package handlers

import (
	"fmt"
	"template/database"

	"github.com/gofiber/fiber/v2"
)

type TxResponse struct {
	Txhash    string `json:"txhash"`
	Height    string `json:"height"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Timestamp string `json:"timestamp"`
}

// Handler to get balance by address
func TestAPIS(c *fiber.Ctx) error {
	// PostgreSQL connection string

	// Prepare the insert statement
	insertQuery := `
        INSERT INTO exchange_coinid (coin_id, symbol, name)
        VALUES ($1, $2, $3)`

	// Sample data to insert
	coins := [][]interface{}{}

	// Insert data
	for _, coin := range coins {
		err := database.DB.Db.Exec(insertQuery, coin[0], coin[1], coin[2])

		fmt.Println(err)

	}

	fmt.Println("Data inserted successfully!")
	return nil

}
