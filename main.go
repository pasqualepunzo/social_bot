package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Twitter user-auth requests with an Access Token (token credential)
// func main() {

// 	godotenv.Load(".env")

// 	// read credentials from environment variables
// 	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
// 	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
// 	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
// 	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
// 	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
// 		panic("Missing required environment variable")
// 	}

// 	config := oauth1.NewConfig(consumerKey, consumerSecret)
// 	token := oauth1.NewToken(accessToken, accessSecret)

// 	// httpClient will automatically authorize http.Request's
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	path := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
// 	path = "https://api.twitter.com/2/users/me" // OK
// 	// path = "https://api.twitter.com/2/users/by/username/:__pasq__"

// 	resp, _ := httpClient.Get(path)
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Printf("Raw Response Body:\n%v\n", string(body))

// 	// Nicer: Pass OAuth1 client to go-twitter API
// 	api := twitter.NewClient(httpClient)
// 	tweets, _, _ := api.Timelines.HomeTimeline(nil)
// 	fmt.Printf("User's HOME TIMELINE:\n%+v\n", tweets)
// }

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	godotenv.Load(".env")

	creds := Credentials{
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	fmt.Printf("%+v\n", client)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "Golang",
	})

	if err != nil {
		log.Print(err)
	}

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", search)

}

func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
