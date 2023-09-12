package dotenv

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var parsed map[string]interface{}

func Load(filenames ...string) error {
	// if filenames not found
	if len(filenames) == 0 {
		return errors.New("please provide file")
	}

	// loop filenames
	for _, filename := range filenames {
		envValues, err := readFile(filename)

		// check for error
		if err != nil {
			return err
		}

		// parse values
		parsed, err = parse(envValues)

		// check for error
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadToMap(filenames []string) (map[string]interface{}, error) {
	// load env file
	err := Load(filenames...)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func parse(mapValues map[string]string) (map[string]interface{}, error) {
	// variables
	var newMap = make(map[string]interface{})

	// loop mapValues
	for key, value := range mapValues {
		// set env
		os.Setenv(key, value)

		// bool
		if value == "true" || value == "false" {
			// parse value
			cvtValue, err := strconv.ParseBool(value)

			// check for error
			if err == nil {
				newMap[key] = cvtValue
				continue
			} else {
				return nil, err
			}
		}

		// int / float
		if regexp.MustCompile(`[[:digit:]]`).MatchString(value) {
			// check again if value is number or float
			if regexp.MustCompile(`^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$`).MatchString(value) {
				// convert string to number
				cvtValue, err := numberHandler(value)

				// check for error
				if err == nil {
					newMap[key] = cvtValue
					continue
				} else {
					return nil, err
				}
			}
		}

		// string
		newMap[key] = stringHandler(value)
	}

	return newMap, nil
}

func numberHandler(s string) (interface{}, error) {
	// check if float / int
	if regexp.MustCompile(`[.][0-9]+$`).MatchString(s) {
		cvtValue, err := strconv.ParseFloat(s, 64)
		return cvtValue, err
	} else {
		cvtValue, err := strconv.Atoi(s)
		return cvtValue, err
	}
}

func readFile(filename string) (map[string]string, error) {
	// variables
	var values map[string]string = make(map[string]string)
	var keyValue string
	var value string

	// Open file
	file, err := os.Open(filename)

	// check for error
	if err != nil {
		return nil, err
	}

	// Create a bufio.Scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// Read and process each line
	for scanner.Scan() {
		// read line
		line := scanner.Text()

		// if line do not contain = then skip
		if !strings.Contains(line, "=") {
			continue
		}

		// split line by =
		lineSplit := strings.Split(line, "=")

		// set key value
		keyValue = strings.Trim(lineSplit[0], " ")

		// if keyValue contain #, it is a comment and ignore
		if keyValue[0] == '#' {
			continue
		}

		// set line and trim
		value = strings.Trim(lineSplit[1], " ")

		// handling value
		value = valueHandler(value)

		// if value empty string
		if value == "" {
			continue
		}

		// set value
		values[keyValue] = stringHandler(value)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// close file
	defer file.Close()

	return values, nil
}

func stringHandler(s string) string {
	// if first letter and last letter has double quotes
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	return s
}

func valueHandler(s string) (result string) {
	s = stringHandler(s)
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		// concat
		result += string(s[i])

		// if i >= sLen then break
		if i >= sLen {
			break
		}

		// if next letter is space and next one is # then break
		if i+2 < sLen {
			if string(s[i+1]) == " " && string(s[i+2]) == "#" {
				break
			}
		}
	}

	return result
}
