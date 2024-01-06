package main

import "github.com/ArshpreetS/Admin/utils"

func main() {
	r := utils.GetRoutes()
	r.Run(":9001")
}
