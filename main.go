package main

import "github.com/lethanhsonthd/EOS_API_create_account/routers"

func main() {
	router := routers.GetRoutes()
	router.Run(":8000")
}
