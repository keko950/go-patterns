package creational

import (
	"fmt"
	"testing"
)

func TestGetSuzukiFactoryType(t *testing.T) {
	builder, err := BuildFactory(SuzukiFactoryType)

	if err != nil {
		t.Errorf("factory should create a SuzukiFactory object")
	}

	suzuki := builder.Build()

	if suzuki.GetSeats() != 2 {
		t.Errorf("a motorbike should have 2 seats")
	}

	if suzuki.GetWheels() != 2 {
		t.Errorf("a motorbike should have 2 seats")
	}

	motorbike := suzuki.(Motorbike)

	if motorbike.GetBrand() != SUZUKI {
		t.Errorf("motorbike object should be a suzuki")
	}

	a := 127 % 127
	fmt.Print(a)

}

func TestGetMercedesFactoryType(t *testing.T) {
	builder, err := BuildFactory(MercedesFactoryType)

	if err != nil {
		t.Errorf("factory should create a MercedesFactory object")
	}

	mercedes := builder.Build()

	if mercedes.GetSeats() != 5 {
		t.Errorf("a car should have 5 seats")
	}

	if mercedes.GetWheels() != 4 {
		t.Errorf("a car should have 4 seats")
	}

	car := mercedes.(Car)

	if car.GetDoors() != 5 {
		t.Errorf("a car should have 5 doors")
	}

	if car.GetBrand() != MERCEDES {
		t.Errorf("car object should be a mercedes")
	}

}
