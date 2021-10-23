package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type article struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}


func Serve(r *gin.Engine) {

	// /articles?categoryId=1
	/* r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	}) */
		//ชุดข้อมูลเดิม
	// GET /api/v1/articles
	// GET /api/v1/articles:id
	// GET /api/v2/articles
	// GET /api/v2/articles:id

	articles := []article{
		{ID: 1,Title: "Title1", Body: "Body1"},
		{ID: 2,Title: "Title2", Body: "Body2"},
		{ID: 3,Title: "Title3", Body: "Body3"},
		{ID: 4,Title: "Title4", Body: "Body4"},
		{ID: 5,Title: "Title5", Body: "Body5"},
	}

	type createArticleForm struct {
		Title string `json:"title" binding:"required"` // ส่งเล็กแต่แปลงไปเป็นใหญ่
		Body string `json:"body" binding:"required"` // binding = validate
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	articlesGroup := r.Group("/api/v1/articles")

	articlesGroup.GET("/", func(c *gin.Context) {
		result := articles
		if limit := c.Query("limit"); limit != ""{
			n, _ := strconv.Atoi(limit)

			result = result[:n]
		}
		
		c.JSON(http.StatusOK, gin.H{"articles" : result})
	})

	articlesGroup.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for _ /* index */, item:= range articles {
			if item.ID == uint(id){
				c.JSON(http.StatusOK, gin.H{"article": item})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error" : "Article not fould"})
	})

	articlesGroup.POST("", func(c *gin.Context) {
		var form createArticleForm
		if err := c.ShouldBindJSON(&form); err!=nil{// &form เขียนทับ
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} 
		
		a := article{
			ID: uint(len(articles) + 1),
			Title: form.Title,
			Body: form.Body,
		}

		articles = append(articles, a)

		c.JSON(http.StatusCreated, gin.H{"article": a})
	})

}