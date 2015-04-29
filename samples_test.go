package ps

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/llgcode/draw2d/draw2dimg"
)

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

func init() {
	os.Mkdir("result", 0666)
}

func TestTiger(t *testing.T) {
	i := image.NewRGBA(image.Rect(0, 0, 600, 800))
	gc := draw2dimg.NewGraphicContext(i)
	gc.Translate(0, 380)
	gc.Scale(1, -1)
	gc.Translate(0, -380)
	src, err := os.OpenFile("samples/tiger.ps", 0, 0)
	if err != nil {
		return
	}
	defer src.Close()
	bytes, err := ioutil.ReadAll(src)
	reader := strings.NewReader(string(bytes))
	interpreter := NewInterpreter(gc)
	interpreter.Execute(reader)
	saveToPngFile("result/TestPostscript.png", i)
}
