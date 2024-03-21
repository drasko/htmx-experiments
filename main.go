package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type SquareDimensions struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

func squareHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	http.ServeFile(w, r, "templates/index.html")
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	// Get new dimensions from the request
	newWidth := r.FormValue("width")
	newHeight := r.FormValue("height")

	// Log the received dimensions for debugging
	log.Printf("Received dimensions: width=%s, height=%s\n", newWidth, newHeight)

	// Perform any necessary validation on the dimensions
	// For simplicity, we'll assume dimensions are valid

	// Write the HTML for the resized square to the response
	fmt.Fprintf(w, `<div id="square" class="resizable"
     style="width: %s; height: %s; background-color: blue;"
     onmousedown="startDrag(event)">
    <!-- This div will be replaced by the resized square -->
    </div>`, newWidth, newHeight)
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/square", squareHandler)
	r.Post("/resize", resizeHandler)

	// Start server
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
