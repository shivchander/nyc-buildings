package main

import (
	. "nyc-buildings/config"
	. "nyc-buildings/dao"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = BuildingsDAO{}

// GET list of Buildings
func AllBuildingsEndPoint(w http.ResponseWriter, r *http.Request) {
	Buildings, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, Buildings)
}


func FindBuildingsBinEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bin, _ := strconv.Atoi(params["BIN"])

	Buildings, err := dao.FindById(bin)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Buildings BIN")
		return
	}
	respondWithJson(w, http.StatusOK, Buildings)
}


func FindBuildingsEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	start, _ := strconv.Atoi(params["Construct_Yr_start"])
	end, _ := strconv.Atoi(params["Construct_Yr_end"])
	//fmt.Println(start)
	Buildings, err := dao.FindByYr(start, end)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Buildings BIN")
		return
	}
	respondWithJson(w, http.StatusOK, Buildings)
}


func UpdateBuildingsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteBuildingsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}


func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/", http.StatusFound)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
	dao.PushAll()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()

	
	r.HandleFunc("/buildings", AllBuildingsEndPoint).Methods("GET")
	//	r.HandleFunc("/Buildings", CreateBuildingsEndPoint).Methods("GET")
	r.HandleFunc("/buildings", UpdateBuildingsEndPoint).Methods("PUT")
	r.HandleFunc("/buildings", DeleteBuildingsEndPoint).Methods("DELETE")
	r.HandleFunc("/buildings/{BIN}", FindBuildingsBinEndpoint).Methods("GET")
	r.HandleFunc("/buildings/{Construct_Yr_start}/{Construct_Yr_end}", FindBuildingsEndpoint).Methods("GET")
		r.HandleFunc("/index", index).Methods("GET")
		s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
		r.PathPrefix("/static/").Handler(s).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
