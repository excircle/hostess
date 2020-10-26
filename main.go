package main

import hsql "github.com/excircle/hostess/sql"
import "fmt"


func init() {
  //Check that we have a database
}

func main() {
    //if the sqlite database exists
    if hsql.DbExists() { 
      fmt.Println("SQLite Database exists.")
    } else {
      fmt.Println("SQLite Database DOES NOT exist!")
    }
    
}
