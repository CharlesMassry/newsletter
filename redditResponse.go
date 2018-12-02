package main

// RedditResponse holds reddit response data
type RedditResponse struct {
	Data RedditResponseData
}

// RedditResponseData holds a collection of children which is a wrapper around a reddit post
type RedditResponseData struct {
	Children []RedditChildData
}

// RedditChildData is a wrapper around a RedditPost
type RedditChildData struct {
	Data RedditPost
}

// RedditPost holds all the data associated to a reddit post
type RedditPost struct {
	title string
}
