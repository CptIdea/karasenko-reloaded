package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"karasenko-reloaded/collectors/friends"
	"karasenko-reloaded/collectors/groups"
	"karasenko-reloaded/controller"
	"karasenko-reloaded/store"

	"github.com/SevereCloud/vksdk/v2/api"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func main() {
	vk := api.NewVK(os.Getenv("TOKEN"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	pg := store.NewStore(conn)

	friendsCollector := friends.NewCollector(vk, pg)
	groupsCollector := groups.NewCollector(pg, vk)

	ctrl := controller.NewController(friendsCollector, groupsCollector, pg)

	err = ctrl.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
