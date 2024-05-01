package main

import (
	"context"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type LambdaResponse struct {
	ID   string `json:"id"`
	Code int    `json:"code"`
	Body []byte `json:"body"`
}

func ProcessRequest(ctx context.Context, event *LambdaEvent) (*LambdaResponse, error) {
	log := slog.With("id", event.ID)
	log.Info("Processing event", "event", event)
	client := &http.Client{}
	request, err := http.NewRequest("GET", event.URL, nil)
	if err != nil {
		log.Error("unable to init a new request", "err", err)
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Error("unable to process request", "err", err)
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("unable to read response body from request", "err", err)
		return nil, err
	}

	lambdaResponse := LambdaResponse{
		ID:   event.ID,
		Code: response.StatusCode,
		Body: bodyBytes,
	}
	return &lambdaResponse, nil
	// return nil, fmt.Errorf("err: invalid processing for event %+v", event)
}

func main() {
	lambda.Start(ProcessRequest)
}
