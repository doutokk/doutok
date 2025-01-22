package main

import (
	auth "douyin-commerce/auth_service/kitex_gen/auth/authservice"
	"log"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
