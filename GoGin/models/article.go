package models

type Article2 struct {
	ID    uint   `form:"id" json:"id" xml:"id"  binding:"required"`
	Title string `form:"title" json:"title" xml:"title"  binding:"required"`
	Body  string `form:"body" json:"body" xml:"body"  binding:"required"`
	Image string `form:"image" json:"image" xml:"image"  binding:"required"`
}
