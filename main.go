package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"gin_graphql/app/validation/customValidateV9"
	"gin_graphql/graph"
	"gin_graphql/graph/directives"
	"gin_graphql/graph/generated"
	"gin_graphql/routes"
	"os"

	_ "gin_graphql/docs"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"

	glmiddleware "gin_graphql/graph/middleware"
)

const defaultGraphQLPort = "8080"

func init() {
	log.SetLevel(log.InfoLevel)
	log.AddHook(newRotateHook("logs", "stdout.log", 3*24*time.Hour, 24*time.Hour))
	// log.AddHook(newRotateHook("logs", "stdout.log", 2*time.Minute, 1*time.Minute))
}

func newRotateHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
	writer, err := rotatelogs.New(
		logPath+"/"+"%Y%m%d_"+logFileName,
		// logPath+"/"+"%Y%m%d_%H%M_"+logFileName,

		// WithLinkName 為最新的 log 建立連結
		rotatelogs.WithLinkName(logFileName),

		// WithRotationTime 設置 log 分割的時間
		rotatelogs.WithRotationTime(rotationTime),

		// WithMaxAge 和 WithRotationCount 只能設置一個
		// WithMaxAge 文件清理前最長保存時間
		// WithRotationCount 文件清理前最多保存個數
		rotatelogs.WithMaxAge(maxAge),
		// rotatelogs.WithRotationCount(maxAge),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &graph.Resolver{}}
	// Schema Directive
	c.Directives.IsAuthenticated = directives.IsAuthenticated
	c.Directives.HasRole = directives.HasRole
	c.Directives.Binding = directives.Binding

	// srv := glmiddleware.AuthMiddleware(handler.NewDefaultServer(generated.NewExecutableSchema(c)))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

/*
@title Swagger Example API With Gin
@version 1.0
@description This is a sample server celler server.
@termsOfService http://swagger.io/terms/

@contact.name API Support
@contact.url http://www.swagger.io/support
@contact.email support@swagger.io

@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html

@host localhost:5566
@BasePath /api/v1
@query.collection.format multi
@x-extension-openapi {"example": "value on a json format"}
*/
func main() {
	// defer db.GormDB.Close()
	// defer db.SqlDB.Close()

	var env = "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
		switch env {
		case "app":
			// fmt.Println("----- run app env -----")
			godotenv.Load(".env")
		case "dev":
			// fmt.Println("----- run develop env -----")
			godotenv.Load("./.env.dev")
		case "qa":
			// fmt.Println("----- run qa env -----")
			godotenv.Load("./.env.qa")
		default:
			// fmt.Println("----- run default env (dev) -----")
			godotenv.Load("./.env.dev")
		}

	} else {
		// fmt.Println("----- run default env (dev)-----")
		godotenv.Load("./.env.dev")
	}

	// GraphQL Server >>>
	graphQLPort := os.Getenv("GRAPHQL_PORT")
	if graphQLPort == "" {
		graphQLPort = defaultGraphQLPort
	}
	graphQLPath := os.Getenv("GRAPHQL_PATH")
	if graphQLPath == "" {
		graphQLPath = "/query"
	}

	/* 另開 port 的 graphql
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", graphQLPath))
	http.Handle(graphQLPath, srv)

	go func() {
		log.Printf("connect to http://localhost:%s/%s for GraphQL playground", graphQLPort, graphQLPath)
		log.Fatal(http.ListenAndServe(":"+graphQLPort, nil))
	}()
	*/
	// GraphQL Server <<<

	// GIN binding validation version
	customValidateV9.Start()

	r := routes.SetupRouter()

	if os.Getenv("GRAPHQL") == "1" {
		// gin 的 router 結合:  GraphQL palyground : http://localhost:5566/graphql
		r.GET("/_graphql", PlaygroundHandler(graphQLPath))
		// GraphQL query : http://localhost:5566/graphql/query
		r.POST("/_graphql"+graphQLPath, glmiddleware.AuthMiddlewareGin(), GraphqlHandler())
		// r.POST("/graphql"+graphQLPath, GraphqlHandler())
	}

	// Listen and Server
	port := os.Getenv("APP_URL")
	addrs, _ := net.InterfaceAddrs()

	if os.Getenv("SWAGGER") == "1" {
		// swagger : http://localhost:5566/swagger/index.html
		if mode := gin.Mode(); mode == gin.DebugMode {
			swagURL := ginSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", "localhost:5566"))
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagURL))
		}
	}

	// PPROF
	if os.Getenv("PPROF") == "1" {
		// pprof.Register(r) // 性能
		adminGroup := r.Group("/admin", func(c *gin.Context) {
			if c.Request.Header.Get("Authorization") != "kimi" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			c.Next()
		})
		pprof.RouteRegister(adminGroup, "pprof")
	}

	log.WithFields(log.Fields{
		"ip":   addrs[0],
		"port": port,
		"env":  env,
	}).Info("===== Start Server ===== ")
	log.Fatal(http.ListenAndServe(port, r))

}
