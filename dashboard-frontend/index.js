const DATA_ADDRESS = "http://localhost:3333/api/data"
const RESOLVE_ALARM_ADDRESS = "http://localhost:3333/api/resolve-alarm"

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

createChart();
setInterval(fetchTemperatureData, 250);
