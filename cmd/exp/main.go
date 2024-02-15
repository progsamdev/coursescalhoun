package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/progsamdev/coursescalhoun/models"
)

// type Tweet struct {
// 	id          string
// 	userId      string
// 	postMessage string
// }

// func getNewTweet(t *User) *Tweet {
// 	return &Tweet{
// 		userId:      t.id,
// 		postMessage: "New post on X",
// 	}
// }

// type Like struct {
// 	userId  string
// 	tweetId string
// }

// func getNewLike(t *Tweet, u *User) *Like {
// 	return &Like{
// 		userId:  u.id,
// 		tweetId: t.id,
// 	}
// }

func main() {
	cfg := models.DefaultPostgresConfig()
	conn, err := models.Open(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	fmt.Println("Connected!")

	userSer := models.UserService{DB: conn}

	user, err := userSer.Create("samuel.msbr2@gmail.com", "123aaa")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	fmt.Println(user)

	/*// Begin a transaction
	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	u := getNewUser()
	row := tx.QueryRow(
		`INSERT INTO users(age, first_name, last_name, email)
			VALUES ($1, $2, $3, $4) RETURNING id`, u.age, u.firstName, u.lastName, u.email,
	)
	err = row.Scan((&u.id))
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		t := getNewTweet(u)
		row := tx.QueryRow(
			`INSERT INTO tweets(user_id, post_message)
				VALUES ($1, $2 ) RETURNING id`, t.userId, t.postMessage)
		err = row.Scan((&t.id))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		for j := 0; j < i; j++ {
			l := getNewLike(t, u)
			_, err = tx.Exec(
				`INSERT INTO likes( user_id, tweet_id)
					VALUES ($1, $2)`, l.userId, l.tweetId,
			)
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	} */
}
