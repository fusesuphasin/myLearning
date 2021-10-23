package routes

import (
	"go-gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormServe(r *gin.Engine) {

	// /articles?categoryId=1
	/* r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	}) */
	//ชุดข้อมูลเดิม
	// GET /api/v1/articles
	// GET /api/v1/articles:id
	// GET /api/v2/articles
	// GET /api/v2/articles:id

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	articlesGroup := r.Group("/api/v2/articles")
	articlesController := controllers.Article{}

	{
		articlesGroup.GET("/", articlesController.FindAll)
		articlesGroup.GET("/:id", articlesController.FindOne)
		articlesGroup.POST("", articlesController.Create)
	}
}
