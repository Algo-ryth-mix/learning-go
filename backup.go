package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

type persistence struct {
	Counter int
	Payload map[string]string
}

func createImage() persistence {
	p := persistence{Counter: counter, Payload: urlStore}
	return p
}

func backupTask() {
	fmt.Print("Saving\n")
	image := createImage()

	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	err := e.Encode(image)

	if err != nil {
		fmt.Printf("[WARN!!!] Failed to create Image!(%v)\n", err)
		return
	}

	err = ioutil.WriteFile("persistent/glob.bin", b.Bytes(), 0644)
	if err != nil {
		fmt.Print("[WARN!!!] Failed to write Image!\n")
		return
	}
}

func readBackup() {

	f, err := os.Open("persistent/glob.bin")
	if err != nil {
		fmt.Print("could not read save file, no data was recovered\n")
		return
	}
	d := gob.NewDecoder(f)

	var p persistence

	err = d.Decode(&p)
	if err != nil {
		panic("Glob input format wrong\n")
	}

	urlStore = p.Payload
	counter = p.Counter

}
