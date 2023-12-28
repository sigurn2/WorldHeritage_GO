package main

import (
	"github.com/schollz/progressbar/v3"
	"github.com/sigurn2/WorldHeritage_GO/api"
	"github.com/sigurn2/WorldHeritage_GO/data"
	"log"
)

func main() {
	f, err := api.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			println(err)
		}
	}()
	heritages, _ := api.GetHeritages(f)
	bar := progressbar.Default(int64(len(heritages)))
	if data.AttributeList != nil {
		for _, v := range data.AttributeList {
			if api.AddCol(f, v) {
				for k, heritage := range heritages {
					ans := api.GetHeritageOutput(v, heritage)
					api.WriteValue(f, k+2, ans)
					bar.Add(1)
				}
			}
		}
	}
}
