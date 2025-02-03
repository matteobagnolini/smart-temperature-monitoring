package http

import (
	"control-unit-backend/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// We want to responde to both GET and POST on /api/data.
// GET should be used to get the N measurements.
// POST should be used to communicate the opening of the window when dashboard goes MANUAL mode.

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "GET" {
		handleGetData(w)
	} else if r.Method == "POST" {
		handlePostData(w, r)
	}
}

func handlePostData(w http.ResponseWriter, r *http.Request) {
	// Will be posted:
	// - Window opening if Dashboard manual mode is selected
	// 	 (new System state should be defined: MANUAL_DASHBOARD)
}

func handleGetData(w http.ResponseWriter) {
	var temps []float32
	for _, d := range models.DataSampler.GetDatas() {
		temps = append(temps, d.Temp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(
		JsonData{
			Temps:             temps,
			Avg:               models.DataSampler.GetLastHistoryData().Avg,
			Max:               models.DataSampler.GetLastHistoryData().Max,
			Min:               models.DataSampler.GetLastHistoryData().Min,
			CurrState:         string(models.System.TempState()),
			WindowOpeningPerc: models.System.WindowPercOpening(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func handleStateRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("could not read body")
		}
		if string(body) == "alarm: resolved" {
			models.System.SetSysState(models.SystemState(models.NORMAL))
		}
	}
	// Need to manage also the MANUAL_DASHBOARD state here
}

func StartHttpServer(addres string, port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", handleDataRequest)
	mux.HandleFunc("/api/state", handleStateRequest)
	server := &http.Server{
		Addr:    addres + ":" + port,
		Handler: mux,
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server one closed")
		} else if err != nil {
			log.Fatal(err)
		}
	}()
}
