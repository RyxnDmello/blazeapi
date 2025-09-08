package main

import "github.com/ryxndmello/flame/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic("An Unexpected Error Has Occurred")
	}
}
