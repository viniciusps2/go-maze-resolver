package main

import (
    "fmt"
    "os"
    "image"
    "image/draw"
    "image/png"
    "github.com/google/gxui"
    "github.com/google/gxui/drivers/gl"
    "github.com/google/gxui/themes/dark"
    "github.com/viniciusps2/mazeresolver/node"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func openImg(imgPath string) *draw.Image {
    fimg,err := os.Open(imgPath)
    defer fimg.Close()
    checkErr(err)
    img,err := png.Decode(fimg)
    checkErr(err)
    // return &img
    
    dimg, ok := img.(draw.Image)
    if !ok {
        panic("cant covert to rgb")
    }
    return &dimg
}

func viewImg(m *draw.Image) func (driver gxui.Driver) {
    return func (driver gxui.Driver) {
        width, height := (*m).Bounds().Max.Y, (*m).Bounds().Max.X
        fmt.Println("==",width, height)
        theme := dark.CreateTheme(driver)
        img := theme.CreateImage()
        window := theme.CreateWindow(width, height, "Image viewer")
        texture := driver.CreateTexture(*m, 1.0)
        img.SetTexture(texture)
        window.AddChild(img)
        window.OnClose(driver.Terminate)
    }
}
func main() {
    img := openImg("fixtures/maze.png")
    n := node.NewNode(img, image.Point{20,80})
    n.SearchRelates()
    // handle(img, image.Point{20,80}, 0)
    // handle(img, image.Point{20,80}, 90)
    // handle(img, image.Point{20,80}, 270)
    // handle(img, image.Point{20,80}, 180)
    gl.StartDriver(viewImg(img))
}
