package main

import (
	"net"
	"net/http"
	"time"

	"gin_graphql/app/validation/customValidateV9"
	"gin_graphql/graph"
	"gin_graphql/graph/generated"
	"gin_graphql/routes"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", graphQLPath))
	http.Handle(graphQLPath, srv)

	go func() {
		log.Printf("connect to http://localhost:%s/%s for GraphQL playground", graphQLPort, graphQLPath)
		log.Fatal(http.ListenAndServe(":"+graphQLPort, nil))
	}()
	// GraphQL Server <<<

	// GIN binding validation version
	customValidateV9.Start()

	r := routes.SetupRouter()

	// Listen and Server
	port := os.Getenv("APP_URL")
	addrs, _ := net.InterfaceAddrs()
	log.WithFields(log.Fields{
		"ip":   addrs[0],
		"port": port,
		"env":  env,
	}).Info("===== Start Server ===== ")

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

	log.Fatal(http.ListenAndServe(port, r))

}
