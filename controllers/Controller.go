package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio-backend/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const ConnectionString = "mongodb://localhost:27017"
const DatabaseName = "portfolio-backend"
const CollectionName = "projects"

var Collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	Collection = client.Database(DatabaseName).Collection(CollectionName)
}

func insertProject(project *models.Project) {
	inserted, err := Collection.InsertOne(context.Background(), project)
	if err != nil {
		panic(err)
	}
	fmt.Print("project inserted successfully------\n", inserted)
}

func updateProject(movieId string, project models.Project) *mongo.UpdateResult {
	oid, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": project}
	updated, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	return updated
}

func deleteProject(movieId string) *mongo.DeleteResult {
	oid, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": oid}
	deletedDoc, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	return deletedDoc
}
func getAllProjects() []bson.M {
	cur, err := Collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	var projects []bson.M
	for cur.Next(context.Background()) {
		var project bson.M
		err := cur.Decode(&project)
		if err != nil {
			panic(err)
		}
		projects = append(projects, project)
	}
	defer cur.Close(context.Background())
	return projects
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allProjects := getAllProjects()
	json.NewEncoder(w).Encode(allProjects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var project models.Project
	_ = json.NewDecoder(r.Body).Decode(&project)
	insertProject(&project)
	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteProject(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var project models.Project
	_ = json.NewDecoder(r.Body).Decode(&project)
	updateProject(params["id"], project)
	json.NewEncoder(w).Encode(params["id"])
}
