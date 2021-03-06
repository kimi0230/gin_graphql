package corsConfig

import (
	"os"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	// https://github.com/gin-contrib/cors
	corsConf := cors.DefaultConfig()

	if os.Getenv("APP_ENV") == "APP" {
		corsConf.AllowOrigins = []string{"https://github.com"}

		// corsConf.AllowHeaders = []string{"Origin"}

		// If this option is set, the content of AllowOrigins is ignored.
		// corsConf.AllowOriginFunc = func(origin string) bool {
		// 	return origin == "https://github.com"
		// }

	} else if os.Getenv("APP_ENV") == "DEV" {
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PUT", "HEAD"}
	} else {
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PUT", "HEAD"}
	}

	return corsConf
}
