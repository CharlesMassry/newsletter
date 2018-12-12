package main

// RedditResponse holds reddit response data
type RedditResponse struct {
	Data RedditResponseData `json:"data"`
}

// RedditResponseData holds a collection of children which is a wrapper around a reddit post
type RedditResponseData struct {
	Posts []RedditPostData `json:"children"`
}

// RedditPostData is a wrapper around a RedditPost
type RedditPostData struct {
	Kind string     `json:"kind"`
	Data RedditPost `json:"data"`
}

// RedditPost holds all the data associated to a reddit post
type RedditPost struct {
	Title     string `json:"title"`
	Subreddit string `json:"subreddit"`
	URL       string `json:"url"`
	IsSelf    bool   `json:"is_self"`
}

// ConvertToPost converts a reddit post to a generic Post
func (redditPost RedditPost) ConvertToPost() Post {
	return Post{Title: redditPost.Title, URL: redditPost.URL}
}
