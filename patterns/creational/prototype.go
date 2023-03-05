package creational

import "fmt"

type ShirtType int

const (
	WhiteShirtType ShirtType = iota
	BlackShirtType
)

type Tshirt struct {
	color string
	size  string
}

var white = &Tshirt{
	color: "white",
	size:  "M",
}

var black = &Tshirt{
	color: "black",
	size:  "M",
}

func GetClone(t ShirtType) (Tshirt, error) {
	switch t {
	case WhiteShirtType:
		return *white, nil
	case BlackShirtType:
		return *black, nil
	default:
		return Tshirt{}, fmt.Errorf("invalid shirt type")
	}

}
