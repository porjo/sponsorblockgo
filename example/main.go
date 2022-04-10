package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	sb "github.com/porjo/sponsorblockgo"
)

func main() {

	client, err := sb.NewClient("https://sponsor.ajay.app/api")
	if err != nil {
		log.Fatalf("Error getting client: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Replace this with a valid Youtube VideoID
	ssParams := sb.GetskipsegmentsParams{VideoID: "xxxxxyyyy"}

	resp, err := client.Getskipsegments(ctx, &ssParams)
	if err != nil {
		log.Fatalf("Error getting segments: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %s\n", err)
	}

	fmt.Printf("response: %s\n", body)

}
