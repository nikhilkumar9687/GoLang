package main

import (
	"Project-RestAPIWithGoMongoDBAndGorillaMux/usecase"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init () {
	//Load env
	err := godotenv.Load("C:\\Users\\Dell\\Downloads\\GoLang\\Project-RestAPIWithGoMongoDBAndGorillaMux\\repository\\.env")
	if err != nil {
		log.Fatal("env load fail", err)
	}

	//create mongo client
	mongoClient, err = mongo.Connect(context.Background(), 
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection failed", err)

	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping Failed", err)
	}

	log.Println("Ping successful")
}


func main() {
	

	//Load env
	err := godotenv.Load("C:\\Users\\Dell\\Downloads\\GoLang\\Project-RestAPIWithGoMongoDBAndGorillaMux\\repository\\.env")
	if err != nil {
		log.Fatal("env load fail", err)
	}
	//"disconect"

	defer mongoClient.Disconnect(context.Background())
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTIONS_NAME"))

	// Create a employee service
	empService := usecase.EmployeeService{MongoCollection: coll}
	r := mux.NewRouter()
	
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/employee",  empService.CreateEmployee).Methods(http.MethodPost)

	r.HandleFunc("/employee/{id}",  empService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee",  empService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee",  empService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}",  empService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee",  empService.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("Server running at 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}