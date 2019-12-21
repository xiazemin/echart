package geo

import (
	"github.com/go-echarts/go-echarts/charts"
	"net/http"
	"os"
	"log"
	"math/rand"
)

var (
	mapData = map[string]float32{
		"北京":   float32(rand.Intn(150)),
		"上海":   float32(rand.Intn(150)),
		"深圳":   float32(rand.Intn(150)),
		"辽宁":   float32(rand.Intn(150)),
		"青岛":   float32(rand.Intn(150)),
		"山西":   float32(rand.Intn(150)),
		"陕西":   float32(rand.Intn(150)),
		"乌鲁木齐": float32(rand.Intn(150)),
		"齐齐哈尔": float32(rand.Intn(150)),
	}

	guangdongData = map[string]float32{
		"深圳市": float32(rand.Intn(150)),
		"广州市": float32(rand.Intn(150)),
		"湛江市": float32(rand.Intn(150)),
		"汕头市": float32(rand.Intn(150)),
		"东莞市": float32(rand.Intn(150)),
		"佛山市": float32(rand.Intn(150)),
		"云浮市": float32(rand.Intn(150)),
		"肇庆市": float32(rand.Intn(150)),
		"梅州市": float32(rand.Intn(150)),
	}

	shantouData = map[string]float32{
		"澄海区": float32(rand.Intn(150)),
		"潮阳区": float32(rand.Intn(150)),
		"潮南区": float32(rand.Intn(150)),
		"南澳县": float32(rand.Intn(150)),
	}
)
func GeoHandler(w http.ResponseWriter, _ *http.Request) {

	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-示例图(effectScatter)"})
	geo.Add("geo", charts.ChartType.EffectScatter, mapData,
		charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"})
	f, err := os.Create("geo.html")
	if err != nil {
		log.Println(err)
	}
	geo.Render(w, f) // Render 可接收多个 io.Writer 接口

}
