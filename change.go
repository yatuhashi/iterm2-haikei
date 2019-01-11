package main

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

func ImgResize(Img string) (image.Image, error) {
	pos := strings.LastIndex(Img, ".")
	if Img[pos:] != ".jpg" {
		return nil, errors.New("This file is not jpeg")
	}
	imgRawdata, err := os.Open(Img)
	defer imgRawdata.Close()
	if err != nil {
		return nil, err
	}
	imgData, _, err := image.Decode(imgRawdata)
	if err != nil {
		return nil, err
	}
	reImg := resize.Resize(960, 0, imgData, resize.Lanczos3)
	return reImg, nil
}

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

func Synthesis(Img1data, Img2data image.Image) {

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

func main() {
	rand.Seed(time.Now().Unix())
	pics := dirwalk("/Users/yatuhashi/Pictures/Fate/Haikei")
	var reimg1 image.Image
	var reimg2 image.Image
	var err error
	for {
		Img1 := fmt.Sprint(pics[rand.Intn(len(pics))])
		reimg1, err = ImgResize(Img1)
		if err == nil {
			break
		}
	}
	for {
		Img2 := fmt.Sprint(pics[rand.Intn(len(pics))])
		reimg2, err = ImgResize(Img2)
		if err == nil {
			break
		}
	}
	Synthesis(reimg1, reimg2)
}
