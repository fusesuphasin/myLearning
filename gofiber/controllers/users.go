package controllers

import (
	"context"
	"fmt"
	"gofiber/models"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	DB *mongo.Client
}

type User struct {
	Name string `bson:"name" form:"name" json:"name" xml:"name" `
	Email    string `bson:"email" form:"email" json:"email" xml:"email" `
	Password string `bson:"password" form:"password" json:"password" xml:"password" `
	Image    string `bson:"image" form:"image" json:"image" xml:"image" `
}

type usersPaging struct {
	Items []UsersRespone `bson:"items" json:"items"`
	Paging *pagingResult `bson:"paging" json:"paging"`
}

type createUsers struct {
	Name string `bson:"name" form:"name" json:"name" xml:"name"  validate:"required"`
	Email    string `bson:"email" form:"email" json:"email" xml:"email"  validate:"required"`
	Password string `bson:"password" form:"password" json:"password" xml:"password"  validate:"required"`
	Image    *multipart.FileHeader `bson:"image" form:"image" json:"image" xml:"image"  validate:"required"`
}

type formUsers struct{
	Name string `bson:"name,omitempty" form:"name,omitempty" json:"name,omitempty" xml:"name" `
	Password string `bson:"password,omitempty" form:"password,omitempty" json:"password,omitempty" xml:"password" `
	Image    string `bson:"image,omitempty" form:"image,omitempty" json:"image,omitempty" xml:"image"  `
}

type updateUsers struct{
	Name string `bson:"name,omitempty" form:"name,omitempty" json:"name,omitempty" xml:"name" `
	Password string `bson:"password,omitempty" form:"password,omitempty" json:"password,omitempty" xml:"password" `
	Image    *multipart.FileHeader `bson:"image,omitempty" form:"image,omitempty" json:"image,omitempty" xml:"image"  `
}

type UsersRespone struct{
	ID    primitive.ObjectID `bson:"_id" form:"_id" json:"_id,omitempty"`
	//ID uint `bson:"id" form:"id" json:"id" xml:"id"`
	Name string `bson:"name" form:"name" json:"name" xml:"name" `
	Email    string `bson:"email" form:"email" json:"email" xml:"email"`
	Password string `bson:"password" form:"password" json:"password" xml:"password" `
	Image    string `bson:"image" form:"image" json:"image" xml:"image" `
}

var users []models.Users = []models.Users{
	{/* ID: 1, */ Name: "A", Email: "a@gmail.com", Password: "123"},
	{/* ID: 2, */ Name: "B", Email: "b@gmail.com", Password: "123"},
	{/* ID: 3, */ Name: "C", Email: "c@gmail.com", Password: "123"},
	{/* ID: 4, */ Name: "D", Email: "d@gmail.com", Password: "123"},
	{/* ID: 5, */ Name: "E", Email: "e@gmail.com", Password: "123"},
}

func (u Users) FindAll(c *fiber.Ctx) error{

	collection := u.DB.Database("myApp").Collection("users")
	cursor, err := collection.Find(context.TODO(),bson.D{{}})
	if err != nil {
		return c.JSON(err)
	}
	
	var results []models.Users
	if err = cursor.All(context.TODO(), &results); err != nil {
		return c.JSON(err)
	}

	/// /articles => limit => 12, page => 1
	// /articles?limit=10 => limit => 10, page => 1
	// /articles?page=10 => limit => 12, page => 10
	// /articles?page=2&limit=4 => limit => 4, page => 2
	//paging := pagingResource(c, u.DB)
	var serializedUsers []UsersRespone
	copier.Copy(&serializedUsers, &results)
	return c.Status(200).JSON(serializedUsers /* usersPaging{Items:results, Paging: } */)
}

func (u Users) FindLimit(c *fiber.Ctx) error{

	paging := pagingResource()
	
	return c.Status(200).JSON(paging)
}

func (u *Users) FindOne(c *fiber.Ctx) (error){
	user,err := u.findUserByID(c)
	if err != nil{
		return c.Status(http.StatusNotFound).JSON("user not found")
	}
	serializedUser := UsersRespone{}
	copier.Copy(&serializedUser, &user)
	c.Status(http.StatusOK).JSON(serializedUser)
	return c.Status(http.StatusOK).JSON(user)
}

