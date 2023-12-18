package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	PacksRequired []int `json:"packsRequired"`
}

func main() {
	// Lambda will call the Handler function
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	itemsOrderedStr, ok := request.QueryStringParameters["itemsOrdered"]

	if !ok {
		log.Println("Missing 'itemsOrdered' query parameter")
		return events.APIGatewayProxyResponse{}, fmt.Errorf("Missing 'itemsOrdered' query parameter")
	}

	itemsOrdered, _ := strconv.Atoi(itemsOrderedStr)

	packs, err := getPackSizes()

	if err != nil {
		log.Print("Error encountered while trying to get pack sizes:", err)
		return events.APIGatewayProxyResponse{}, err
	}

	sortedPackSizes := sortPackSizes(packs.PackSizes)
	possiblePackSizes := possiblePackCombinations(sortedPackSizes)
	combinations := findCombinations(sortedPackSizes, possiblePackSizes)

	packsRequired := handlePacks(itemsOrdered, sortedPackSizes, possiblePackSizes, combinations)

	responseData := Response{
		PacksRequired: packsRequired,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		log.Print("Error marshaling JSON:", err)
		return events.APIGatewayProxyResponse{}, fmt.Errorf("Error creating JSON response")
	}

	headers := map[string]string{
		"Access-Control-Allow-Origin":  "localhost:3000, https://pack-optimisation-d42ubt7qr-toms-projects-41b46903.vercel.app/, *",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "OPTIONS, GET, POST, ANY",
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}
