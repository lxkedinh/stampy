package env

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func Load(path string) error {

	envFile, err := os.Open(path)
	if err != nil {
		log.Panicf("Could not load .env file\n%v", err)
		return errors.New("Could not load .env file")
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envVar := strings.Split(scanner.Text(), "=")
		os.Setenv(envVar[0], envVar[1])
	}

	if err = scanner.Err(); err != nil {
		log.Panicf("Error during scanning env file\n%v", err)
		return errors.New("Error during scanning env file")
	}

	return nil
}
