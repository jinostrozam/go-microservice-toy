package main

import "fmt"

func main() {

	videos := getVideos()

	for i, _ := range videos {
		videos[i].Description = videos[i].Description + "\n Remember to like and suscribe"
	}

	fmt.Println(videos)

	saveVideos(videos)

}
