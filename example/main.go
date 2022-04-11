package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	sb "github.com/porjo/sponsorblockgo"
)

func main() {

	videoID := flag.String("videoID", "", "Youtube video ID")
	flag.Parse()

	if *videoID == "" {
		log.Printf("videoID required\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	client, err := sb.NewClientWithResponses("https://sponsor.ajay.app/api")
	if err != nil {
		log.Fatalf("Error getting client: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Replace this with a valid Youtube VideoID
	ssParams := sb.GetskipsegmentsParams{VideoID: sb.RequiredVideoID(*videoID)}

	resp, err := client.GetskipsegmentsWithResponse(ctx, &ssParams)
	if err != nil {
		log.Fatalf("Error getting segments: %s\n", err)
	}

	fmt.Printf("Segments for videoID %s\n", *videoID)
	fmt.Printf("%20s %20s\n", "start", "stop")
	fmt.Println("---------------------------------------------")
	for _, s := range *resp.JSON200 {
		fmt.Printf("%20f %20f\n", (*s.Segment)[0], (*s.Segment)[1])
	}

}
