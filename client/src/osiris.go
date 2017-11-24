package main

import (
	"bufio"
	"os"
	"log"
)

/*
Read the osiris file
Go through each package and install it
*/
func readOsirisFile(){
	file, err := os.Open("osiris") // open the osiris file
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	var osirisFile []string // array of lines
	scanner := bufio.NewScanner(file) // scanner to read the file

	for scanner.Scan(){ // read the entire file
		osirisFile = append(osirisFile, scanner.Text()) // append to array
	}
	
}
