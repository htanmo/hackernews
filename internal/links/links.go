package links

import (
	"context"
	"log"

	"github.com/htanmo/hackernews/internal/database"
	"github.com/htanmo/hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link *Link) Save(ctx context.Context) int64 {
	var id int64
	err := database.Pool.QueryRow(ctx, "INSERT INTO Links (Title, Address, UserID) VALUES($1, $2, $3) RETURNING ID;", link.Title, link.Address, link.User.ID).Scan(&id)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll(ctx context.Context) []Link {
	rows, err := database.Pool.Query(
		ctx,
		"SELECT L.ID, L.Title, L.Address, L.UserID, U.Username from Links L INNER JOIN Users U on L.UserID = U.ID;",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
