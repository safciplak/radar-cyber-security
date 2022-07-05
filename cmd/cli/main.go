package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goware/urlx"

	"github.com/jessevdk/go-flags"
	jsonWrapper "github.com/radar-cyber-security/internal/json"
	"github.com/radar-cyber-security/pkg/numbersapi"
)

func main() {

	// clients
	numbersApiClient := numbersapi.New()

	// services

	jsonService := &jsonWrapper.Service{Numbers: numbersApiClient}

	// command line args
	var opts struct {
		InputFileName  string `short:"i" long:"input file name" description:"Input file name" required:"true"`
		OutputFileName string `short:"o" long:"output file name" description:"Output file name" required:"true"`
	}

	args := []string{
		"-i", "sample.csv",
		"-o", "sample.json",
	}

	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		fmt.Println(err)
	}

	// file reader
	file, err := os.Open(opts.InputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var jsonResponse = []interface{}{}

	var url string
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")

		t1 := items[0]
		ip := items[1]
		url = items[2]
		s := items[3]
		size, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(items[0]) // time
		// fmt.Println(items[1]) // ip
		// fmt.Println(items[2]) // url
		// fmt.Println(items[3]) // size

		url, err := urlx.Parse(url)
		if err != nil {
			log.Fatal(err)
		}

		text, err := jsonService.GetRandomText(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		time, e := time.Parse(time.RFC3339, t1)
		if e != nil {
			log.Fatal(e)
		}

		var URLItem jsonWrapper.URLItem
		URLItem.Scheme = url.Scheme
		URLItem.Host = url.Host
		URLItem.Path = url.Path
		URLItem.Opaque = url.Opaque

		var JSONItem jsonWrapper.JSONItem

		JSONItem.Ts = time.Unix()
		JSONItem.SourceIP = ip
		JSONItem.URLItem = URLItem
		JSONItem.Size = size
		JSONItem.Note = text

		jsonItem, err := json.Marshal(JSONItem)
		if err != nil {
			log.Fatal(err)
		}

		res := jsonWrapper.JSONItem{}
		json.Unmarshal([]byte(jsonItem), &res)

		jsonResponse = append(jsonResponse, res)
	}

	jsonFile, _ := json.MarshalIndent(jsonResponse, "", " ")

	_ = ioutil.WriteFile("test.json", jsonFile, 0644)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
