package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if getLink("https://campus-booster.net/") {
		sendTelegramNotification("Campus Booster", "https://campus-booster.net/")
	}

	if getLink("https://canvas.supinfo.com/") {
		sendTelegramNotification("Canvas", "https://canvas.supinfo.com/")
	}
}

func getLink(url string) bool {
	_, err := http.Get(url)

	if err != nil {
		return false
	}

	return false
}

func sendTelegramNotification(website string, websiteUrl string) {
	url := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendMessage"

	reqBody, err := json.Marshal(map[string]string{
		"chat_id": os.Getenv("TELEGRAM_CHANNEL"),
		"text":    website + " is up ! You can check it : " + websiteUrl,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	resp, err := http.Post(url,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		print(err)
	}
	log.Printf("Telegram response : \n%s\n", string(body))
}
