package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

var doctorStatus = "happy"
var patientStatus = "happy"

func serveApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func setDoctorStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	doctorStatus = r.Form["status"][0]
}

func setPatientStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	patientStatus = r.Form["status"][0]
}

func getPatientStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, patientStatus)
}

func getDoctorStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, doctorStatus)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/setDoctorStatus", setDoctorStatus)
	mux.HandleFunc("/setPatientStatus", setPatientStatus)
	mux.HandleFunc("/getDoctorStatus", getDoctorStatus)
	mux.HandleFunc("/getPatientStatus", getPatientStatus)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
