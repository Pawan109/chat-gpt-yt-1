package main

import (
	"context" //request scoped values ke liye use krte hai and also for data-passing & timeout // request scoped values matlb jo request var hai uska ek naya instance bana deta hai
	//so a request scoped var ka instance does not exist in the context of another request.
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3" //gpt-3 API client enabling go/golang programs to interact with the gpt3 APIs
	"github.com/joho/godotenv"          //.env files ko load kra ne ke liye library
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatalln("missing API Key")
	}

	ctx := context.Background() //returns an empty Context.- never cancelled, has no values , had no deadline.
	//it initialises and tests and can be used as the top-level{first context} Context for incoming requests.

	client := gpt3.NewClient(apiKey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"the first thing you should know about Golang is "}, //we are giving gpt3 a prompt, and it'll complete it and give it back to us .
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text) //just printing out the response i got from the gpt3 api
}
