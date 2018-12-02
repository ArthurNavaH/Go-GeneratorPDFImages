package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

type config struct {
	ImagesInput string  `json:"imagesInput"`
	FileOutput  string  `json:"fileOutput"`
	WidthImage  float64 `json:"widthImage"`
	HeightImage float64 `json:"heightImage"`
	TopImage    float64 `json:"topImage"`
	LeftImage   float64 `json:"leftImage"`

	SheetStyle string `json:"sheetStyle"`
	ExtImage   string `json:"extImage"`
}

func main() {
	var config config
	// config.imagesInput = "./images"
	// config.fileOutput = "MiJardin.pdf"
	// config.widthImage = 190.0
	// config.heightImage = 190.0
	// config.topImage = 10.0
	// config.leftImage = 40.0
	// config.sheetStyle = "A4"
	// config.extImage = "jpg"

	cj, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(cj, &config)
	if err != nil {
		fmt.Println(err)
	}

	files, _ := ioutil.ReadDir(config.ImagesInput)

	pdf := gofpdf.New("P", "mm", config.SheetStyle, "")

	for i := 0; i < len(files); i++ {
		pdf.AddPage()
		file := fmt.Sprintf("%s/%d.%s", config.ImagesInput, i+1, config.ExtImage)
		pdf.Image(file, config.LeftImage, config.TopImage, config.WidthImage, config.HeightImage, false, "", 0, "")
	}
	err = pdf.OutputFileAndClose(config.FileOutput)
	if err != nil {
		fmt.Println(err)
	}
}
