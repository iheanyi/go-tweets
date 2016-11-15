package main

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
)

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	s := api.PublicStreamSample(v)

	for t := range s.C {
		log.Print("In the range function!")
		switch v := t.(type) {
		case anaconda.Tweet:
			log.Printf("%v", v)
		}
	}
}
