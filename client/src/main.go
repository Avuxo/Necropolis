package main

import (
	"flag"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)
/*
Check if file exists
*/
func fileExists(path string) (bool){
	_, err := os.Stat(path) // stat the file
	exists := !os.IsNotExist(err) // does the file exist?
	return exists
}
/*
Configure the project
This includes creation of the package directory, and the initial setup of the osiris file.
*/
func setupProject(){
	os.Mkdir("packages", 0700)
	out, err := os.Create("osiris")
	if err != nil{
		log.Fatal(err)
	}
	out.Close()
}

/*
Append line to file
Used for the addition of files to OSIRIS and to add to .gitignore if necessary
*/
func fileAppend(line string, path string){
	// open file in append mode
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	// append line to file
	_,err = file.WriteString(line)
	if err != nil{
		log.Fatal(err)
	}
}

/*
Install packages
Download the file from the server and put it where it should be
*/
func installPkg(pkgName string, addToOsiris bool){
	if fileExists("./packages/" + pkgName){
		fmt.Printf("Package %s is already installed\n", pkgName)
		return
	}
	// download the file
	res, err := http.Get("http://localhost:8080/packages/" + pkgName)

	if err != nil{
		log.Fatal(err)
	}
	
	defer res.Body.Close()
	file, err := ioutil.ReadAll(res.Body) // write the response into a file
	if err != nil{
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./packages/" + pkgName, file, 0444) // create the library file
	if err != nil{
		log.Fatal(err)
	}
	if addToOsiris == true{ // user wants add to the osiris file?
		fileAppend("pkgName", "./osiris")
	}

	
}

func main(){
	// command line flags
	pkgFlag := flag.String("install", "", "The package you wish to install")
	reviveFlag := flag.Bool("revive", true, "Install the packages for the program from the osiris file")
	flag.Parse()

	setupProject() // create the packages dir and config project

	if *pkgFlag != ""{ // check if the install flag was set
		installPkg(*pkgFlag, true)
	}

	if *reviveFlag{
		readOsirisFile()
	}
}
