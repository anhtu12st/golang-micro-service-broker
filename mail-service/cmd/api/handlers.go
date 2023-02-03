package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}
	log.Println("SendMail() start:")

	var requestPayload mailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	log.Printf("SendSMTPMessage - %v", requestPayload)

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJSON(w, err)
		log.Printf("Error: SendSMTPMessage - %v", err.Error())
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Sent to " + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
