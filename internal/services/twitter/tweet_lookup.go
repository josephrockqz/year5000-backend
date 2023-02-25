package tweet_lookup

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
 )

 func fetchTweets() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	
	id := "your-user-id";
	url := "https://api.twitter.com/2/users/" + id + "/blocking"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer")
	resp, err := client.Do(req)
	resp, err := client.Get(url)
 }
