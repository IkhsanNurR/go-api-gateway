package middleware

import (
	appconfig "api-gateway/config/app_config"
	pathconfig "api-gateway/config/path_config"
	"api-gateway/response"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfgPath, err := pathconfig.LoadConfigPath("./public_path.yaml")
		if err != nil {
			log.Fatalf("Failed to load path config: %v", err)
		}
		// Cek API Key terlebih dahulu untuk semua request
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" || apiKey != appconfig.APIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse(
				nil,
				"Unauthorized. Invalid or missing API Key.",
				http.StatusUnauthorized,
				nil,
			))
			return
		}

		// Path yang hanya perlu API key saja (tanpa token)
		publicPaths := cfgPath.PublicPaths
		if publicPaths == nil {
			publicPaths = []string{}
		}
		currentPath := c.Request.URL.Path
		for _, path := range publicPaths {
			if strings.EqualFold(currentPath, path) {
				// Hanya cek API Key untuk path ini
				c.Next()
				return
			}
		}

		// Untuk path lain, cek token juga
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse(
				nil,
				"Unauthorized. Missing or invalid token!",
				http.StatusUnauthorized,
				nil,
			))
			return
		}

		// TODO: Tambahkan logika validasi token di sini jika ada
		c.Next()
	}
}
