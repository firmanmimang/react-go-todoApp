package controllers

import (
	"context"
	"time"

	"github.com/firmanmimang/react-go-todo/config"
	"github.com/firmanmimang/react-go-todo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.DB.Collection("todos")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error fetching todos"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error decoding todo"})
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.DB.Collection("todos")
	result, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error inserting todo"})
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.DB.Collection("todos")
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error updating todo"})
	}

	return c.JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.DB.Collection("todos")
	filter := bson.M{"_id": objID}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error deleting todo"})
	}

	return c.JSON(fiber.Map{"success": true})
}
