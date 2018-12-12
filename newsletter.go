package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var netClient = &http.Client{Timeout: time.Second * 10}

func newsletter(languages []string) string {
	numberOfChannels := len(languages)
	channels := make(PostChannel, numberOfChannels)
	for _, language := range languages {
		go worker(language, channels)
	}

	posts := collectPosts(numberOfChannels, channels)

	html := generateHTML(posts)

	// file, err := os.Create("result.html")

	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// fmt.Fprintf(file, html)

	return html
}

func generateHTML(posts []Post) string {
	html := "<html><head></head><body>"
	for _, post := range posts {
		html += post.toHtml()
	}

	html += "</body></html>"

	return html
}

func collectPosts(numberOfChannels int, channels PostChannel) []Post {
	var posts []Post
	for x := 0; x < numberOfChannels; x++ {
		for _, post := range <-channels {
			posts = append(posts, post)
		}
	}

	return posts
}

func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func worker(language string, channels PostChannel) {
	channels <- getDataFromReddit(language)
}

func getDataFromReddit(language string) []Post {
	subreddit := getSubredditForLanguage(language)
	url := getURLForSubreddit(subreddit)
	redditResponse := RedditResponse{}
	getJSON(url, &redditResponse)
	var posts []Post

	for _, post := range redditResponse.Data.Posts {
		data := post.Data
		if !data.IsSelf {
			posts = append(posts, data.ConvertToPost())
		}
	}

	return posts
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
