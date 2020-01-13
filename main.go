package main

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// Book structure to collection book
type Book struct {
	Title   string `json:"title"`
	NoPages int32  `json:"noPages"`
}

func main() {
	ctx := context.Background()
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	PrintError(err, "error 01")

	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	PrintError(err, "error 01")

	// Open "examples_books" database
	var db driver.Database
	foundDB, err := c.DatabaseExists(ctx, "examples_books")
	PrintError(err, "error 02.exite databasen")
	if foundDB {
		db, err = c.Database(ctx, "examples_books")
		PrintError(err, "error 02.open database")
	} else {
		db, err = c.CreateDatabase(ctx, "examples_books", &driver.CreateDatabaseOptions{})
		PrintError(err, "error 02.create database")
	}

	// Open "books" collection
	var col driver.Collection
	found, err := db.CollectionExists(ctx, "books")
	PrintError(err, "error 03.exite collection")
	if found {
		col, err = db.Collection(ctx, "books")
		PrintError(err, "error 03")
	} else {
		col, err = db.CreateCollection(ctx, "books", &driver.CreateCollectionOptions{})
		PrintError(err, "error 04")
	}

	// Create document
	book := Book{
		Title:   "ArangoDB Cookbook",
		NoPages: 257,
	}
	meta, err := col.CreateDocument(ctx, book)
	PrintError(err, "error 05")

	fmt.Println(meta)
	fmt.Printf("Created document in collection '%s' in database '%s'\n", col.Name(), db.Name())

	var doc Book
	meta2, err := col.ReadDocument(ctx, meta.Key, &doc)
	PrintError(err, "error 06")

	fmt.Println("reading from collection")
	fmt.Println(meta2)

}

func PrintError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err.Error())
	}
}
