package main

/* 
#cgo LDFLAGS: -L /usr/local/lib  -lboost_program_options  -lboost_thread -lboost_system -lboost_filesystem
#cgo LDFLAGS: -lboost_chrono -lboost_date_time -lboost_coroutine -lboost_context
#cgo CXXFLAGS: -std=c++11

#include "wrap_point.hxx"
*/
import "C"

import "fmt"

func main() {
	fmt.Println("Hi from Go, about to calculate distance in C++ ...")
	go func(msg string) {
        	fmt.Println(msg)
		distance := C.distance_between(1.0, 1.0, 2.0, 2.0)
		fmt.Printf("Go has result, distance is: %v\n", distance)
    	}("going")

	go func(msg string) {
        	fmt.Println(msg)
		distance := C.distance_between(1.0, 1.0, 2.0, 2.0)
		fmt.Printf("Go2 has result, distance is: %v\n", distance)
    	}("going2")
	fmt.Scanln()
    	fmt.Println("done")
}
