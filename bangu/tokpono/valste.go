package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// .i lo Valsi cu pa tokpono valsi
type Valsi struct {
	Word        string               `json:"word"`
	Definitions map[string]Vlavelcki `json:"definitions"`
	Examples    []Mupli              `json:"examples"`
	Misc        struct {
		Notes string `json:"notes"`
	} `json:"misc"`
}

// .i lo Mupli cu tokpono je glico be mupli
type Mupli struct {
	TokiPona string `json:"toki-pona"`
	English  string `json:"english"`
}

// .i lo Vlavelcki cu pa tokpono vlavelcki
type Vlavelcki struct {
	Meanings []string `json:"meanings"`
	FullName string   `json:"full_name"` // selma'o
}

// .i Morji le tokpono valste
func Morji() ([]Valsi, error) {
	fin, err := os.Open("tokpono-valste.json")
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	var result []Valsi
	err = json.NewDecoder(fin).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var fintiJubme = `
CREATE TABLE IF NOT EXISTS valsi
	( cmene TEXT NOT NULL
	, selmaho TEXT NOT NULL
	, selvla TEXT NOT NULL
	);`

func FintiJubme(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, fintiJubme)
	if err != nil {
		return err
	}

	return nil
}

var (
	dbPath = flag.String("db-path", "./tokpono.db", "database path")
)

func main() {
	flag.Parse()

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("database opened")

	ctx := context.Background()

	valste, err := Morji()
	if err != nil {
		log.Fatal(err)
	}
	log.Println ("pu morji")

	err = FintiJubme(ctx, db)
	if err != nil {log.Fatal(err)}
	log.Println("pu finti jubme")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	for _, valsi := range valste {
		for _, def := range valsi.Definitions {
			for _, selvla := range def.Meanings {
				_, err := tx.ExecContext(ctx, "INSERT INTO valsi (cmene, selmaho, selvla) values (?,?,?)", valsi.Word, def.FullName, selvla)
				if err != nil {log.Fatal(err)}

				log.Printf("%s: %s - %s", valsi.Word, def.FullName, selvla)
			}
		}
	}

	log.Println("pu valste morji")
}
