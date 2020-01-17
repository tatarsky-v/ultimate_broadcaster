package main

import (
	"net/http"
	"os/exec"
	"log"
	"math/rand"
	"io/ioutil"
)

const SAMPLES_FOLDER = "./samples"
const PLAYER = "mpg123"

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listen on port :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	go (func(){
		cmd := exec.Command(PLAYER, randomFile())
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})()
}

func randomFile() string {
 	files, err := ioutil.ReadDir(SAMPLES_FOLDER)
 	if err != nil {
	  log.Fatal(err)
  }
  filename := SAMPLES_FOLDER + "/" + files[rand.Intn(len(files) - 1)].Name()
  log.Println("Play", filename)

  return filename
}