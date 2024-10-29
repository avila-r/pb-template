package main

import "github.com/pocketbase/pocketbase"

func main() {
	pb := pocketbase.New()

	if err := pb.Start(); err != nil {
		panic(err.Error())
	}
}
