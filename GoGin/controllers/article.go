package controllers

import (
	"fmt"
	"go-gin/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)


type Article struct {

}


type createArticleForm struct {
	Title string `form:"title" json:"title" xml:"title"  ` // ส่งเล็กแต่แปลงไปเป็นใหญ่
	Body string `form:"body" json:"body" xml:"body"  ` // binding = validate
	Image *multipart.FileHeader `form:"image" json:"image" xml:"image"  `
}

var articles []models.Article2 = []models.Article2{
	{ID: 1,Title: "Title1", Body: "Body1"},
	{ID: 2,Title: "Title2", Body: "Body2"},
	{ID: 3,Title: "Title3", Body: "Body3"},
	{ID: 4,Title: "Title4", Body: "Body4"},
	{ID: 5,Title: "Title5", Body: "Body5"},
}


func (a *Article) FindAll(c *gin.Context) {
	result := articles
	if limit := c.Query("limit"); limit != ""{
		n, _ := strconv.Atoi(limit)

		result = result[:n]
	}
	
	c.JSON(http.StatusOK, gin.H{"articles" : result})
}

func (a *Article) FindOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _ /* index */, item:= range articles {
		if item.ID == uint(id){
			c.JSON(http.StatusOK, gin.H{"article": item})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "Article not fould"})
}

func (a *Article) Create(c *gin.Context) {
	fmt.Print(555)
	var form createArticleForm
	if err := c.ShouldBind(&form); err!=nil{// &form เขียนทับ
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	
	article := models.Article2{
		ID: uint(len(articles) + 1),
		Title: form.Title,
		Body: form.Body,
	}

	// Get file
	file, err := c.FormFile("Image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	// Create Path
	// ID => 8, uploads/articles/8/image.png
	path := "uploads/articles/" + strconv.Itoa(int(article.ID))
	os.MkdirAll(path, 0755)


	// Upload File
	filename := path + "/" + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		// ...
	}

	// Attach file to article2
	article.Image = os.Getenv("HOST") + "/" + filename

	articles = append(articles, article)

	c.JSON(http.StatusCreated, gin.H{"article": article})
}