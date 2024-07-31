package main

import (
	"errors"
	"fmt"
	"log"

	courseSdk "github.com/zchelalo/go_microservices_course_sdk/course"
)

func main() {
	courseTransport := courseSdk.NewHTTPClient("http://localhost:8002", "")

	course, err := courseTransport.Get("9c25d2b1-0865-49b8-af2d-3ef840b5f64c")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			log.Fatal("course not found: ", err.Error())
		}

		log.Fatal(err)
	}

	fmt.Println(course)
}
