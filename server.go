package main

import (
	"github.com/labstack/gommon/log"
	"go-ms-session/config"
	"go-ms-session/context"
	"go-ms-session/repository"
	"go-ms-session/router"
	"os"
)

func main() {
	s := context.CreateServer(
		"go-upload-ImgProfile",
		config.Config{
			Application: "application",
			Env:         os.Getenv("ENV"),
			Resource:    "resource",
			ParenPath:   "",
		},
		nil,
	)

	if client, err := repository.CreateMongoClient(&s.Config); err != nil {
		log.Error(err)
		return
	} else {
		s.MongoRepository = repository.NewMongoRepository(client, &s.Config)
		r := router.NewRoutes(s.Serve, s.MongoRepository)
		r.InitRoute()
		log.Fatal(s.Serve.Start(s.GetPort(s.Config.Server.Port)))
	}
}
