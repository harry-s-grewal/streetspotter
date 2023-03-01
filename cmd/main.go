package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	montagne    = "https://ville.montreal.qc.ca/Circulation-Cameras/GEN274.jpeg"
	guystjaques = "https://ville.montreal.qc.ca/Circulation-Cameras/GEN379.jpeg"
)

func main() {
	go download(montagne, "./montagne/")
	go download(guystjaques, "./guystjaques/")

	<-make(chan int)
}

func download(url string, filename string) {
	for {
		err := downloadFile(url, filename)
		if err != nil {
			panic(err)
		}
		// to do: Make this check if the image is the same as the previous one and query more often
		time.Sleep(5 * time.Minute)
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	fileName = fileName + time.Now().Format("2006-01-02_15:04:05") + ".jpeg"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
