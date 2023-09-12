package main

import (
	"fmt"
	"os"

	"github.com/theadamz/dotenv/examples/config"
)

func main() {
	// print use OS
	keys := []string{"MY_SECRET", "BASE_URL", "DEBUG", "BUDGET_INIT", "BUDGET_FLOAT"}
	for _, key := range keys {
		fmt.Printf(key+" : %v, type : %T", os.Getenv(key), os.Getenv(key))
		fmt.Println()
	}

	fmt.Println("======================================================================")

	// print use struct
	mySecret := config.Env.MY_SECRET
	fmt.Printf("MY_SECRET : %v, type : %T", mySecret, mySecret)
	fmt.Println()
	fmt.Printf("BASE_URL : %v, type : %T", config.Env.BASE_URL, config.Env.BASE_URL)
	fmt.Println()
	fmt.Printf("DEBUG : %v, type : %T", config.Env.DEBUG, config.Env.DEBUG)
	fmt.Println()
	fmt.Printf("BUDGET_INIT : %v, type : %T", config.Env.BUDGET_INIT, config.Env.BUDGET_INIT)
	fmt.Println()
	fmt.Printf("BUDGET_FLOAT : %v, type : %T", config.Env.BUDGET_FLOAT, config.Env.BUDGET_FLOAT)
	fmt.Println()
}
