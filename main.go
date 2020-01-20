package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
)

const SAMPLES_FOLDER = "./samples"
const PLAYER = "mpg123"

func main() {
	http.HandleFunc("/random_phrase", randomPhrase)
	http.HandleFunc("/polly", polly)
	log.Println("Listen on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func randomPhrase(w http.ResponseWriter, r *http.Request) {
	go (func() {
		cmd := exec.Command(PLAYER, randomFile())
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})()
}

func polly(w http.ResponseWriter, r *http.Request) {
	go (func() {
    text := r.URL.Query().Get("text")
    if len(text) == 0 {
    	log.Println("Url Param 'text' is missing")
    	return
    }

		cmd := exec.Command(
			"aws", "polly", "synthesize-speech", "--output-format", "mp3", "--text",
			text, "--voice-id", "Maxim", "/tmp/voice.mp3")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		cmd = exec.Command("mpg123", "/tmp/voice.mp3")
		err = cmd.Run()
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
	filename := SAMPLES_FOLDER + "/" + files[rand.Intn(len(files)-1)].Name()
	log.Println("Play", filename)

	return filename
}
