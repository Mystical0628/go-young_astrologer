package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	package_core "go-young_astrologer/package-core"
	"go-young_astrologer/package-core/scalar"
	"go-young_astrologer/service-api/application/controllers"
	"log"
	"net/http"
	"strconv"
)

func init() {
	path := flag.String("config", "", "Path to .env")

	flag.Parse()

	if err := godotenv.Load(*path); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	core, err := package_core.NewCore()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/nasa/apod", func(w http.ResponseWriter, r *http.Request) {
		controller := controllers.NewNasaController(core)
		resp, err := controller.ApodAction()
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "application/json")
		_, err = fmt.Fprint(w, string(b))
		if err != nil {
			return
		}
	})

	// /nasa/apod/date/2023-10-24
	http.HandleFunc("/nasa/apod/date/", func(w http.ResponseWriter, r *http.Request) {
		var date scalar.Date
		err := date.ParseString(r.URL.Path[len("/nasa/apod/date/"):])
		if err != nil {
			return
		}

		controller := controllers.NewNasaController(core)
		resp, err := controller.ApodDateAction(date)
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "application/json")
		_, err = fmt.Fprint(w, string(b))
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(core.Config.ServerPort), nil))
}
