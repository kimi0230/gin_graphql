package corsConfig

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	// https://github.com/gin-contrib/cors
	corsConf := cors.Config{
		MaxAge:           12 * time.Hour,
		AllowCredentials: false,
	}

	if os.Getenv("APP_ENV") == "APP" {
		corsConf.AllowOrigins = []string{"https://foo.com"}
		corsConf.AllowHeaders = []string{"Origin"}

		// If this option is set, the content of AllowOrigins is ignored.
		// corsConf.AllowOriginFunc = func(origin string) bool {
		// 	return origin == "https://github.com"
		// }

	} else if os.Getenv("APP_ENV") == "DEV" {
		corsConf.AllowAllOrigins = true
		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host"}
	}

	return corsConf
}
