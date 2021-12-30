package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	in := os.Getenv("INPUT_FILE")
	if len(in) == 0 {
		in = ".env"
	}
	out := os.Getenv("INPUT_OUTPUT")
	if len(out) == 0 {
		out = in
	}
	sec := os.Getenv("INPUT_SECRETS")
	if len(sec) == 0 {
		log.Fatal("required input 'secrets' is missing")
	}
	var data interface{}
	err := json.Unmarshal([]byte(sec), &data)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse secrets from json '%v': %w", sec, err))
	}
	t, err := template.ParseFiles(in)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse input file '%v': %w", in, err))
	}
	f, err := os.Create(out)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create output file '%v': %w", out, err))
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to execute template '%v': %w", out, err))
	}
}
