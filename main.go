package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

const link string = "https://ethereumprice.org/btc/"

var help bool

func init() {
	flag.BoolVarP(&help, "help", "h", false, "Display this help message")
	flag.Parse()
}

func main() {
	if help == true {
		PrintHelpMessage()
	}
	data := GetData()
	fmt.Print(data)
}

func GetData() string {
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	price, _ := scrape.Find(root, scrape.ById("ep-price"))
	change, _ := scrape.Find(root, scrape.ById("ep-percent-change"))

	data := fmt.Sprintf("$%s (%s)", scrape.Text(price), scrape.Text(change))
	return data
}

func PrintHelpMessage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
