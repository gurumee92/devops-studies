package people

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"

	"studydts/lib/model"
)

const dbURL = "root:root@tcp(127.0.0.1:3306)/test_db"

// Repository is
type Repository struct {
	db *sql.DB
}

// NewRepository is
func NewRepository() *Repository {
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Cannot ping the db: %v", err)
	}

	return &Repository{
		db: db,
	}
}

// GetPerson is
func (r *Repository) GetPerson(ctx context.Context, name string) (model.Person, error) {
	query := "select title, description from people where name = ?"
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		"get-person",
		opentracing.Tag{Key: "db.statement", Value: query},
	)
	defer span.Finish()

	rows, err := r.db.QueryContext(ctx, query, name)

	if err != nil {
		return model.Person{}, nil
	}

	defer rows.Close()

	for rows.Next() {
		var title, description string
		err := rows.Scan(&title, &description)

		if err != nil {
			return model.Person{}, nil
		}

		return model.Person{
			Name:        name,
			Title:       title,
			Description: description,
		}, nil
	}

	return model.Person{
		Name: name,
	}, nil
}

// Close is
func (r *Repository) Close() {
	r.db.Close()
}
