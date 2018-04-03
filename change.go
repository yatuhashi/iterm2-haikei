package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func Shiro() image.Image {
	x := 0
	y := 0
	width := 1920
	height := 1200
	// 1948x1414

	// RectからRGBAを作る(ゼロ値なので黒なはず)
	img := image.NewRGBA(image.Rect(x, y, width, height))
	return img
}

func Gousei(Img1data, Img2data image.Image) {

	ShiroImgdata := Shiro()

	startPointLogo := image.Point{0, 0}
	logoRectangle := image.Rectangle{startPointLogo, startPointLogo.Add(Img1data.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, ShiroImgdata.Bounds().Size()}

	rgba := image.NewRGBA(originRectangle)
	draw.Draw(rgba, originRectangle, ShiroImgdata, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, logoRectangle, Img1data, image.Point{0, 0}, draw.Over)

	//オリジナル画像上のどこからlogoイメージを重ねるか
	//これだと左上
	startPointLogo = image.Point{960, 0}
	logoRectangle = image.Rectangle{startPointLogo, startPointLogo.Add(Img2data.Bounds().Size())}
	originRectangle = image.Rectangle{image.Point{0, 0}, ShiroImgdata.Bounds().Size()}

	draw.Draw(rgba, logoRectangle, Img2data, image.Point{0, 0}, draw.Over)

	out, err := os.Create("/Users/yatuhashi/Pictures/Fate/Haikei/result.jpg")
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(out, rgba, &opt)
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}

func ImgResize(Img1, Img2 string) (image.Image, image.Image) {
	img1rawdata, err := os.Open(Img1)
	if err != nil {
		fmt.Println(err)
	}
	img2rawdata, err := os.Open(Img2)
	if err != nil {
		fmt.Println(err)
	}
	img1data, _, err := image.Decode(img1rawdata)
	if err != nil {
		fmt.Println(err)
	}
	img2data, _, err := image.Decode(img2rawdata)
	if err != nil {
		fmt.Println(err)
	}
	reimg1 := resize.Resize(960, 0, img1data, resize.Lanczos3)
	reimg2 := resize.Resize(960, 0, img2data, resize.Lanczos3)
	return reimg1, reimg2
}

func main() {
	rand.Seed(time.Now().Unix())
	pics := dirwalk("/Users/yatuhashi/Pictures/Fate/Haikei")
	Img1 := fmt.Sprint(pics[rand.Intn(len(pics))])
	Img2 := fmt.Sprint(pics[rand.Intn(len(pics))])
	// fmt.Println(Img1, Img2)
	reimg1, reimg2 := ImgResize(Img1, Img2)
	Gousei(reimg1, reimg2)
}
