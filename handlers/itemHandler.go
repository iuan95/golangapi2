package handlers

import (
	"context"

	"github.com/iuan95/golangapi2/db"
	"github.com/iuan95/golangapi2/models"
)

func CreateItem(ctx context.Context, name string, desc string) error {
	_, err := db.DB.Exec(ctx, "INSERT INTO items (name, description) VALUES ($1, $2)", name, desc)
	return err
}

func DeleteItem(ctx context.Context, id string) error {
	_, err := db.DB.Exec(ctx, "DELETE FROM items WHERE id = $1", id)
	return err
}

func GetItemById(ctx context.Context, id string) (models.Item, error) {
	var item models.Item
	err := db.DB.QueryRow(ctx, "SELECT * FROM items WHERE id = $1", id).Scan(&item.Id, &item.Name, &item.Description)
	return item, err
}

func GetAllItems(ctx context.Context) (*[]models.Item, error) {
	var items []models.Item
	conn, err := db.DB.Query(ctx, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	for conn.Next() {
		var item models.Item
		err := conn.Scan(&item.Id, &item.Name, &item.Description)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return &items, nil
}
