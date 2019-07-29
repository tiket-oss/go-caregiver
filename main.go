package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	command := os.Args[1:]
	configPath := os.Getenv("CONFIG")

	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(configRaw, &config); err != nil {
		log.Fatal(err)
	}

	for key, value := range config["service"].(map[string]interface{}) {
		value, err := toString(value)
		if err != nil {
			log.Fatal(err)
		}

		os.Setenv(key, value)
	}

	cmd := exec.Command(command[0], command[1:]...)
	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", buf.String())
}

func toString(i interface{}) (string, error) {
	switch v := i.(type) {
	case string:
		return v, nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	default:
		msg := fmt.Sprintf("%T type unsupported\n", v)
		return "", errors.New(msg)
	}
}
