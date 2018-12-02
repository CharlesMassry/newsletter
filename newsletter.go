package newsletter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var netClient = &http.Client{Timeout: time.Second * 10}

func newsletter(languages []string) {
	for _, language := range languages {
		subreddit := getSubredditForLanguage(language)
		url := getURLForSubreddit(subreddit)
		redditResponse := RedditResponse{}
		getJSON(url, &redditResponse)
		fmt.Println(redditResponse.Data)
	}
}

func getJSON(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "golang-tutorial")
	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getURLForSubreddit(subreddit string) string {
	return "https://www.reddit.com/r/" + subreddit + ".json"
}

func getSubredditForLanguage(language string) string {
	languages := map[string]string{
		"go":     "golang",
		"ruby":   "ruby",
		"golang": "golang",
		"elixir": "elixir",
	}

	subreddit, ok := languages[language]

	if ok != true {
		log.Fatal("Language Not Found")
	}

	return subreddit
}
