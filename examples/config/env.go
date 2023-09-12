package config

import (
	"encoding/json"
	"log"

	"github.com/theadamz/dotenv"
)

type Environment struct {
	MY_SECRET    string
	BASE_URL     string
	DEBUG        bool
	BUDGET_INIT  int
	BUDGET_FLOAT float64
}

var Env Environment

func init() {
	LoadEnv()
}

func LoadEnv(files ...string) {
	if len(files) == 0 {
		files = []string{"./.env"}
	}

	parsed, err := dotenv.LoadToMap(files)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Convert the map to JSON
	jsonData, _ := json.Marshal(parsed)

	// Convert the JSON to a struct
	json.Unmarshal(jsonData, &Env)
}
