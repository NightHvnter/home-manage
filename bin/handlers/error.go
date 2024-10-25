package handlers

import (
	"log"
)

type error struct {
	err string
}

func (e error) handleFatal() {
	log.Fatal(e.err)
}

func (e error) handleWarning() {
	log.Println(e.err)
}
