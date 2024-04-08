package repository

import (
	"Project-RestAPIWithGoMongoDBAndGorillaMux/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
	options.Client().ApplyURI("mongodb+srv://kumarnikhil9687:4891Nch06!@cluster0.cp91rfw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("Error while connecting to MongoDB", err)
	}

	log.Println("MongoDB connected Successfully")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error while Pinging to MongoDB", err)
	}

	log.Println("MongoDB Ping Successfully")

	return mongoTestClient
}

func TestMongoOperations (t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data

	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	coll := mongoTestClient.Database("Project1DB").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	//Insert Employee 1 Data

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name: "Abs Go",
			Department: "Programming",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}

		t.Log("Insert 1 successful", result)
	})

	t.Run("Insert emp 2", func(t *testing.T) {
		emp := model.Employee {
			Name: "Nikhil",
			Department: "Go Learner",
			EmployeeID: emp2,
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert 2 failed", err)
		}
		t.Log("insert 2 successful", result)
	})

	t.Run("Get Emp 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Get emp 1 failed", err)
		}

		t.Log("Get emp 1 passed", result)
	})
}