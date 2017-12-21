package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"github.com/fatih/color"
	"net/http"
	"os"
)

var help bool

func init() {
	flag.BoolVarP(&help, "help", "h", false, "Display this help message")
	flag.Parse()
}

func main() {
	if help == true {
		PrintHelpMessage()
	}

  links := map[string]string{
    "BTC": "https://ethereumprice.org/btc",
    "ETH": "https://ethereumprice.org",
    "LTC": "https://ethereumprice.org/ltc",
    "XMR": "https://ethereumprice.org/xmr/",
  }

  resp := map[string]string{}

  blue := color.New(color.FgBlue).SprintFunc()
  red := color.New(color.FgRed).SprintFunc()
  yellow := color.New(color.FgYellow).SprintFunc()
  green := color.New(color.FgGreen).SprintFunc()

  for k, v := range links {
    resp[k] = GetData(v)
  }

  fmt.Printf("%s: %s\n", blue("BTC"), resp["BTC"])
  fmt.Printf("%s: %s\n", red("ETH"), resp["ETH"])
  fmt.Printf("%s: %s\n", yellow("LTC"), resp["LTC"])
  fmt.Printf("%s: %s\n", green("XMR"), resp["XMR"])

}

func GetData(link string) string {
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

	data := fmt.Sprintf("$%s (%s%%)", scrape.Text(price), scrape.Text(change))
	return data
}

func PrintHelpMessage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
