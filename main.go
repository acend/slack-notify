package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	EnvSlackWebhook = "SLACK_WEBHOOK"
	EnvSlackMessage = "SLACK_MESSAGE"
)

type Webhook struct {
	Text string `json:"text,omitempty"`
}

func main() {
	endpoint := os.Getenv(EnvSlackWebhook)
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}
	text := os.Getenv(EnvSlackMessage)
	if text == "" {
		g
		fmt.Fprintln(os.Stderr, "Message is required")
		os.Exit(1)
	}

	msg := Webhook{
		Text: text,
	}

	if err := send(endpoint, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(2)
	}
}

func send(endpoint string, msg Webhook) error {
	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("Error on message: %s\n", res.Status)
	}
	fmt.Println(res.Status)
	return nil
}
