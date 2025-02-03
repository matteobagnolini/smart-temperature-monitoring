let temperatureChart;
let tempData = [];

async function fetchTemperatureData() {
    try {
        const response = await fetch("http://localhost:3333/api/data");
        const data = await response.json();

        // Aggiorna il grafico
        updateChart(data.Temps);

        // Aggiorna i valori nella sezione info
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
                    suggestedMin: 15,  // Limite inferiore 10°C
                    suggestedMax: 25,  // Limite superiore 30°C
                    title: { display: true, text: "Temperature (°C)" }
                }
            }
        }
    });
}

function resolveAlarm() {
    fetch("http://localhost:3333/api/resolve-alarm", {
        method: "POST",
    }).then(() => {
        alert("Alarm resolved!");
    }).catch(error => console.error("Error resolving alarm:", error));
}

// Creazione del grafico e aggiornamento automatico ogni 2 secondi
createChart();
setInterval(fetchTemperatureData, 250);
