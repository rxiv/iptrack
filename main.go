package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
        "strconv"
        "math/rand"

	"github.com/gorilla/mux"
)

type Computer struct {
        ID      string  `json:"id"`
	Name    string  `json:"name"`
	IP      string  `json:"ip"`
}

var computers []Computer

func getComputers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(computers)
}

func updateComputer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var computer Computer
	_ = json.NewDecoder(r.Body).Decode(&computer)
	for index, item := range computers {
		if item.Name == computer.Name {
			computers = append(computers[:index], computers[index+1:]...)
                        computer.ID = item.ID
            	        computers = append(computers, computer)
			json.NewEncoder(w).Encode(computer)
                        return
		}
	}
	computer.ID = strconv.Itoa(rand.Intn(10000000000))
	computers = append(computers, computer)
	json.NewEncoder(w).Encode(computer)

}

func main() {
	r := mux.NewRouter()

    computers = append(computers, Computer{ID: "1", Name: "edgemax", IP: "192.168.1.1"})
    computers = append(computers, Computer{ID: "2", Name: "cupid", IP: "192.168.1.5"})
    computers = append(computers, Computer{ID: "3",Name: "pphole", IP: "192.168.1.10"})

	r.HandleFunc("/comp", getComputers).Methods("GET")
	r.HandleFunc("/comp", updateComputer).Methods("PUT")

	fmt.Printf("Starting Server on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
