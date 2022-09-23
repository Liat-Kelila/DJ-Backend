package main

import (
  "fmt"
  "os"
  "log"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PreviousGigs struct {
    gorm.Model

    Title string
    Date string
    Location string
    Notes string

}

var (
    previousGigs = &PreviousGigs{
      Title: "Opened for Themba",
      Date: "April 8th, 2022",
      Location: "Hidden Door, Tampa",
      Notes: "A jungle-themed night with both women & plants hanging from the ceiling, nude tigress women dancing on either side of the stage.",
    }
)

  //Globally declared varialbes
var db *gorm.DB
var err error


func main () {
    // Loading environment variables
    dialect := os.Getenv("DIALECT")
    host := os.Getenv("HOST")
    dbPort := os.Getenv("DBPORT")
    user := os.Getenv("USER")
    dbName := os.Getenv("NAME")
    password := os.Getenv("PASSWORD")

    // Database connection string
    dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)



  // Opening connection to Database
    db, err = gorm.Open(dialect, dbURI)
    if err != nil {
       log.Fatal(err)
    } else {
      fmt.Println("Successfully connected to database!")
    }

  // Close connection to database when the main function finishes
  defer db.Close()

  // Make migrations to the database (only done once!)
  db.AutoMigrate(&PreviousGigs{})

  //Create data
  // db.Create(previousGigs)
  // {
  //   db.Create(&previousGigs)
  // }

  //Api routes
  router := mux.NewRouter()

  //Get Route
  router.HandleFunc("/previousGigs", getPreviousGigs).Methods("GET")

  //Get Function
}
