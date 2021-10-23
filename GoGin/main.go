package main

import (
	"go-gin/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	r := gin.Default()

	r.Static("/uploads", "./uploads") //ให้มองเห็น

	uploadDirs := [...]string{"articles", "users"}

	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755 /*Unix permission calculate*/)
	}

	
	//routes.Serve(r)
	routes.FormServe(r)
	
	r.Run()
	//r.Run(":" + os.Getenv("PORT"))
} 