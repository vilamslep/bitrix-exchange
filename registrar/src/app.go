package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/vi-la-muerto/bx24-service/http/registrar"
)

func main() {

	brokerAddr := getEnv("KAFKA_HOST", "localhost")

	brokerPort, err := strconv.Atoi(getEnv("KAFKA_PORT", "9092"))
	if err != nil {
		log.Fatalln(err)
	}
	topic := getEnv("KAFKA_TOPIC", "tocrm")

	partition, err := strconv.Atoi(getEnv("KAFKA_TOPIC_PARTITION", "0"))

	if err != nil {
		log.Fatalln(err)
	}

	messageKey := getEnv("KAFKA_MESSAGE_KEY", "")

	wPort, err := strconv.Atoi(getEnv("SERVICE_PORT", "8080"))

	if err != nil {
		log.Fatalln(err)
	}

	method := getEnv("SERVICE_METHOD", "")

	checkInput := true
	if v := getEnv("SERVICE_CHECK_INPUT", "0"); v == "1" {
		checkInput = false
	}

	s := registrar.NewServer(wPort, brokerAddr, brokerPort, topic, partition, method, checkInput, messageKey)

	//new chanel and exec subscription for handing
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go s.Run()

	//wait signal
	<-done

	s.Close()
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
