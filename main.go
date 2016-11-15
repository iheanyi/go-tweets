package main

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
)

type Streamer struct {
	api *anaconda.TwitterApi
}

type TweetEvent struct {
	filter string
	tweet  *anaconda.Tweet
}

func (stream *Streamer) doStream(i int, ch chan interface{}) {
	v := url.Values{}
	s := stream.api.PublicStreamSample(v)

	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			//if v.HasCoordinates() {
			log.Printf("Pushing in stream :%d", i)
			log.Printf("Here's the User: %v", v.User.ScreenName)
			log.Printf("Here's the Tweet Text: %v", v.Text)
			log.Printf("Here's the Coordinates: %v", v.Coordinates)
			log.Printf("Here's the user's location: %v", v.User.Location)
			ch <- v
			//}
		}
	}
}

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	s := Streamer{api: api}
	ch := make(chan interface{})
	go s.doStream(1, ch)
	go s.doStream(2, ch)

	for _ = range ch {
		log.Print("Printing stuff!")
	}
}
