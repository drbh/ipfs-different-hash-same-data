package main

import (
	"database/sql"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	database, _ := sql.Open("sqlite3", "./dhsd.db")
	sh := shell.NewShell("localhost:5001")
	dhsd := DuplicationManager{database, sh}
	dhsd.createDatabase()

	argsWithProg := os.Args

	if len(argsWithProg) < 3 {
		fmt.Println("Please specify either a `compare` or `eq` action with arguments")
		return
	}
	action_prt := argsWithProg[1]
	cid_1_ptr := argsWithProg[2]

	if action_prt == "compare" {

		cid_2_ptr := argsWithProg[3]
		dhsd.compareContentsOfCID(cid_1_ptr, cid_2_ptr)
	}

	if action_prt == "eq" { // fetch my equivlant ids
		dhsd.getEqualCIDs(cid_1_ptr)
	}

	if action_prt == "" {

	}
}
