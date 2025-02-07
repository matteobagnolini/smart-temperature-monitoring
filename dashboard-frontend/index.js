const PORT = "3333"
const BASE_ADDRESS = "http://localhost:" + PORT
const DATA_ADDRESS = BASE_ADDRESS + "/api/data"
const RESOLVE_ALARM_ADDRESS = BASE_ADDRESS + "/api/resolve-alarm"
const MANUAL_STATE_ADDRESS = BASE_ADDRESS + "/api/manual-state"
const AUTO_STATE_ADDRESS = BASE_ADDRESS + "/api/auto-state"
const WINDOW_OPENING_ADDRESS = BASE_ADDRESS + "/api/window-opening"

let temperatureChart;
let tempData = [];

async function fetchTemperatureData() {
    try {
        const response = await fetch(DATA_ADDRESS);
        const data = await response.json();

        updateChart(data.Temps);

        document.getElementById("avgTemp").innerText = data.Avg.toFixed(2);
        document.getElementById("maxTemp").innerText = data.Max.toFixed(2);
        document.getElementById("minTemp").innerText = data.Min.toFixed(2);
        document.getElementById("currState").innerText = data.CurrState;
        document.getElementById("windowPerc").innerText = data.WindowOpeningPerc;

        // Button is disabled if state != ALARM
        const alarmBtn = document.getElementById("resolveAlarmBtn");
        if (data.CurrState === "ALARM") {
            alarmBtn.disabled = false;
        } else {
            alarmBtn.disabled = true;
        }

    } catch (error) {
        console.error("Error fetching data:", error);
    }
}

function updateChart(newTemps) {
    tempData = newTemps;
    if (temperatureChart) {
        temperatureChart.data.labels = Array.from({ length: tempData.length }, (_, i) => i + 1);
        temperatureChart.data.datasets[0].data = tempData;
        temperatureChart.update();
    }
}

function createChart() {
    const ctx = document.getElementById("temperatureChart").getContext("2d");
    temperatureChart = new Chart(ctx, {
        type: "line",
        data: {
            labels: [],
            datasets: [{
                label: "Temperature (°C)",
                borderColor: "red",
                backgroundColor: "rgba(255, 99, 132, 0.2)",
                data: [],
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                x: {
                    title: { display: true, text: "Time" }
                },
                y: {
                    suggestedMin: 15,
                    suggestedMax: 25,
                    title: { display: true, text: "Temperature (°C)" }
                }
            }
        }
    });
}

function resolveAlarm() {
    fetch(RESOLVE_ALARM_ADDRESS, {
        method: "POST",
    }).then(() => {
        alert("Alarm resolved!");
    }).catch(error => console.error("Error resolving alarm:", error));
}

document.addEventListener("DOMContentLoaded", function () {
    const modeSelector = document.getElementById("modeSelector");

    modeSelector.addEventListener("change", function () {
        const selectedMode = modeSelector.value;
        const url = selectedMode === "MANUAL"
            ? MANUAL_STATE_ADDRESS
            : AUTO_STATE_ADDRESS;

        fetch(url, { method: "POST" })
            .then(() => console.log(`Mode changed to: ${selectedMode}`))
            .catch(error => console.error("Error changing mode:", error));
    });

    // Display or hide slider for controlling window opening
    modeSelector.addEventListener("change", function () {
        if (modeSelector.value === "MANUAL") {
            manualControls.style.display = "block";
        } else {
            manualControls.style.display = "none";
        }
    });
});

document.addEventListener("DOMContentLoaded", function () {
    const slider = document.getElementById("windowSlider");
    const windowValue = document.getElementById("windowValue");
    const setWindowBtn = document.getElementById("setWindowBtn");

    slider.addEventListener("input", function () {
        windowValue.textContent = slider.value + "%";
    });

    setWindowBtn.addEventListener("click", function () {
        const percentage = parseInt(slider.value);
        fetch(WINDOW_OPENING_ADDRESS, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ WindowOpeningPerc: percentage }),
        })
        .then(() => console.log(`Window opening set to: ${percentage}%`))
        .catch(error => console.error("Error setting window opening:", error));
    });
});

createChart();
setInterval(fetchTemperatureData, 250);
