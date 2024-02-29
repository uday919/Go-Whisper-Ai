package main

import (
	"fmt"
	"log"
	"os"

	"github.com/uday919/go-whisper-ai/api/whisper"
)

func main() {
	client := whisper.NewClient(whisper.WithKey(os.Getenv("OPENAI_API_KEY")))

	response, err := client.TranscribeFile("jus.mp3")
	if err != nil {
		log.Fatalf("Error transcribing the file: %v", err)
	}
	fmt.Printf("Transcription: %s\n", response.Text)
}