func (u *Users) Create(c *fiber.Ctx) error{
	var form createUsers
	if err := c.BodyParser(&form); err!=nil{// &form เขียนทับ
		return c.Status(http.StatusBadRequest).JSON("can not get users")
	} 
	
	var user User
	/* collection := u.DB.Database("myApp").Collection("users")
	countID, err := collection.EstimatedDocumentCount(context.TODO())
	if err != nil {
		panic(err)
	}
	copier.Copy(&user, models.Users{
		ID: uint(countID + 1),
	} ) */
	copier.Copy(&user, &form)

	//u.DB.Database("myapp").Collection("users")
	//collection := u.DB.Database("myapp").Collection("users")
	//fmt.Println(collection)
	u.setUserImage(c, &user)
	serializedUser := UsersRespone{}
	copier.Copy(&serializedUser, &user)
	//users = append(users, user)

	return c.Status(http.StatusCreated).JSON(user)
}

func (u *Users) Update(c *fiber.Ctx) error{
	/* id := c.Params("id")
	
	objectId, err := primitive.ObjectIDFromHex(id) */
	
	var form formUsers
	if err := c.BodyParser(&form); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON("err: " + err.Error())
	}
	
	user, err := u.findUserByID(c)
	collection := u.DB.Database("myApp").Collection("users")
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("err: " + err.Error())
	}
	filter := bson.D{{"_id", user.ID}}
	
	// Get file
	file, err := c.FormFile("Image")
	if err != nil || file == nil{
		return c.Status(http.StatusBadRequest).JSON("message" + "No file is received")
	}
	fmt.Println(file.Filename ," WHATTT")

	// ถ้ามีรูปอยู่ก่อนให้ลบออก
	user.Image = strings.Replace(user.Image, os.Getenv("HOST"), "", 1)
	pwd, _ := os.Getwd()
	newPath := pwd + "/uploads/users/" + ((user.Email)) + "/oldfile/"
	os.MkdirAll(newPath, 0777)
	err =  os.Rename(pwd + user.Image, newPath)
	if err != nil {
		fmt.Println( err)
	}
	os.Remove(pwd + user.Image)
	fmt.Println(user.Image)
	// Create Path
	// ID => 8, uploads/articles/8/image.png
	path := "uploads/users/" + ((user.Email))
	os.MkdirAll(path, 0755)

	// Upload File
	filename := path + "/" + file.Filename

	if err := c.SaveFile(file, filename); err != nil {
		return err
	}

	// Attach file to article2
	form.Image = os.Getenv("HOST") + "/" + filename
	fmt.Println(form.Image ," WHATTT2")

	update := bson.D{{Key: "$set", Value: form}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(http.StatusNotModified).JSON("err: " + err.Error())
	}
	
	serializedUser := UsersRespone{}
	copier.Copy(&serializedUser, &result)
	return c.Status(http.StatusOK).JSON(result)
}

func (u *Users) Delete(c *fiber.Ctx) error{
	/* id := c.Params("id")
	
	objectId, err := primitive.ObjectIDFromHex(id) */
	user, err := u.findUserByID(c)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("err: " + err.Error())
	}
	collection := u.DB.Database("myApp").Collection("users")
	filter := bson.D{{"_id", user.ID}}
	//filter := bson.D{{"id", &user.ID}}
	result, err := collection.DeleteOne(context.TODO(),filter)
	if err != nil {
		return c.Status(400).JSON("err: " + err.Error())
	}
	fmt.Println(result)
	return c.Status(http.StatusOK).JSON("Delete Successfully")
}

func (u *Users) setUserImage(c *fiber.Ctx, user *User) error{
	
	// connect db
	collection := u.DB.Database("myApp").Collection("users")

	// Get file
	file, err := c.FormFile("Image")
	if err != nil || file == nil{
		return c.Status(http.StatusBadRequest).JSON("message" + "No file is received")
	}

	// ถ้ามีรูปอยู่ก่อนให้ลบออก
	if user.Image != "" {
		user.Image = strings.Replace(user.Image, os.Getenv("HOST"), "", 1)
		pwd, _ := os.Getwd()
		os.Remove(pwd + user.Image)
	}

	// Create Path
	// ID => 8, uploads/articles/8/image.png
	path := "uploads/users/" + ((user.Email))
	os.MkdirAll(path, 0755)

	// Upload File
	filename := path + "/" + file.Filename

	if err := c.SaveFile(file, filename); err != nil {
		return err
	}

	// Attach file to article2
	user.Image = os.Getenv("HOST") + "/" + filename
	fmt.Println(user.Image + " WHAT!")
	updateResult, err :=  collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("InspertUser :", updateResult)
	return nil
}

func (u *Users) findUserByID(c *fiber.Ctx) (UsersRespone, error){
	id := c.Params("id")
	
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {fmt.Println(err)}
	//newid := uint(id)
	filter := bson.M{"_id": objectId}
	collection := u.DB.Database("myApp").Collection("users")
	var result UsersRespone
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, err
} 