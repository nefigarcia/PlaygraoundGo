package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func(app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct{
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Go lang Accenture students",
		Version: "1.0.0",
	}
	
	out, err := json.Marshal(payload)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllStudents(w http.ResponseWriter, r *http.Request){
	var students []models.Student
	rd, _ :=time.Parse("2006-01-02", "2024-03-12")

	highlander := models.Student {
		ID:1,
		Name:"Nefi Garcia",
		Company: "Accenture",
		Band: "CL10",
		EnrrolledDate: rd,
	}
	students = append(students, highlander)

	out,err := json.Marshal(students)
	if err!= nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}