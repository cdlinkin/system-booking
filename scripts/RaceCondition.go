package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type BookingRequest struct {
	UserID     int `json:"user_id"`
	ResourceID int `json:"resource_id"`
}

const (
	resourceID      = 1
	goroutinesCount = 50
)

func main() {
	fmt.Printf("Test/Race Condition: %d goroutines and %d resource.\n", goroutinesCount, resourceID)

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	doneBookings := 0

	for userID := 1; userID <= goroutinesCount; userID++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			time.Sleep(time.Duration(userID) * 10 * time.Millisecond)

			reqBody := BookingRequest{
				UserID:     userID,
				ResourceID: resourceID,
			}

			jsonData, _ := json.Marshal(reqBody)

			resp, err := http.Post("http://localhost:9090/api/bookings", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Printf("User %d: network error\n", userID)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode == 201 {
				mu.Lock()
				doneBookings++
				mu.Unlock()
				fmt.Printf("User %d: done. resource %d booked\n", userID, resourceID)
			} else {
				fmt.Printf("User %d: fail %d (%s)\n", userID, resp.StatusCode, string(body))
			}
		}(userID)
	}

	wg.Wait()

	fmt.Println()
	fmt.Printf("Result RaceCondition/Test:\n")
	fmt.Printf("Total attempts: %d\n", goroutinesCount)
	fmt.Printf("Done bookings: %d - bug\n", doneBookings)
	fmt.Printf("Race Condition: %d users have received one resource\n", doneBookings)
}
