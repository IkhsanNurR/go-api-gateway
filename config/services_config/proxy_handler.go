package servicesconfig

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	response "api-gateway/response"
)

func ProxyHandler(cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, svc := range cfg.Services {
			if strings.HasPrefix(path, svc.PathPrefix) {
				trimmedPath := strings.TrimPrefix(path, svc.PathPrefix)
				if !strings.HasPrefix(trimmedPath, "/") {
					trimmedPath = "/" + trimmedPath
				}
				target := svc.TargetURL + trimmedPath

				// Baca body request jika ada
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, response.FailedResponse(
						nil,
						"Failed to read request body",
						http.StatusInternalServerError,
						err.Error(),
					))
					return
				}
				// Reset body supaya bisa dibaca ulang jika perlu middleware lain
				c.Request.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

				// Buat request baru ke service target
				req, err := http.NewRequest(c.Request.Method, target, strings.NewReader(string(bodyBytes)))
				if err != nil {
					c.JSON(http.StatusInternalServerError, response.FailedResponse(
						nil,
						"Failed to create proxy request",
						http.StatusInternalServerError,
						err.Error(),
					))
					return
				}

				// Copy semua header dari request asli
				for k, v := range c.Request.Header {
					for _, vv := range v {
						req.Header.Add(k, vv)
					}
				}

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					c.JSON(http.StatusBadGateway, response.FailedResponse(
						nil,
						"Service Unavailable",
						http.StatusBadGateway,
						nil,
					))
					return
				}
				defer resp.Body.Close()

				// Copy semua header dari response target ke response client
				for k, v := range resp.Header {
					for _, vv := range v {
						c.Writer.Header().Add(k, vv)
					}
				}

				// Set status code dari response target
				c.Writer.WriteHeader(resp.StatusCode)

				// Streaming langsung responsenya ke client
				_, err = io.Copy(c.Writer, resp.Body)
				if err != nil {
					// opsional log error jika streaming gagal
					log.Println("Error streaming response:", err)
				}

				return
			}
		}

		// Jika tidak ada service yang cocok dengan path
		c.JSON(http.StatusNotFound, response.FailedResponse(
			nil,
			"Route Not Found",
			http.StatusNotFound,
			nil,
		))
	}
}
