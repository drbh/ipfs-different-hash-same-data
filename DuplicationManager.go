package main

import (
	"database/sql"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	_ "github.com/mattn/go-sqlite3"
)

type DuplicationManager struct {
	db      *sql.DB
	ipfsapi *shell.Shell
}

func (d *DuplicationManager) createDatabase() int {
	statement, _ := d.db.Prepare("CREATE TABLE IF NOT EXISTS addresses (id INTEGER PRIMARY KEY, originCID TEXT, similarCID TEXT)")
	statement.Exec()
	return 1
}

func (d *DuplicationManager) getEqualCIDs(cid string) int {
	query := fmt.Sprint("SELECT similarCID FROM addresses WHERE originCID = '", cid, "'; ")
	rows, _ := d.db.Query(query)
	var similarCID string
	for rows.Next() {
		rows.Scan(&similarCID)
		fmt.Println(similarCID)
	}
	return 1
}

func (d *DuplicationManager) compareContentsOfCID(c1 string, c2 string) {
	cid, _ := d.ipfsapi.ObjectGet(c1)
	cid2, _ := d.ipfsapi.ObjectGet(c2)
	if cid.Data == cid2.Data {
		// files are equal
		query := fmt.Sprint("SELECT COUNT(*) FROM addresses WHERE originCID = '", c1, "' AND similarCID = '", c2, "'; ")
		rows, _ := d.db.Query(query)
		var count int
		for rows.Next() {
			rows.Scan(&count)
		}
		if count < 1 {
			// insert new link
			statement, _ := d.db.Prepare("INSERT INTO addresses (originCID, similarCID) VALUES (?, ?)")
			statement.Exec(c1, c2)
			// save the inverse relationship
			statement, _ = d.db.Prepare("INSERT INTO addresses (originCID, similarCID) VALUES (?, ?)")
			statement.Exec(c2, c1)
			// tell user about it
			fmt.Println("+1 CIDs return equal byte arrays, DB updated")
		} else {
			fmt.Println("We already have this entry")
		}
	}
}
