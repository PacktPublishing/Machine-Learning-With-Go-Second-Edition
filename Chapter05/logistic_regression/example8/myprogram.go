package main

import (
	"github.com/zupzup/ml-in-go-examples/logisticregression"
	"log"
)

func main() {
	if err := logisticregression.Run(); err != nil {
		log.Fatal(err)
	}
}