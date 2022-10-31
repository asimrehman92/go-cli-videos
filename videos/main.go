package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "Youtube video Id")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addId := addCmd.String("id", "", "Youtube Video Id")
	addTitle := addCmd.String("title", "", "Youtube Video Title")
	addUrl := addCmd.String("url", "", "Youtube Video Url")
	addImageUrl := addCmd.String("imageurl", "", "Youtube Video Image Url")
	addDesc := addCmd.String("desc", "", "Youtube Video Desc")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getId)
	case "add":
		HandleAdd(addCmd, addId, addTitle, addImageUrl, addUrl, addDesc)
	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \t \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageUrl, desc)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *desc,
		Imageurl:    *imageUrl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, desc *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *desc == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}
