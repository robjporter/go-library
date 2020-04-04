package main

import (
	"fmt"

	"github.com/robjporter/go-library/xwallpaper"
)

func main() {
	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current wallpaper:", background)
	//wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
}
