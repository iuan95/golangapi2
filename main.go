package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/iuan95/golangapi2/db"
	"github.com/iuan95/golangapi2/handlers"
)
func init(){
	err:=db.Connect()
	if err!=nil {
		panic(err)
	}
}

func main(){

	app:=fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{"message": "home"})
	})
	items:= app.Group("/items")
	items.Get("/", func(c *fiber.Ctx) error {
		items, err:=handlers.GetAllItems(context.Background())
		fmt.Println(err)
		if err!= nil {
			return c.Status(400).JSON(&fiber.Map{
				"message": "can not find items",
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message":"ok",
			"data": items,
		})
	})
	items.Get("/:id", func(c *fiber.Ctx) error {
		id:=c.Params("id")
		item,err:= handlers.GetItemById(context.Background(), id)
		fmt.Println(err)
		if err!=nil {
			return c.Status(400).JSON(&fiber.Map{
				"message": "can not find item",
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message":"ok",
			"data": item,
		})
	})

	items.Post("/", func(c *fiber.Ctx) error {
		name:= c.Params("name")
		desc:=c.Params("description")
		err:= handlers.CreateItem(context.Background(), name, desc)
		fmt.Println(err)
		if err!=nil {
			return c.Status(400).JSON(&fiber.Map{
				"message": "can not create item",
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "success",
		})
	})

	app.Listen(":3000")

}