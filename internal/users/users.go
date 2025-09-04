package users

import (
	"context"
	"fmt"
	"log"

	"github.com/htanmo/hackernews/internal/database"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create(ctx context.Context) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = database.Pool.Exec(ctx, "INSERT INTO Users (Username, Password) VALUES($1, $2);", user.Username, hashedPassword)

	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIDByUsername(ctx context.Context, username string) (int, error) {
	row := database.Pool.QueryRow(ctx, "SELECT ID FROM Users WHERE Username = $1;", username)

	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("user not found: %w", err)
		}
		return 0, fmt.Errorf("failed to get user ID by username: %w", err)
	}

	log.Printf("Found user ID %d for username %s", id, username)
	return id, nil
}

func (user *User) Authenticate(ctx context.Context) bool {
	var hashedPassword string
	err := database.Pool.QueryRow(
		ctx,
		"SELECT Password FROM Users WHERE Username = $1;",
		user.Username,
	).Scan(&hashedPassword)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}
