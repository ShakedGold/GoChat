package server

import (
	"fmt"
	"net/http"
	"shaked/gochat/client"
	"strings"

	"github.com/rs/cors"
)

const (
	// Port is the port number the server will listen on
	Port = 8080
	// Host is the host address the server will listen on
	Host = "localhost"
)

func Start() {
	client.Setup()
	router := http.NewServeMux()
	router.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		// serve the client/frontend/dist directory
		http.ServeFile(w, r, "client/frontend/dist/index.html")
	}))

	// serve assets
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("client/frontend/dist/assets"))))

	// create the hub map
	hub_names := []string{"general", "admin", "help", "java", "nizan", "nitai", "almog"}
	hubs := make(map[string]*client.Hub)
	for _, name := range hub_names {
		hubs[name] = client.NewHub()
		go hubs[name].Run()
	}

	// setup /hubs endpoint
	router.HandleFunc("/hubs", func(w http.ResponseWriter, r *http.Request) {
		// wrap the hub names in quotes
		json_hub_names := make([]string, len(hub_names))
		copy(json_hub_names, hub_names)
		for i, name := range json_hub_names {
			json_hub_names[i] = fmt.Sprintf("\"%s\"", name)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("[%s]", strings.Join(json_hub_names, ","))))
	})

	// serve the websocket
	router.HandleFunc("/ws/{hub}", func(w http.ResponseWriter, r *http.Request) {
		hub := r.PathValue("hub")
		if _, ok := hubs[hub]; !ok {
			http.Error(w, "Invalid hub", http.StatusBadRequest)
			return
		}

		client.ServeWs(hubs[hub], w, r)
	})

	fmt.Printf("Attempting to start server on http://%s:%d\n", Host, Port)

	handler := cors.AllowAll().Handler(router)
	err := http.ListenAndServe(Host+":"+fmt.Sprint(Port), handler)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
