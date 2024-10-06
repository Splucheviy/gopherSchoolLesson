package main

import (
	"log"

	"github.com/Splucheviy/gopherSchoolLesson/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
