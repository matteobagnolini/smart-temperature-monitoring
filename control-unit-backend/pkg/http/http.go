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
const autoStateAddress = "/api/auto-state"
const dashboardWindowOpeningAddress = "/api/window-opening"
const historyDataAddress = "/api/history"

func StartHttpServer(addres string, port string) {
	mux := http.NewServeMux()
	mux.HandleFunc(dataAddress, handleGetData)
	mux.HandleFunc(resolveAlarmAddress, handleResolveAlarmRequest)
	mux.HandleFunc(autoStateAddress, handleAutoStateRequest)
	mux.HandleFunc(manualStateAddress, handleManualStateRequest)
	mux.HandleFunc(dashboardWindowOpeningAddress, handleDashboardWindowOpeningRequest)
	mux.HandleFunc(historyDataAddress, handleHistoryDataRequest)

	server := &http.Server{
		Addr:    addres + ":" + port,
		Handler: mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func enableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func handleGetData(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
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
	enableCors(w, r)
	models.System.ResolveAlarm()
}

func handleManualStateRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if models.System.SysState() == models.SystemState(models.AUTOMATIC) {
		models.System.SetSysState(models.SystemState(models.DASHBOARD_MANUAL))
	}
}

func handleDashboardWindowOpeningRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	var windowOpeningFromDashboard JsonWindowOpening
	err := json.NewDecoder(r.Body).Decode(&windowOpeningFromDashboard)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	models.System.SetWindPercOpening(windowOpeningFromDashboard.WindowOpeningPerc)
}

func handleAutoStateRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if models.System.SysState() == models.SystemState(models.DASHBOARD_MANUAL) {
		models.System.SetSysState(models.SystemState(models.AUTOMATIC))
	}
}

func handleHistoryDataRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(models.DataSampler.GetHistoryDatas())
	if err != nil {
		log.Fatal(err)
	}
}
