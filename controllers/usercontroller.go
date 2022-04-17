package controllers

import (
	"admybrand/configs"
	"admybrand/model"
	"admybrand/response"
	"context"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HelloWorld(c *fiber.Ctx) {
	c.JSON(&fiber.Map{
		"api/users":           "To Get All The Users",
		"api/user/:id":        "To Get Single User",
		"api/user/insert":     "To Insert New User",
		"api/user/update/:id": "To Update User",
		"api/user/delete/:id": "To Delete User",
	})
}

func GetAllUsers(c *fiber.Ctx) {
	db := configs.DB
	var user model.User
	var users []model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := db.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrProcessing.Code,
			Message: "Error in Finding All Users",
			Data:    &fiber.Map{"data": err},
		})
	}
	for cursor.Next(ctx) {
		if err := cursor.Decode(&user); err != nil {
			c.JSON(response.UserResponse{
				Status:  fiber.ErrProcessing.Code,
				Message: "Error in Decoding",
				Data:    &fiber.Map{"data": err},
			})
		}
		users = append(users, user)
	}
	c.JSON(response.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": users},
	})
}

func GetAUser(c *fiber.Ctx) {
	db := configs.DB
	var user model.User
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := db.FindOne(ctx, bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(response.UserResponse{
				Status:  fiber.ErrNotFound.Code,
				Message: "User Not Found",
				Data:    &fiber.Map{"data": err},
			})
		}
	}
	c.JSON(response.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"data": user},
	})
}

func InsertAUser(c *fiber.Ctx) {
	db := configs.DB
	var user model.User
	// var temp model.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrProcessing.Code,
			Message: "Error in Parsing Body",
			Data:    &fiber.Map{"data": err},
		})
	}

	// err := db.FindOne(ctx, bson.M{"id": user.Id}).Decode(&temp)
	// fmt.Println(user.Id == temp.Id)
	// if temp.Id == user.Id {
	// 	c.JSON(response.UserResponse{
	// 		Status:  fiber.ErrProcessing.Code,
	// 		Message: "User Already Found",
	// 		Data:    &fiber.Map{"data": err},
	// 	})
	// }

	result, err := db.InsertOne(ctx, user)
	if err != nil {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrProcessing.Code,
			Message: "Error in Inserting User",
			Data:    &fiber.Map{"data": err},
		})
	}
	c.JSON(response.UserResponse{
		Status:  fiber.StatusOK,
		Message: "User Inserted",
		Data:    &fiber.Map{"data": result},
	})
}

func UpdateAUser(c *fiber.Ctx) {
	db := configs.DB
	var user model.User
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrProcessing.Code,
			Message: "Error in Parsing Body",
			Data:    &fiber.Map{"data": err},
		})
	}
	Upsert := true
	opt := &options.FindOneAndUpdateOptions{
		Upsert: &Upsert,
	}
	err := db.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": user}, opt).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(response.UserResponse{
				Status:  fiber.ErrNotFound.Code,
				Message: "User Not Found",
				Data:    &fiber.Map{"data": err},
			})
		} else {
			c.JSON(response.UserResponse{
				Status:  fiber.ErrProcessing.Code,
				Message: "Error in Updating User",
				Data:    &fiber.Map{"data": err},
			})
		}
	}
	c.JSON(response.UserResponse{
		Status:  fiber.StatusOK,
		Message: "User Updated",
		Data:    &fiber.Map{"data": user},
	})
}

func DeleteAUser(c *fiber.Ctx) {
	db := configs.DB
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrProcessing.Code,
			Message: "Error in Deleting User",
			Data:    &fiber.Map{"data": err},
		})
	}
	if result.DeletedCount == 1 {
		c.JSON(response.UserResponse{
			Status:  fiber.StatusOK,
			Message: "User Deleted",
			Data:    &fiber.Map{"data": id},
		})
	} else {
		c.JSON(response.UserResponse{
			Status:  fiber.ErrNotFound.Code,
			Message: "User Not Found",
			Data:    &fiber.Map{"data": id},
		})
	}
}
