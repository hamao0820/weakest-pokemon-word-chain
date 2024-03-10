package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Pokemon struct {
	No   int    `json:"no"`
	Name string `json:"name"`
}

type Data struct {
	Data []Pokemon `json:"data"`
}

func main() {
	url := "https://pente.koro-pokemon.com/zukan/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	ulList := doc.Find(".ul2.zukan_img")

	pokemons := []Pokemon{}
	ulList.Each(func(i int, ul *goquery.Selection) {
		liList := ul.Find("li")
		liList.Each(func(j int, li *goquery.Selection) {
			p := Pokemon{}
			a := li.Find("a").First()
			noSpan := a.Find("span.no")
			nameSpan := a.Find("span.name")
			if noSpan.Length() == 0 || nameSpan.Length() == 0 {
				return
			}
			p.No, err = strconv.Atoi(noSpan.Text())
			if err != nil {
				panic(err)
			}
			p.Name = nameSpan.Text()
			pokemons = append(pokemons, p)
		})
	})

	d := Data{Data: pokemons}
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("data/zukan.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		panic(err)
	}
}
