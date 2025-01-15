package main

import (
	"net/http"
	"strings"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestGetAllGroceries(t *testing.T) {
	resp, err := http.Get(baseURL + "/groceries")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestGetGroceryByID(t *testing.T) {
	resp, err := http.Get(baseURL + "/groceries/1")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestAddGrocery(t *testing.T) {
	payload := `{"id": 3, "name": "Oranges", "price": 2.49}`
	resp, err := http.Post(baseURL+"/groceries", "application/json", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestUpdateGrocery(t *testing.T) {
	payload := `{"id": 1, "name": "Apples", "price": 1.49}`
	req, err := http.NewRequest(http.MethodPut, baseURL+"/groceries/1", strings.NewReader(payload))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send PUT request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestDeleteGrocery(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/groceries/1", nil)
	if err != nil {
		t.Fatalf("Failed to create DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

/*
go test -v


Assignment:
	Based on postman collection
		Autogenerate these test cases using `postman-to-code` tool

*/
