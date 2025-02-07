const HISTORY_PATH = "http://localhost:3333/api/history"

document.addEventListener("DOMContentLoaded", function () {
    fetch(HISTORY_PATH)
        .then(response => response.json())
        .then(data => {
            const tableBody = document.getElementById("historyTableBody");
            tableBody.innerHTML = "";

            data.forEach(entry => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${entry.Date}</td>
                    <td>${entry.Avg.toFixed(2)}</td>
                    <td>${entry.Min.toFixed(2)}</td>
                    <td>${entry.Max.toFixed(2)}</td>
                `;
                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error("Error loading history data:", error));
});
