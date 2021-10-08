package models

import "fmt"

type Post struct{
	Id string `json:"id"`
	Userid string `json:"blogger"`
	Title string `json:"title"`
	Message string `json:"message"`
	Like int `json:"like"`
	Comments []Comment
}

type Comment struct{
	CommentId string
	PostId string
	Blogger string
	Message string
}
type Iblog interface {
	AddPostToDatabase(blg *Post) error
	AllPost() (posts []Post, err error)
}
func (d DBModel) AddPostToDatabase(blg *Post) error{
	_, err :=
		d.Db.Exec(`INSERT INTO blogPosts (id, blogger,title, message,likes) VALUES (?, ?, ?,?,?)`,
			blg.Id, blg.Userid, blg.Title,blg.Message,blg.Like)
	if err != nil {
		return err
	}
	return nil
}

func(d DBModel) AllPost() (posts []Post, err error) {
	// An albums slice to hold data from returned rows.
	rows, err := d.Db.Query(`SELECT id, blogger,title, message  FROM post`)
	if err != nil {
		return nil, fmt.Errorf("blog post %q", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Message, &post.Like); err != nil {
			return nil, fmt.Errorf("blog post %q: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q", err)
	}
	return posts, nil
}
