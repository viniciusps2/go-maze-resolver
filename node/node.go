package node

import (
    "fmt"
    "image"
    "math"
    "image/color"
    "image/draw"
)

const (
    degree = float64(math.Pi / 180)
)
var (
    blue = color.RGBA{0, 0, 255, 255}
    white = color.RGBA{255, 255, 255, 255}
)

type Node struct {
	start, end image.Point
	img *draw.Image
	relates []*Node
}

func NewNode(img *draw.Image, start image.Point) *Node {
	return &Node{
		img: img,
		start: start,
	}
}

func (n *Node) SearchRelates() {
    // n := node.NewNode(img, image.Point{20,80})
	angles := []float64{0, 90, 180, 270}
	for _, angle := range angles {
		x,y := n.WalkAt(angle)
		d := Distance(float64(n.start.X), float64(n.start.Y), float64(x), float64(y))
		fmt.Println("d--", x, y, d)
		if d > 20 {
			fmt.Println("new--", x, y, d)
			relate := n.NewRelate(image.Point{x, y})
			n.AddRelate(relate)
			n.SearchRelates()
		}
	}
}

func (n *Node) NewRelate(start image.Point) *Node {
	return NewNode(n.img, start)
}

func (n *Node) AddRelate(relate *Node) {
	n.relates = append(n.relates, relate)
}

func (n Node) WalkAt(angle float64) (int, int) {
    t := math.Tan(angle * degree)
    var x0, y0 int
    var factorX, factorY float64
    switch angle {
        case 0:
            factorY = 1
        case 90:
            factorX = 1
        case 180:
            factorY = -1
        case 270:
            factorX = -1
        default:
            factorX = 1
            factorY = t
    }
    // fmt.Println("t", t, factorX, factorY)
    for i := 0; i < (*n.img).Bounds().Max.X; i++ {
        x := n.start.X + int(float64(i) * factorX)
        y := n.start.Y + int(float64(i) * factorY)
        if !n.verify(x0, x, y0, y) {
            break
        }
        x0 = x
        y0 = y
    }
    return x0, y0
}

func (n *Node) verify(x0, x, y0, y int) bool {
    lineSize := 1
    c := (*n.img).At(x,y)
    // fmt.Println(x,y, c)
    if colorEquals(white, c) || colorEquals(blue, c) {
       rect := image.Rect(x, y, x+lineSize, y+lineSize)
       draw.Draw(*n.img, rect, &image.Uniform{blue}, image.ZP, draw.Src)
       return true
    }
    return false
}

