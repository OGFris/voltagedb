package main

import (
	"fmt"
	"github.com/OGFris/voltagedb/database"
	"github.com/OGFris/voltagedb/routes/player"
	"github.com/OGFris/voltagedb/utils"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.InitDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/player/create", player.Create)
	router.HandleFunc("/player", player.Get)
	router.HandleFunc("/player/ban", player.Ban)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("Server is running Port: ", port)
	utils.PanicErr(s.ListenAndServe())
}
