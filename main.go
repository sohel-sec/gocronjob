package main

import (
    "fmt"
    "time"

    "github.com/jasonlvhit/gocron"
)

func task() {
    fmt.Println("Executing task at", time.Now())
    // Task logic goes here
}

func otherFunction() {
    fmt.Println("Executing otherFunction at", time.Now())
    // Other function logic goes here
}

func main() {
    // Create a new scheduler
    s := gocron.NewScheduler()

    // Schedule the task to run every 2 seconds
    s.Every(2).Seconds().Do(task)

    // Start the scheduler in a separate goroutine
    go s.Start()

    // Call the otherFunction directly from the main function
    otherFunction()

    // Keep the main goroutine running to allow the scheduled task to execute continuously
    select {}
}
