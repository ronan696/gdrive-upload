package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/drive/v3"
)

func main() {

	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	fileName := "<FileName>"
	mimeType := "<MimeType>"
	parentId := "<ParentID>"

	d := &drive.File{
		Name:     fileName,
		MimeType: mimeType,
		Parents:  []string{parentId},
	 }
	
	content, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}	

	file, err := srv.Files.Create(d).Media(content).Do()
  
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}

	if file != nil	{
		fmt.Println("Uploaded file: " + file.Id)
	}
	 
}