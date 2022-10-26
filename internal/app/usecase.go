package app

import (
	"fmt"
	"log"

	"github.com/tetsurowada/test_server_2022/internal/database"
)

func getOwnedItems(param GetOwnedItemsInput) (GetOwnedItemsOutput, error) {
	result := GetOwnedItemsOutput{Items: make(GetOwnedItemsOutputItems, 0)}
	db, err := database.Connect()
	if err != nil {
		log.Println(err.Error())
		return result, err
	}
	defer db.Close()
	stmt := fmt.Sprintf("SELECT * FROM db1.items where item_owner_id = %v", param.UserId)
	rows, err := db.Query(stmt)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var goioi GetOwnedItemsOutputItem
		err := rows.Scan(&goioi.ItemId, &goioi.ItemName, &goioi.CategoryId, &goioi.ItemOwnerId)
		if err != nil {
			log.Println(err.Error())
			return result, err
		}
		result.Items = append(result.Items, goioi)
		log.Println(goioi.ItemId, goioi.ItemName)
	}

	return result, err
}
