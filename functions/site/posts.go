package site

import (
	"time"
)

var posts map[string]Post

// Post contains the list of posts
type Post struct {
	Author  User
	Name    string
	Content string
	Date    time.Time
}

// NewPost creates a new post
func NewPost(author User, name string, content string) Post {
	return Post{
		Author:  author,
		Name:    name,
		Content: content,
	}
}

// AddPost add the post to the post list. its identified by the post name
func AddPost(p Post) {
	posts[p.Name] = p
}

// GetUsersPosts get the list of the posts made by the user
func GetUsersPosts(username string) []Post {
	var postList []Post
	for _, v := range posts {
		if v.Author.Username == username {
			postList = append(postList, v)
		}
	}
	return postList
}
