package controllers

import (
	"context"
	"fmt"
	"gofiber/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	DB *mongo.Client
}

type SignupRequest struct {
	Name    string `json:"name" bson:"name" validate:"required,email"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}

type authResponse struct {
	ID    uint   `json:"id" bson:"email"`
	Email string `json:"email" bson:"password"`
}

func (a *Auth) Signup(c *fiber.Ctx) error {
	req := new(SignupRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}

	// save this info in the database
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.Users{
		Name:    req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	collection := a.DB.Database("myApp").Collection("users")
	updateResult, err :=  collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(updateResult)

	token, exp, err := createJWTToken(*user)
		if err != nil {
			return err
		}
		// create a jwt token

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user, "status": updateResult})

	
}

func createJWTToken(user models.Users) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
