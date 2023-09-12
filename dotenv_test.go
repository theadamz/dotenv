package dotenv

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type Config struct {
	MY_STRING                     string
	MY_STRING_SECRET              string
	MY_STRING_QUOTES_SPACE_QUOTES string
	MY_BOOL_TRUE                  bool
	MY_NUMBER_INT                 int // or int64
	MY_NUMBER_FLOAT               float64
}

var AppConfig Config

func TestMain(t *testing.T) {
	files := []string{"./.env"}
	err := Load(files...)
	if err != nil {
		t.Fatal(err.Error())
	}

	// test os env
	tests := []struct {
		key    string
		expect string
	}{
		{
			key:    "MY_STRING",
			expect: "foo",
		},
		{
			key:    "MY_STRING_QUOTES",
			expect: "bar",
		},
		{
			key:    "MY_STRING_SPACE",
			expect: "foo bar",
		},
		{
			key:    "MY_STRING_QUOTES_SPACE_QUOTES",
			expect: "foo bar",
		},
		{
			key:    "MY_URL",
			expect: "http://www.google.com",
		},
		{
			key:    "MY_URL_QUOTES",
			expect: "http://www.google.com",
		},
		{
			key:    "MY_EMAIL",
			expect: "my_email@email.com",
		},
		{
			key:    "MY_EMAIL_QUOTES",
			expect: "my_email@email.com",
		},
		{
			key:    "MY_BOOL_TRUE",
			expect: "true",
		},
		{
			key:    "MY_BOOL_FALSE",
			expect: "false",
		},
		{
			key:    "MY_BOOL_BUT_STRING",
			expect: "false",
		},
		{
			key:    "MY_NUMBER_INT",
			expect: "123456789000",
		},
		{
			key:    "MY_NUMBER_FLOAT",
			expect: "1234567890.0589",
		},
		{
			key:    "MY_NUMBER_INT_BUT_STRING",
			expect: "123456789000",
		},
		{
			key:    "MY_STRING_SECRET",
			expect: "secret#bar123",
		},
	}

	// run loop tests
	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			if os.Getenv(test.key) != test.expect {
				t.Error("result must be: " + test.expect)
			}
		})
	}
}

func TestStruct(t *testing.T) {
	files := []string{"./.env", "./.env"}
	parsed, err := LoadToMap(files)

	if err != nil {
		t.Fatal(err.Error())
	}

	// Convert the map to JSON
	jsonData, _ := json.Marshal(parsed)

	// Convert the JSON to a struct
	json.Unmarshal(jsonData, &AppConfig)

	// test string
	t.Run("MY_STRING", func(t *testing.T) {
		expect := "foo"

		// check value
		if AppConfig.MY_STRING != expect {
			t.Error("result must be:", expect)
		}

		// check type
		if getType(AppConfig.MY_STRING) != reflect.TypeOf(AppConfig.MY_STRING) {
			t.Error("result type must be:", reflect.TypeOf(AppConfig.MY_STRING))
		}
	})

	t.Run("MY_STRING_SECRET", func(t *testing.T) {
		expect := "secret#bar123"

		// check value
		if AppConfig.MY_STRING_SECRET != expect {
			t.Error("result must be:", expect)
		}

		// check type
		if getType(AppConfig.MY_STRING_SECRET) != reflect.TypeOf(AppConfig.MY_STRING_SECRET) {
			t.Error("result type must be:", reflect.TypeOf(AppConfig.MY_STRING_SECRET))
		}
	})

	t.Run("MY_STRING_QUOTES_SPACE_QUOTES", func(t *testing.T) {
		expect := "foo bar"

		// check value
		if AppConfig.MY_STRING_QUOTES_SPACE_QUOTES != expect {
			t.Error("result must be:", expect)
		}

		// check type
		if getType(AppConfig.MY_STRING_QUOTES_SPACE_QUOTES) != reflect.TypeOf(AppConfig.MY_STRING_QUOTES_SPACE_QUOTES) {
			t.Error("result type must be:", reflect.TypeOf(AppConfig.MY_STRING_QUOTES_SPACE_QUOTES))
		}
	})

	t.Run("MY_BOOL_TRUE", func(t *testing.T) {
		expect := true

		// check value
		if AppConfig.MY_BOOL_TRUE != expect {
			t.Error("result must be:", expect)
		}

		// check type
		if getType(AppConfig.MY_BOOL_TRUE) != reflect.TypeOf(AppConfig.MY_BOOL_TRUE) {
			t.Error("result type must be:", reflect.TypeOf(AppConfig.MY_BOOL_TRUE))
		}
	})

	t.Run("MY_NUMBER_INT", func(t *testing.T) {
		expect := 123456789000

		// check value
		if AppConfig.MY_NUMBER_INT != expect {
			t.Error("result must be:", expect)
		}

		// check type
		if getType(AppConfig.MY_NUMBER_INT) != reflect.TypeOf(AppConfig.MY_NUMBER_INT) {
			t.Error("result type must be:", reflect.TypeOf(AppConfig.MY_NUMBER_INT))
		}
	})
}

func getType(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}
