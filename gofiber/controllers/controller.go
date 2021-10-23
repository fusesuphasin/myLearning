package controllers

import (
	"context"
	"fmt"
	"gofiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pagingResult struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	PrevPage  int `json:"prevPage"`
	NextPage  int `json:"nextPage"`
	Count     int `json:"count"`
	TotalPage int `json:"totalPage"`
}

var c *fiber.Ctx
var DB *mongo.Client

func pagingResource() interface{}{
	// 1. Get limit, page ?limit=10&page=2
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))
	
	collection := DB.Database("myApp").Collection("users")
	fmt.Println(collection)
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"id", -1}}).SetLimit(int64(limit)).SetSkip(int64((page * 1) * limit))
	curcer, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		fmt.Println(err)
	}
	var results []models.Users
	if err = curcer.All(context.TODO(), &results); err != nil {
		fmt.Println(err)
	}

	/// /articles => limit => 12, page => 1
	// /articles?limit=10 => limit => 10, page => 1
	// /articles?page=10 => limit => 12, page => 10
	// /articles?page=2&limit=4 => limit => 4, page => 2
	var serializedUsers []UsersRespone
	copier.Copy(&serializedUsers, &results)
	
	

	// 2. count records
	count, err := collection.EstimatedDocumentCount(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("number in collection: %d\n", count)

	// 3. Find Records
	//limit, offset(skip)
	// page => 1, 1 - 10, offset => 0
	// page => 2, 11 - 20, offset => 10
	// page => 3, 21 - 30, offset => 20
	
	// 4. total page
	/* totalPage := int(math.Ceil(float64(count)/float64(limit)))

	// 5. Find nextPage
	var nextPage int
	if page == totalPage{
		nextPage = totalPage
	}else{
		nextPage = -1
	} */

	// 6. create pagingResult
	return nil/* pagingResult{
		Page:      page,
		Limit:     limit,
		Count:     int(count),
		PrevPage:  page - 1,
		NextPage:  nextPage,
		TotalPage: totalPage,
	} */
}
