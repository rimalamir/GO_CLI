package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "YouTube video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "A", "YouTube video id")
	addTitle := addCmd.String("title", "", "YouTube video title")
	addURL := addCmd.String("url", "", "YouTube video url")
	addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
	addDesc := addCmd.String("desc", "", "YouTube video description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "get":
		handleGet(getCmd, getAll, getID)
	case "add":
		handleAdd(addCmd, addID, addTitle, addURL, addImageUrl, addDesc)
	default:

	}
}

func handleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description\n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description\n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
	}
}

func validateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageURL *string, desc *string) {

	if *id == "" || *title == "" || *url == "" || *imageURL == "" || *desc == "" {
		fmt.Print("All fields are required for adding a video\n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func handleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageURL *string, desc *string) {
	addCmd.Parse(os.Args[2:])
	validateVideo(addCmd, id, title, url, imageURL, desc)

	video := video{
		Id:          *id,
		Title:       *title,
		Url:         *url,
		Imageurl:    *imageURL,
		Description: *desc,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
