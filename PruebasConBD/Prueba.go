package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Datos struct {
	Data []Membresia `json:"data"`
}
type Membresia struct {
	Id_membresia string `json:"id_membresia"`
	Tipo_membresia string `json:"tipo_membresia,omitempty"`
}

var membresias []Membresia
var membresia Membresia
var data Datos
// EndPoints
func GetMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
    for _,item:= range membresias{
    	if item.Id_membresia==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//RecuperarXId(tabla,1)
	json.NewEncoder(w).Encode(membresia)
}
func GetMembershipsEndpoint(w http.ResponseWriter, req *http.Request){
	//enableCors(&w)
	data.Data=membresias
	json.NewEncoder(w).Encode(data)
}
func CreateMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)

	if (*req).Method == "OPTIONS" {
		return
	}
	if (*req).Method=="POST"{
		var memberish Membresia
		_ = json.NewDecoder(req.Body).Decode(&memberish)
		abrirConexionDB()
		agregarDatosBD(memberish)
		resultadosQuery(tabla)
		terminarConexion()
		data.Data=membresias
		json.NewEncoder(w).Encode(data)
	}
}
func UpdateMembershipEndpoint(w http.ResponseWriter, req *http.Request){
	//params := mux.Vars(req)
		var memberish Membresia
		params := mux.Vars(req)
		_ = json.NewDecoder(req.Body).Decode(&memberish)
		abrirConexionDB()
		actualizarDatosBD(memberish,params["id"])
		terminarConexion()
		data.Data=membresias
		json.NewEncoder(w).Encode(data)
}
func DeleteMembershipEndpoint(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	params := mux.Vars(req)
	abrirConexionDB()
	eliminarDatosBD(strconv.ParseInt(params["id"],10,64) )
	terminarConexion()
	data.Data=membresias
	json.NewEncoder(w).Encode(data)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func Iniciar() {
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/memberships", GetMembershipsEndpoint).Methods("GET")
	router.HandleFunc("/memberships/{id}", GetMembershipEndpoint).Methods("GET")
	router.HandleFunc("/memberships/agregar", CreateMembershipEndpoint).Methods("POST")
	router.HandleFunc("/memberships/actualizar/{id}",UpdateMembershipEndpoint).Methods("POST")
	router.HandleFunc("/memberships/eliminar/{id}", DeleteMembershipEndpoint).Methods("DELETE")
	//
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS","DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

