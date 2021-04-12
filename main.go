package main

import (
	"kindle_clipping_exporter/kindle"
	"kindle_clipping_exporter/kindle_emailer"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting...")
	creds := loadEnv()
	kindleHandler := kindle.InitHandler()
	if creds != nil {
		for _, device := range kindleHandler.Kindles {
			kindle_emailer.SendEmail(device, creds)
		}
	} else {
		log.Println(".env file was not loaded. Skipping emailing.")
	}
}

func loadEnv() *kindle_emailer.Credentials {
	var myEnv map[string]string
	creds := &kindle_emailer.Credentials{}
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Printf("There was an error loading the env file: %v\n", err)
		return nil
	}

	creds.FromEmail = myEnv["FROM_EMAIL"]
	creds.FromEmailPassword = myEnv["FROM_EMAIL_PASSWORD"]
	creds.ToEmail = myEnv["TO_EMAIL"]
	log.Printf(".env file loaded. Email will be sent to %s\n", creds.ToEmail)
	return creds
}
