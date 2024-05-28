package handlers

import (
	"sipinter-api/models"
	"sipinter-api/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	var services []models.Service
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := utils.ServiceCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var service models.Service
		cursor.Decode(&service)
		services = append(services, service)
	}

	json.NewEncoder(w).Encode(services)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var service models.Service

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := utils.ServiceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	json.NewDecoder(r.Body).Decode(&service)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	service.ID = primitive.NewObjectID()
	_, err := utils.ServiceCollection.InsertOne(ctx, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(service)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var service models.Service
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := utils.ServiceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&service)
	_, err = utils.ServiceCollection.ReplaceOne(ctx, bson.M{"_id": id}, service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(service)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := utils.ServiceCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
