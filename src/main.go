package main

import (
	s "github.com/gabrielanjos/golang/service"

	_ "github.com/lib/pq"
)

func main() {
	s.ReadAndManipulateFile()
}
