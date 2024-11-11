package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Specify 'producer' or 'consumer'")
	}
	
	switch os.Args[1] {
	case "producer":
		cmd := exec.Command("go", "run", "producer/producer.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "consumer":
		cmd := exec.Command("go", "run", "consumer/consumer.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	default:
		log.Fatal("Invalid argument. Use 'producer' or 'consumer'")
	}
}