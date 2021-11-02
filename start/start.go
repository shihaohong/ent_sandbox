package main

import (
	"context"
	"ent_sandbox/ent"
	"ent_sandbox/ent/user"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	_, err = QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed to query user: %v", err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(27).
		SetName("shihaohong").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.
		User.
		Query().
		Where(user.Name("shihaohong")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
