package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Go Lang App</title>
                <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
                <style>
                    body {
                        font-family: Arial, sans-serif;
                        margin: 20px;
                        background-color: #f0f8ff; /* AliceBlue */
                    }
                    .navbar {
                        background-color: #ff6347; /* Tomato */
                    }
                    .navbar-brand, .nav-link {
                        color: #ffffff !important; /* White */
                    }
                    .path-display {
                        margin-top: 20px;
                        padding: 10px;
                        border: 1px solid #ccc;
                        background-color: #e6e6fa; /* Lavender */
                        color: #4b0082; /* Indigo */
                    }
                    .container {
                        background-color: #fafad2; /* LightGoldenRodYellow */
                        padding: 20px;
                        border-radius: 10px;
                    }
                    h1 {
                        color: #4682b4; /* SteelBlue */
                    }
                    p.lead {
                        color: #2e8b57; /* SeaGreen */
                    }
                </style>
            </head>
            <body>
                <nav class="navbar navbar-expand-lg">
                    <a class="navbar-brand" href="#">Go App</a>
                    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarNav">
                        <ul class="navbar-nav">
                            <li class="nav-item active">
                                <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">Features</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">Pricing</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">Contact us</a>
                            </li>
                        </ul>
                    </div>
                </nav>
                <div class="container">
                    <div class="row">
                        <div class="col-md-12">
                            <h1 class="mt-5">Welcome to Go App</h1>
                            <p class="lead">This is a simple web application using Golang.</p>
                            <div class="path-display">
                                Current Path: <span id="current-path"></span>
                            </div>
                        </div>
                    </div>
                </div>
                <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
                <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.5.4/dist/umd/popper.min.js"></script>
                <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
                <script>
                    // Get the current path
                    const currentPath = window.location.pathname;

                    // Display the current path in the span element
                    document.getElementById('current-path').textContent = currentPath;
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
