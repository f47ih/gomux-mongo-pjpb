package main

import "fmt"
import "go.mongodb.org/mongo-driver/mongo"
import "github.com/f47ih/gomux-mongo-pjpb/config"

func main() {
	fmt.Println("Hello")

	_, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}
	
}