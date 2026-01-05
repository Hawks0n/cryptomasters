CryptoMasters
CryptoMasters is a small Go project I built to practice concurrency, HTTP requests, and project structure in Go.
It fetches live cryptocurrency prices using goroutines and prints the results.


What it does

- Fetches crypto prices (Bitcoin, Ethereum, etc.)
- Makes multiple API calls concurrently
- Uses `WaitGroup` to synchronize goroutines
- Parses JSON responses into Go data


Project structure

cryptomasters/
├── cmd/server/main.go # Entry point
├── internal/
│ ├── client/ # HTTP requests
│ ├── config/ # Configuration
│ ├── models/ # Data models
│ └── service/ # Concurrency logic
├── go.mod
└── LICENSE


How to run


go mod tidy
go run cmd/server/main.go
![Screenshot_2026-01-05_09_28_09_page-0001](https://github.com/user-attachments/assets/3ba0393d-b1e0-4c75-84d2-2dfb07746bcb)


Example output:
![Screenshot_2026-01-05_09_28_09_page-0002](https://github.com/user-attachments/assets/177636df-eb0b-4c54-9f68-cffc89e63ba2)

nginx
Copy code
CryptoMasters started

bitcoin → $43210.00
ethereum → $2310.45
Done

