package usecase

import (
	"Project-RestAPIWithGoMongoDBAndGorillaMux/model"
	"Project-RestAPIWithGoMongoDBAndGorillaMux/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee (w http.ResponseWriter, r *http.Request){
	log.Println("LOG:In function CreateEmployee")
	fmt.Println("FMT:In function CreateEmployee")
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = uuid.NewString()
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	//insert EMp
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee inserted with ID", insertID, emp)
}

func (svc *EmployeeService) GetEmployeeByID (w http.ResponseWriter, r *http.Request){
	log.Println("LOG:In function GetEmployeeByID")
	fmt.Println("FMT:In function GetEmployeeByID")
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get Emp id

	empId := mux.Vars(r)["id"]
	log.Println("Employee ID", empId)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	
	emp, err := repo.FindEmployeeByID(empId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) GetAllEmployee (w http.ResponseWriter, r *http.Request){
	log.Println("LOG:In function GetAllEmployee")
	fmt.Println("FMT:In function GetAllEmployee")
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	
	emp, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) UpdateEmployeeByID (w http.ResponseWriter, r *http.Request){
	log.Println("LOG:In function UpdateEmployeeByID")
	fmt.Println("FMT:In function UpdateEmployeeByID")
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get Emp id

	empId := mux.Vars(r)["id"]
	log.Println("Employee ID", empId)

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid emp")
		res.Error = "Invalid emp"
		return
	}

	var updatedEmpData model.Employee

	if err := json.NewDecoder(r.Body).Decode(&updatedEmpData); err != nil {
        http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
        return
    }
	log.Println(updatedEmpData)
    // Update employee data in the repository
    repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
    count, err := repo.UpdateEmployeeByID(empId, &updatedEmpData)
    if err != nil {
        http.Error(w, "Failed to update employee data: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return success response
    response := map[string]interface{}{
        "message":      "Employee updated successfully",
        "modifiedCount": count,
    }
    json.NewEncoder(w).Encode(response)
}

func (svc *EmployeeService) DeleteEmployeeByID (w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get Emp id

	empId := mux.Vars(r)["id"]
	log.Println("Employee ID", empId)

	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid emp")
		res.Error = "Invalid emp"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid Body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empId
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	

	count, err := repo.DeleteEmployeeByID(empId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) DeleteAllEmployee (w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	
	emp, err := repo.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}