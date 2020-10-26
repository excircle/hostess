package sql

import "database/sql"

import _ "github.com/mattn/go-sqlite3"

func DbExists() bool { 
    _, err := sql.Open("sqlite3", "./sqlite-database.db")
    if err != nil {
        return true
      } else {
        return false
    }
}