package utils

import (
	"encoding/json"
    "bytes"
	"fmt"
	"github.com/golang/freetype"
    "github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
    "embed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strings"
)


const (
	FreeMono   = "FreeMono.ttf"
	FreeSans   = "FreeSans.ttf"
	UbuntuMono = "UbuntuMono.ttf"
)


var (
	// configuration models.Config
	// fontFile      string

	// go:embed static/*
	static embed.FS

	// go:embed views/*
	// views embed.FS
)

func GetJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

// func AssertType(files interface{}) interface{} {
//     m, ok := files.(map[string]interface{})
//     if !ok {
//         return fmt.Errorf("want type map[string]interface{};")
//     }
//     for k, v := range m {
//         fmt.Println(k, "=>", v)
//     }
//     return m
// }

func GetFileName(files map[string]interface{}) string {
    keys := make([]string, 0, len(files))
    for k := range files {
        keys = append(keys, k)
    }
    var filename string
    for i := 0; i<len(keys); i++ {
        if i>=0 && i<len(keys) {
            filename = keys[i]
        }
    }
    return filename
}

func getFontMap(f string) string {
	if f == "FreeSans" {
		return FreeSans
	} else if f == "FreeMono" {
		return FreeMono
	} else if f == "UbuntuMono" {
		return UbuntuMono
	}
	return UbuntuMono
}

func loadFont(fn string) (*truetype.Font, error) {
	fontFile := fmt.Sprintf("static/fonts/%s", getFontMap(fn))
	fontBytes, err := static.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GenerateImage(textContent string, fgColorHex string, bgColorHex string, fontSize float64) ([]byte, error) {

	fgColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	if len(fgColorHex) == 7 {
		_, err := fmt.Sscanf(fgColorHex, "#%02x%02x%02x", &fgColor.R, &fgColor.G, &fgColor.B)
		if err != nil {
			log.Println(err)
			fgColor = color.RGBA{0x2e, 0x34, 0x36, 0xff}
		}
	}

	bgColor := color.RGBA{0x30, 0x0a, 0x24, 0xff}
	if len(bgColorHex) == 7 {
		_, err := fmt.Sscanf(bgColorHex, "#%02x%02x%02x", &bgColor.R, &bgColor.G, &bgColor.B)
		if err != nil {
			log.Println(err)
			bgColor = color.RGBA{0x30, 0x0a, 0x24, 0xff}
		}
	}

	loadedFont, err := loadFont("UbuntuMono")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	code := strings.Replace(textContent, "\t", "    ", -1) // convert tabs into spaces
	text := strings.Split(code, "\n")                      // split newlines into arrays

	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	rgba := image.NewRGBA(image.Rect(0, 0, 1200, 630))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(loadedFont)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 50
	textYOffset := 10 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return nil, err
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, rgba); err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}
