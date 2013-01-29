// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
	"net/http"
)

var (
	user     = flag.String("user", "root", "user to connect to the database")
	password = flag.String("password", "", "password to connect to the database")
	pswd     string
)

func init() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("You must provide listen address!")
	}
	if *password != "" {
		pswd = ":" + *password
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf("%s%s@/", *user, pswd)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("select 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	log.Printf("Listening at %s...", flag.Args()[0])
	log.Fatal(http.ListenAndServe(flag.Args()[0], nil))
}
