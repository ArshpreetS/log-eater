package routes

import (
	controllers "github.com/ArshpreetS/log-ingester/handlers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(es *elasticsearch.Client) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/", func(c *gin.Context) {
			controllers.Ingest(c, es)
		})
	}

	return r
}
