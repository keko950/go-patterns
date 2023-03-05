package creational

import (
	"testing"
)

func TestCloneWhiteTshirt(t *testing.T) {
	shirt, err := GetClone(WhiteShirtType)

	if err != nil {
		t.Errorf("should return a shirt instance")
	}

	shirt2, err := GetClone(WhiteShirtType)

	if err != nil {
		t.Errorf("should return a shirt instance")
	}

	if &shirt == &shirt2 {
		t.Errorf("shirt objects should be different")
	}

}

func TestCloneBlackTshirt(t *testing.T) {
	shirt, err := GetClone(BlackShirtType)

	if err != nil {
		t.Errorf("should return a shirt instance")
	}

	shirt2, err := GetClone(BlackShirtType)
	if err != nil {
		t.Errorf("should return a shirt instance")
	}

	if &shirt == &shirt2 {
		t.Errorf("shirt objects should be different")
	}

}
