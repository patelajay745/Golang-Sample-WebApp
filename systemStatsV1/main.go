package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
}

func getSystemStats() (SystemStats, error) {
	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStats{}, err
	}

	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	return SystemStats{
		CPUUsage:    cpuPercentages[0],
		MemoryUsage: virtualMemory.UsedPercent,
	}, nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>System Stats</title>
                <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
                <style>
                    body {
                        font-family: Arial, sans-serif;
                        margin: 20px;
                    }
                    .chart-container {
                        width: 80%%;
                        margin: auto;
                    }
                </style>
            </head>
            <body>
                <div class="container">
                    <h1 class="mt-5">System Stats</h1>
                    <div class="chart-container">
                        <canvas id="cpuChart"></canvas>
                    </div>
                    <div class="chart-container">
                        <canvas id="memoryChart"></canvas>
                    </div>
                </div>
                <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
                <script>
                    async function fetchStats() {
                        const response = await fetch('/stats');
                        return response.json();
                    }

                    async function updateCharts(cpuChart, memoryChart) {
                        const stats = await fetchStats();
                        const currentTime = new Date().toLocaleTimeString();

                        if (cpuChart.data.labels.length > 10) {
                            cpuChart.data.labels.shift();
                            cpuChart.data.datasets[0].data.shift();
                        }

                        if (memoryChart.data.labels.length > 10) {
                            memoryChart.data.labels.shift();
                            memoryChart.data.datasets[0].data.shift();
                        }

                        cpuChart.data.labels.push(currentTime);
                        cpuChart.data.datasets[0].data.push(stats.cpu_usage);

                        memoryChart.data.labels.push(currentTime);
                        memoryChart.data.datasets[0].data.push(stats.memory_usage);

                        cpuChart.update();
                        memoryChart.update();
                    }

                    document.addEventListener('DOMContentLoaded', () => {
                        const ctxCpu = document.getElementById('cpuChart').getContext('2d');
                        const ctxMemory = document.getElementById('memoryChart').getContext('2d');

                        const cpuChart = new Chart(ctxCpu, {
                            type: 'line',
                            data: {
                                labels: [],
                                datasets: [{
                                    label: 'CPU Usage (%)',
                                    data: [],
                                    borderColor: 'rgba(75, 192, 192, 1)',
                                    borderWidth: 1,
                                    fill: false
                                }]
                            },
                            options: {
                                scales: {
                                    y: {
                                        beginAtZero: true,
                                        max: 100
                                    }
                                }
                            }
                        });

                        const memoryChart = new Chart(ctxMemory, {
                            type: 'line',
                            data: {
                                labels: [],
                                datasets: [{
                                    label: 'Memory Usage (%)',
                                    data: [],
                                    borderColor: 'rgba(153, 102, 255, 1)',
                                    borderWidth: 1,
                                    fill: false
                                }]
                            },
                            options: {
                                scales: {
                                    y: {
                                        beginAtZero: true,
                                        max: 100
                                    }
                                }
                            }
                        });

                        setInterval(() => updateCharts(cpuChart, memoryChart), 2000);
                    });
                </script>
            </body>
            </html>
        `)
	})

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
