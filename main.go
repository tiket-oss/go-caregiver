package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"gopkg.in/natefinch/lumberjack.v2"
)

type serviceConfig struct {
	Service map[string]interface{} `json:"service"`
	Log     *logConfig             `json:"log"`
}

type logConfig struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsizemb"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxagedays"`
	Compress   bool   `json:"compress"`
}

func main() {
	command := os.Args[1:]
	configPath := os.Getenv("CONFIG")

	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config serviceConfig
	if err := json.Unmarshal(configRaw, &config); err != nil {
		log.Fatal(err)
	}

	for key, value := range config.Service {
		value, err := toString(value)
		if err != nil {
			log.Fatal(err)
		}

		os.Setenv(key, value)
	}

	cmd := exec.Command(command[0], command[1:]...)

	var writer io.Writer
	if config.Log != nil {
		writer = &lumberjack.Logger{
			Filename:   config.Log.Filename,
			MaxSize:    config.Log.MaxSize,
			MaxBackups: config.Log.MaxBackups,
			MaxAge:     config.Log.MaxAge,
			Compress:   config.Log.Compress,
		}
	} else {
		writer = os.Stdout // stdout main process
	}

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
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
