package main

import (
  "fmt"
  "os"
  "log"
  "encoding/json"
  "net/http"

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


  //Get a single object
  router.HandleFunc("/previousGig/{id}", getPreviousGig).Methods("GET")

  //Create route
  router.HandleFunc("/create/previousGigs", createPreviousGigs).Methods("POST")

  //Delete route
  router.HandleFunc("/delete/previousGig/{id}", deletePreviousGigs).Methods("DELETE")



  //Connect to server
  log.Fatal(http.ListenAndServe(":8080", router))
}

  //API Controllers
  //Get Function
  func getPreviousGigs(w http.ResponseWriter, r *http.Request) {
    var previousGigs []PreviousGigs

    db.Find(&previousGigs)

    json.NewEncoder(w).Encode(&previousGigs)
  }

  func getPreviousGig(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var gig PreviousGigs

    db.First(&gig, params["id"])

    json.NewEncoder(w).Encode(gig)
  }

  func createPreviousGigs(w http.ResponseWriter, r *http.Request) {
    var previousGigs PreviousGigs
    json.NewDecoder(r.Body).Decode(&previousGigs)

    createdPreviousGigs := db.Create(&previousGigs)
    err = createdPreviousGigs.Error
    if err != nil {
      json.NewEncoder(w).Encode(err)
    } else {
      json.NewEncoder(w).Encode(&previousGigs)
    }
  }

  //Delete
  func deletePreviousGigs(w http.ResponseWriter,r *http.Request) {
    params := mux.Vars(r)

    var previousGigs PreviousGigs

    db.First(&previousGigs, params["id"])
    db.Delete(&previousGigs)

    json.NewEncoder(w).Encode(&previousGigs)
  }
