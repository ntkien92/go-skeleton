package main

import "blog-api/cmd"

func main() {
	server := cmd.ApiServer{}
	server.Run()
}
