package models

import (
	"database/sql"
	"fmt"
	"log"
)


type BlogUser struct{
	Id string
	Name string `json:"name"`
	PassWord string `json:"password"`
	Email string `json:"email"`
	Bio string	`json:"bio"`
	State bool `json:"state"`
	Followers uint32 `json:"follower"`
}

type DBModel struct {
	Db *sql.DB
}

type IUsers interface {
	SignUpUser(user *BlogUser) error
	QueryUser(user *BlogUser,email string) error
	GetAllUsers(user BlogUser) (users []BlogUser, err error)
	QueryEmail(email string) (user BlogUser, err error)
}

func (d *DBModel)SignUpUser(user *BlogUser) error{
	stmt := `INSERT INTO user (userid,name, password, email, bio, state, follower) VALUES(?, ?, ?, ?, ?, ?, ?)`
	sql, err := d.Db.Prepare(stmt)

	if err != nil {
		return err
	}
	_ ,err = sql.Exec(user.Id, user.Name, user.PassWord, user.Email, user.Bio, user.State, user.State)
	return err
}
func (d *DBModel)QueryUser(user *BlogUser,email string) error {
	row := d.Db.QueryRow(`SELECT email FROM user WHERE email = ?`, email)
	if err := row.Scan(&user.Email); err != nil {
		return err
	}

	return nil
}

func (d *DBModel) QueryEmail(email string) (user BlogUser, err error) {
	log.Println(" db working first")
	row := d.Db.QueryRow(`SELECT * FROM user WHERE  email = ?`,email)
	log.Println(email," db working after scan")
	err = row.Scan(&user.Id, &user.Name, &user.PassWord, &user.Email, &user.Bio, &user.State, &user.Followers)

	if err != nil {
		return user,err
	}
	return user, nil
}

func (d *DBModel)GetAllUsers(user BlogUser) (users []BlogUser, err error) {
	// An albums slice to hold data from returned rows.

	rows, err := d.Db.Query(`SELECT * FROM user`)
	if err != nil {
		return nil, fmt.Errorf("blogUser %q", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {

		if err := rows.Scan(&user.Id, &user.Name, &user.PassWord, &user.Email, &user.Bio, &user.State, &user.Followers); err != nil {
			return nil, fmt.Errorf("blogUser %q: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("blogUser %q", err)
	}
	return users, nil
}


//func (d DBModel) GetAllPosts() ([]pkg.Post, error) {
//	// An albums slice to hold data from returned rows.
//	var posts []pkg.Post
//	rows, err := db.Query(`SELECT * FROM post`)
//	if err != nil {
//		return nil, fmt.Errorf("blogUser %q", err)
//	}
//	defer rows.Close()
//	// Loop through rows, using Scan to assign column data to struct fields.
//	for rows.Next() {
//		var post pkg.Post
//		if err := rows.Scan(&post.Id, &post.Title, &post.Message, &post.Like); err != nil {
//			return nil, fmt.Errorf("blogUser %v", err)
//		}
//		posts = append(posts, post)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("blogUser %q", err)
//	}
//	return posts, nil
//}
//
//func GetAllComments() ([]pkg.Comment, error) {
//	// An albums slice to hold data from returned rows.
//	var comments []pkg.Comment
//	rows, err := db.Query(`SELECT * FROM comments`)
//	if err != nil {
//		return nil, fmt.Errorf("blogUser %q", err)
//	}
//	defer rows.Close()
//	// Loop through rows, using Scan to assign column data to struct fields.
//	for rows.Next() {
//		var comment pkg.Comment
//		if err := rows.Scan(&comment.CommentId, &comment.Id, &comment.Blogger, &comment.Message); err != nil {
//			return nil, fmt.Errorf("blogUser %v", err)
//		}
//		comments = append(comments, comment)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("blogUser %q", err)
//	}
//	return comments, nil
//}

//func getAllfollower () []pkg.BlogUser{
//	var comments []pkg.Comment
//	rows, err := db.Query(`SELECT * FROM comments`)
//	if err != nil {
//		return fmt.Errorf("blogUser %q", err)
//	}
//	defer rows.Close()
//	// Loop through rows, using Scan to assign column data to struct fields.
//	for rows.Next() {
//		var comment pkg.Comment
//		if err := rows.Scan(&comment.CommentId, &comment.Id, &comment.Blogger, &comment.Message); err != nil {
//			return  fmt.Errorf("blogUser %v", err)
//		}
//		comments = append(comments, comment)
//	}
//	if err := rows.Err(); err != nil {
//		return  fmt.Errorf("blogUser %q", err)
//	}
//	return comments, nil
//}




//

//
//type Notification interface{
//
//}


//func (b *BlogUser) Login(){
//
//}
//
//func (p *Post) View(){
//
//}
//
//func (b *BlogUser) LikeNotification(){
//
//}
//
//func (p *Post) SharePost(){
//}