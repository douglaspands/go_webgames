package main

import (
	webgames "webgames/app"
)

func main() {
	app := webgames.CreateApp()
	app.Run(":5000")
}
