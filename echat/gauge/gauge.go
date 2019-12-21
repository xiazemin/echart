package gauge

import (
	"github.com/go-echarts/go-echarts/charts"
	"math/rand"
	"net/http"
	"os"
	"log"
)

func GaugeHandler(w http.ResponseWriter, _ *http.Request)  {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-示例图"})
	m := make(map[string]interface{})
	m["工作进度"] = rand.Intn(50)
	gauge.Add("gauge", m)
	f, err := os.Create("geo.html")
	if err != nil {
		log.Println(err)
	}
	gauge.Render(w, f) // Render 可接收多个 io.Writer 接口
}
