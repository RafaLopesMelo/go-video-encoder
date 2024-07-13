package env

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Load(filename string) error {
    _, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)

	file, err := os.Open(basepath + "/" +filename)

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
