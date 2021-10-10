package models

import (
	"database/sql"
	"fmt"
)

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
		d.Db.Exec(`INSERT INTO post (id, userid,title, message,likes) VALUES (?, ?, ?,?,?)`,
			blg.Id, blg.Userid, blg.Title,blg.Message,blg.Like)
	if err != nil {
		return err
	}
	return nil
}
func (d DBModel) SelectSinglePost (id string) (Post, error) {
	var post Post

	row := d.Db.QueryRow(`SELECT * FROM post WHERE id, = ?`, id)
	if err := row.Scan(&post.Id, &post.Title, &post.Message, &post.Like); err != nil {
		if err == sql.ErrNoRows {
			return post, fmt.Errorf("blogId %d: no such blogPost", id)
		}
		return post, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return post, nil
}
func(d DBModel) AllPost() (posts []Post, err error) {
	// An albums slice to hold data from returned rows.
	rows, err := d.Db.Query(`SELECT id, userid,title, message,likes  FROM post`)
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
		return nil, fmt.Errorf("Blog Post %q", err)
	}
	return posts, nil
}
