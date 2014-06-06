package main

import (
  "fmt"
  "os"
  "flag"
  "time"
  "github.com/rakyll/globalconf"
  "github.com/ChimeraCoder/anaconda"
)

func main() {

  var (
    consumer_key    = flag.String("consumerkey", "", "the app consumer key")
    consumer_secret = flag.String("consumersecret", "", "the app consumer secret")
    access_token    = flag.String("accesstoken", "", "app access token")
    access_secret   = flag.String("accesstokensecret", "", "app access token secret")
    show_help       = flag.Bool("h", false, "show this dialog")
    graphite        = flag.Bool("graphite", false, "print stats in Graphite format")
  )

  conf, _ := globalconf.New("twanity")
  conf.ParseAll()

  if *show_help {
    flag.PrintDefaults()
    os.Exit(1)
  }


	flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "usage: twanity  username\n")
    flag.PrintDefaults()
    os.Exit(1)
  }

	flag.Parse()

	username := flag.Arg(0)

	if username == "" {
		flag.Usage()
	}

  anaconda.SetConsumerKey(*consumer_key)
  anaconda.SetConsumerSecret(*consumer_secret)
  api := anaconda.NewTwitterApi(*access_token, *access_secret)

  result, err := api.GetUsersShow(username, nil)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to retrieve user details\n")
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  if *graphite {
    timestamp := time.Now().Unix()
    fmt.Printf("twanity.%s.following %d %d\n", username, result.FriendsCount, timestamp)
    fmt.Printf("twanity.%s.followers %d %d\n", username, result.FollowersCount, timestamp)
    fmt.Printf("twanity.%s.tweets %d %d\n", username, result.StatusesCount, timestamp)

  } else {
    fmt.Printf("Getting twitter stats for %s...\n", username)
    fmt.Printf("Following: %d\n", result.FriendsCount)
    fmt.Printf("Followers: %d\n", result.FollowersCount)
    fmt.Printf("Tweets: %d\n", result.StatusesCount)
  }
}
