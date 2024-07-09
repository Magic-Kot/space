package main

import (
	tgClient "4.space/clients/telegram"
	event_consumer "4.space/consumer/event-consumer"
	"4.space/events/telegram"
	"4.space/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	) // bot -tg-bot-token 'my token'

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
