package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var doctorStatus = "neutral"
var patientStatus = "neutral"

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
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/setDoctorStatus", setDoctorStatus)
	http.HandleFunc("/setPatientStatus", setPatientStatus)
	http.HandleFunc("/getDoctorStatus", getDoctorStatus)
	http.HandleFunc("/getPatientStatus", getPatientStatus)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
