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

var happy = "0"
var angry = "0"
var sad = "0"
var fear = "0"
var surprise = "0"
var disgust = "0"

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

func setDoctorStatusAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	happy = r.Form["happy"][0]
	angry = r.Form["angry"][0]
	sad = r.Form["sad"][0]
	fear = r.Form["fear"][0]
	surprise = r.Form["surprise"][0]
	disgust = r.Form["disgust"][0]
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

func getDoctorStatusAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprint(w, happy+" "+angry+" "+sad+" "+fear+" "+surprise+" "+disgust)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/setDoctorStatus", setDoctorStatus)
	mux.HandleFunc("/setPatientStatus", setPatientStatus)
	mux.HandleFunc("/getDoctorStatus", getDoctorStatus)
	mux.HandleFunc("/getPatientStatus", getPatientStatus)
	mux.HandleFunc("/setDoctorStatusAll", setDoctorStatusAll)
	mux.HandleFunc("/getDoctorStatusAll", getDoctorStatusAll)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
