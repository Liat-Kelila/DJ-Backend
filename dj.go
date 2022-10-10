package main

import (
  "fmt"
  "os"
  "log"
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  "github.com/gorilla/handlers"
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
    // dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

    dbURI := postgres://obqsbjrxwcyqjq:62e0ead0553e1f53b5a699ae1aa400b82c4626dd76abad4f0a3d74127f9f39c8@ec2-23-20-140-229.compute-1.amazonaws.com:5432/ddblm95hab27e3


  // Opening connection to Database
    db, err = gorm.Open(dialect, os.Getenv("dbURI")
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
  router.HandleFunc("/previousGigs", createPreviousGigs).Methods("POST")

  //Delete route
  router.HandleFunc("/previousGig/{id}", deletePreviousGigs).Methods("DELETE")

  //Update route
  router.HandleFunc("/previousGig/{id}", updateGig).Methods("PUT")

  //Enable CORS
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
  origins := handlers.AllowedOrigins([]string{"*"})

  //Connect to server
  log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
  }

  // respondJSON makes the response with payload as json format
  func respondJSON(w http.ResponseWriter, status int, payload interface{}) {

      response, err := json.Marshal(payload)
        if err != nil {
          w.WriteHeader(http.StatusInternalServerError)
          w.Write([]byte(err.Error()))
          return
    }
        w.WriteHeader(status)
        w.Write([]byte(response))
    }

// respondError makes the error response with payload as json format
  func respondError(w http.ResponseWriter, code int, message string) {
      respondJSON(w, code, map[string]string{"error": message})
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

  //Update
  func updateGig(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var previousGigs PreviousGigs
    db.First(&previousGigs, params["id"])
    json.NewDecoder(r.Body).Decode(&previousGigs)
    db.Save(&previousGigs)
    json.NewEncoder(w).Encode(previousGigs)
  }


  //Delete
  func deletePreviousGigs(w http.ResponseWriter,r *http.Request) {
    params := mux.Vars(r)
    var previousGigs PreviousGigs
    db.Delete(&previousGigs, params["id"])
    json.NewEncoder(w).Encode(&previousGigs)
  }
