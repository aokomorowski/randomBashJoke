package main

import (
	"net/http"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func main() {

	lambda.Start(getRandomJokeHandler)
}
