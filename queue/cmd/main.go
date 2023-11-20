package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func JetStreamInit() (nats.JetStreamContext, error) {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	return js, nil
}

const (
	StreamName     = "REVIEWS"
	StreamSubjects = "REVIEWS.*"
)

func CreateStream(jetStream nats.JetStreamContext) error {
	stream, err := jetStream.StreamInfo(StreamName)

	// stream not found, create it
	if stream == nil {
		log.Printf("Creating stream: %s\n", StreamName)

		_, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     StreamName,
			Subjects: []string{StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type Review struct {
	Id      string `json:"_id"`
	Author  string `json:"author"`
	Store   string `json:"store"`
	Text    string `json:"text"`
	Rating  int    `json:"rating"`
	Created string `json:"created"`
}
