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

	brokerPort, err := strconv.Atoi(getEnv("KAFKA_PORT", "9092"))

	if err != nil {
		log.Fatalln(err)
	}

	wPort, err := strconv.Atoi(getEnv("SERVICE_PORT", "15643"))

	if err != nil {
		log.Fatalln(err)
	}

	brokerAddr := getEnv("KAFKA_HOST", "127.0.0.1")
	topic := getEnv("KAFKA_TOPIC", "change")
	partition, err := strconv.Atoi(getEnv("KAFKA_TOPIC_PARTITION", "0"))


	if err != nil {
		log.Fatalln(err)
	}
	
	s := registrar.NewServer(wPort, brokerAddr, brokerPort, topic, partition)

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