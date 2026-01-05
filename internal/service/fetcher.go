package service

import (
	"fmt"
	"sync"

	"cryptomasters/internal/client"
	"cryptomasters/internal/config"
)

func FetchAllPrices(coins []string) {
	var wg sync.WaitGroup
	wg.Add(len(coins))

	for _, coin := range coins {
		go func(c string) {
			defer wg.Done()

			url := fmt.Sprintf(
				"%s?ids=%s&vs_currencies=usd",
				config.BaseURL,
				c,
			)

			prices, err := client.FetchPrices(url)
			if err != nil {
				fmt.Printf("[ERROR] %s → %v\n", c, err)
				return
			}

			fmt.Printf("\n%s price data:\n", c)
			for name, data := range prices {
				fmt.Printf("  %s → $%.2f\n", name, data["usd"])
			}
		}(coin)
	}

	wg.Wait()
}

