package discord_over_sms

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func PostWebhook() int {
	var bod io.Reader
	bod = strings.NewReader("{\"username\": \"test\", \"content\": \"hello\"}")
	resp, err := http.Post("https://discord.com/api/webhooks/921964797673742376/T4CW-_-BFANBKKTqcPc3vvmBqymKWZcvDFGWNZkL81JK-M7C4wtlOHsmBk7Koch8lJmb",
		"application/json",
		bod)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	code := resp.StatusCode
	resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	return code
}
