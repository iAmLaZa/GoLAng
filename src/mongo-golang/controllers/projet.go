package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iAmLaZa/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProjetController struct {
	session *mgo.Session
}

func NewProjetController(s *mgo.Session) *ProjetController {
	return &ProjetController{s}
}
func (pc ProjetController) GetProjet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	pr := models.Projet{}

	if err := pc.session.DB("mongo-golang").C("projets").FindId(oid).One(&pr); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(pr)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
func (pc ProjetController) CreateProjet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pr := models.Projet{}
	json.NewDecoder(r.Body).Decode(&p)
	pr.Id = bson.NewObjectId()
	pc.session.DB("mongo-golang").C("projets").Insert(pr)
	uj, err := json.Marshal(pr)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
func (pc UserController) DeleteProjet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.IsObjectIdHex(id)
	if err := pc.session.DB("mongo-golang").C("projets").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Projet", oid, "\n")
}
