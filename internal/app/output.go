package app

type GetOwnedItemsOutput struct {
	Items GetOwnedItemsOutputItems `json:"items"`
}

type GetOwnedItemsOutputItems []GetOwnedItemsOutputItem

type GetOwnedItemsOutputItem struct {
	ItemId      int    `json:"itemId"`
	ItemName    string `json:"ItemName"`
	CategoryId  int    `json:"CategoryId"`
	ItemOwnerId int    `json:"ItemOwnerId"`
}
