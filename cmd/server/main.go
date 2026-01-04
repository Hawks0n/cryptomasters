package main

import (
	"fmt"

	"cryptomasters/internal/service"
)

func main() {
	fmt.Println(" CryptoMasters started on Kali Linux")

	coins := []string{
		"bitcoin",
		"ethereum",
		"dogecoin",
		"litecoin",
	}

	service.FetchAllPrices(coins)

	fmt.Println("\nAll prices fetched successfully")
}
