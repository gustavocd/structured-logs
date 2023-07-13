package main

import (
	"log"
	"os"
)

// START OMIT
func main() {
	defaultLogger := log.Default()
	defaultLogger.SetOutput(os.Stdout)
	log.Println("Hello from the std-log package!")

	logger := log.New(
		os.Stderr,
		"Go application: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)

	logger.Println("Hello from the custom logger!")
	logger.Printf("username=%s user_id=%s request_id:%s", "John Doe", "123", "abc123")
}

// END OMIT
