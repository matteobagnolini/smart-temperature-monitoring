package http

import (
	"control-unit-backend/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

const dataAddress = "/api/data"
const resolveAlarmAddress = "/api/resolve-alarm"
const manualStateAddress = "/api/manual-state"
const dashboardWindowOpeningAddress = "/api/window-opening"

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "GET" {
		handleGetData(w)
	}
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

func handleResolveAlarmRequest(w http.ResponseWriter, r *http.Request) {
	models.System.SetSysState(models.SystemState(models.NORMAL))
}

func handleManualStateRequest(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func handleDashboardWindowOpeningRequest(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func StartHttpServer(addres string, port string) {
	mux := http.NewServeMux()
	mux.HandleFunc(dataAddress, handleDataRequest)
	mux.HandleFunc(resolveAlarmAddress, handleResolveAlarmRequest)
	mux.HandleFunc(manualStateAddress, handleManualStateRequest)
	mux.HandleFunc(dashboardWindowOpeningAddress, handleDashboardWindowOpeningRequest)

	server := &http.Server{
		Addr:    addres + ":" + port,
		Handler: mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}
