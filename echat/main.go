package main

import (
	"github.com/go-echarts/go-echarts/charts"
	"net/http"
	"os"
	"log"
	"math/rand"
	"time"
	"github.com/xiazemin/echart/echat/geo"
	"github.com/xiazemin/echart/echat/heatMap"
	"github.com/xiazemin/echart/echat/page"
	"github.com/xiazemin/echart/echat/gauge"
)
var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63()) % 50)
	}
	return r
}
func handler(w http.ResponseWriter, _ *http.Request) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(w, f) // Render 可接收多个 io.Writer 接口
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/geo", geo.GeoHandler)
	http.HandleFunc("/heatMap", heatMap.HeatMapHandler)
	http.HandleFunc("/page", page.PagHandler)
	http.HandleFunc("/gauge", gauge.GaugeHandler)
	http.ListenAndServe(":8081", nil)
}