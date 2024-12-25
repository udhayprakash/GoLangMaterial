package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "log"
)

const apiUrl = "http://api.open-notify.org/iss-now.json"

type ISSResponse struct {
    Timestamp    int64 `json:"timestamp"`
    ISSPosition  struct {
        Latitude  string `json:"latitude"`
        Longitude string `json:"longitude"`
    } `json:"iss_position"`
    Message string `json:"message"`
}

func main() {
    // Task 1: Parsing sampleResponse
    sampleResponse := []byte(`{"timestamp": 1710861415, "iss_position": {"latitude": "22.9866", "longitude": "152.7521"}, "message":"success"}`)
    var issResponse ISSResponse
    if err := json.Unmarshal(sampleResponse, &issResponse); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Latitude: %s, Longitude: %s\n", issResponse.ISSPosition.Latitude, issResponse.ISSPosition.Longitude)

    // Task 2: Making a request to the live API
    response, err := http.Get(apiUrl)
    if err != nil {
        log.Fatal("Error making request:", err)
    }
    defer response.Body.Close()

    var liveResponse ISSResponse
    if err := json.NewDecoder(response.Body).Decode(&liveResponse); err != nil {
        log.Fatal("Error decoding JSON response:", err)
    }

    if liveResponse.Message == "success" {
        fmt.Printf("Live Latitude: %s, Longitude: %s\n", liveResponse.ISSPosition.Latitude, liveResponse.ISSPosition.Longitude)
    } else {
        log.Fatal("Error: ", liveResponse.Message)
    }
}
