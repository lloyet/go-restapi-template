package fooController

import (
	fooModel "api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	b := bson.M{}
	name := query["name"]
	if len(name) > 0 {
		b = bson.M{"name": name[0]}
	}
	people := fooModel.Search(b)
	if js, err := json.Marshal(people); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	people := fooModel.FindById(params["id"])
	if js, err := json.Marshal(people); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var p fooModel.People
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := fooModel.Register(p)
	if js, err := json.Marshal(id); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fooModel.DeleteById(params["id"])
	w.Write([]byte("Ok"))
}
