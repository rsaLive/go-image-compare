package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

func main() { //如果读取出错会panic,返回图像矩阵img
	s1 := "./images/1.png"
	s2 := "./images/2.png"

	img1 := readImage(s1)
	img2 := readImage(s2)

	if img1 == nil || img2 == nil {
		fmt.Println("Failed to read image(s)")
		return
	}

	if img1.Bounds() != img2.Bounds() {
		fmt.Println("Images are not of the same size")
		return
	}

	diff := compareImages(img1, img2)

	fmt.Printf(" score between %s and %s is %f\n", s1, s2, diff)
}

func readImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return img
}

func compareImages(img1, img2 image.Image) float64 {
	// 获取图片的边界
	bounds := img1.Bounds()
	// 获取图片的宽和高
	w, h := bounds.Max.X, bounds.Max.Y
	// 初始化差异值
	var diff float64
	// 遍历每个像素点
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// 获取第一个图片的像素点的RGBA值
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			// 获取第二个图片的像素点的RGBA值
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			// 计算每个像素点的差异值
			diff += math.Abs(float64(r1)-float64(r2)) / 0xffff
			diff += math.Abs(float64(g1)-float64(g2)) / 0xffff
			diff += math.Abs(float64(b1)-float64(b2)) / 0xffff
		}
	}
	// 计算图片的像素点总数
	nPixels := w * h
	// 计算图片的相似度得分，3是RPG颜色通道
	score := diff / (3 * float64(nPixels))
	return 1 - score
}
