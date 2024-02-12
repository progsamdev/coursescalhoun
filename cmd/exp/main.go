package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) toString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

type User struct {
	id        string
	age       uint8
	firstName string
	lastName  string
	email     string
}

func getNewUser() *User {
	return &User{
		firstName: "Test",
		lastName:  "TestLastName",
		age:       23,
		email:     "tes2t22@demo33.com",
	}
}

type Tweet struct {
	id          string
	userId      string
	postMessage string
}

func getNewTweet(t *User) *Tweet {
	return &Tweet{
		userId:      t.id,
		postMessage: "New post on X",
	}
}

type Like struct {
	userId  string
	tweetId string
}

func getNewLike(t *Tweet, u *User) *Like {
	return &Like{
		userId:  u.id,
		tweetId: t.id,
	}
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	conn, err := sql.Open("pgx", cfg.toString())
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

	// Begin a transaction
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
	}
}
