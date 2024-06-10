package env

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Load(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		err = errors.New("Could not find .env file at the project root")
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.SplitN(text, "=", 2)

		key := parts[0]
		value := strings.ReplaceAll(parts[1], `"`, "")

		os.Setenv(key, value)
	}

	return nil
}

func Get(key string) string {
	return os.Getenv(key)
}
