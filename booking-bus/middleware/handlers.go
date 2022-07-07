package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/archit-tiket/booking-system/booking-bus/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//create connection with postgresql :/

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("USER"), os.Getenv("NAME"), os.Getenv("PASSWORD"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the db!")
	return db
}

func GetAllBuses(w http.ResponseWriter, r *http.Request) {
	buses, err := getAllBuses()

	if err != nil {
		log.Fatalf("Unable to get all users.%v", err)
	}
	json.NewEncoder(w).Encode(buses)
}

//handlers ----------------->>>>>>>>>>>>>>>

func getAllBuses() ([]models.Buses, error) {
	db := createConnection()

	defer db.Close()

	var buses []models.Buses

	sqlStatement := `SELECT * FROM buses`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query.%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var bus models.Buses
		err = rows.Scan(&bus.ID, &bus.Name, &bus.Start, &bus.End)
		if err != nil {
			log.Fatalf("unable to scan the row.%v", err)
		}

		buses = append(buses, bus)
	}
	return buses, err
}
