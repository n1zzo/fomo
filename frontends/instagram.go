package frontends

import (
	"image/draw"
	"image/png"
    "bufio"
    "flag"
    "fmt"
    "strings"
    "image"
    "io/ioutil"
    "os"

    . "github.com/n1zzo/fomo/data"

    "github.com/golang/freetype"
    "golang.org/x/image/font"
    "github.com/kr/pretty"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./fonts/WorkSans/WorkSans-Light.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 180, "font size in points")
	spacing  = flag.Float64("spacing", 1.3, "line spacing (e.g. 2 means double spaced)")
)

//func downloadImage(url string) (img image.Image) {
//    resp, err := client.Get(url)
//    defer resp.Body.Close()
//    image = Decode(resp.Body)
//    return image
//}

func makeStory(e Event) (story image.Image) {
    fontBytes, err := ioutil.ReadFile(*fontfile)
    if err != nil {
		fmt.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

    // TODO: Download image from event
    // TODO: Scale to portrait FullHD resolution

    // Initialize the context.
    fg, bg := image.White, image.Black
	rgba := image.NewRGBA(image.Rect(0, 0, 1080, 1920))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

    // Split text into lines
    lines := make([]string, 3)
    chars, l := 0, 0
    for _, word := range strings.Split(e.Name, " ") {
        lines[l] = lines[l] + " " + word
        chars += len(word)
        if chars > 8 {
            l += 1
            chars = 0
        }
    }

    pretty.Println(lines)

	// Draw the text.
    pt := freetype.Pt(100, 960)
    for _, l := range lines {
        c.DrawString(l, pt)
        pt.Y += c.PointToFixed(*size * *spacing)
    }

    return rgba
}

func saveStory(s image.Image) {
	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}

func Instagram(e Event) {
    // Create story image
    story := makeStory(e)

    // TODO: publish story on Instagram

    // Save story to file
    saveStory(story)
}
