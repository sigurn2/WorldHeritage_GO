package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/schollz/progressbar/v3"
	"github.com/sigurn2/WorldHeritage_GO/api"
	"github.com/sigurn2/WorldHeritage_GO/data"
	"github.com/xuri/excelize/v2"
	"log"
	"sync"
	"time"
)

type TaskParams struct {
	p1 *excelize.File
	p2 int
	p3 data.Attribute
	p4 string
	p5 *progressbar.ProgressBar
}

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

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(20, func(i interface{}) {
		writeCell(i)
		defer wg.Done()
	})
	defer p.Release()
	heritages, _ := api.GetHeritages(f)
	bar := progressbar.Default(int64(len(heritages)))
	if data.AttributeList != nil {
		for _, v := range data.AttributeList {
			if api.AddCol(f, v) {
				for k, heritage := range heritages {
					wg.Add(1)
					_ = p.Invoke(TaskParams{
						p1: f,
						p2: k,
						p3: v,
						p4: heritage,
						p5: bar,
					})
				}
			}
		}
	}
	wg.Wait()
	fmt.Printf("\nrunning goroutines: %d\n", p.Running())
}
func writeCell(i interface{}) {
	n := i.(TaskParams)
	time.Sleep(500 * time.Millisecond)
	ans := api.GetHeritageOutput(n.p3, n.p4)
	api.WriteValue(n.p1, n.p2+2, ans)
	n.p5.Add(1)
}
