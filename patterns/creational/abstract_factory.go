package creational

import (
	"fmt"
)

type CarBrand int
type MotorbikeBrand int
type FactoryType int

const (
	SUZUKI MotorbikeBrand = iota
	KAWASAKI
)

const (
	BMW CarBrand = iota + 2
	MERCEDES
)

const (
	SuzukiFactoryType FactoryType = iota + 4
	MercedesFactoryType
)

type Vehicle interface {
	GetWheels() int
	GetSeats() int
}

type Car interface {
	GetDoors() int
	GetBrand() CarBrand
}
type Motorbike interface {
	GetBrand() MotorbikeBrand
}

type VehicleFactory interface {
	Build() Vehicle
}

type SuzukiFactory struct {
}
type MercedesFactory struct{}

type Suzuki struct {
}

func (m *Suzuki) GetWheels() int {
	return 2
}

func (m *Suzuki) GetSeats() int {
	return 2
}

func (m *Suzuki) GetBrand() MotorbikeBrand {
	return SUZUKI
}

type Mercedes struct{}

func (m *Mercedes) GetWheels() int {
	return 4
}

func (m *Mercedes) GetSeats() int {
	return 5
}

func (m *Mercedes) GetDoors() int {
	return 5
}

func (m *Mercedes) GetBrand() CarBrand {
	return MERCEDES
}

func (m *SuzukiFactory) Build() Vehicle {
	return new(Suzuki)
}

func (m *MercedesFactory) Build() Vehicle {
	return new(Mercedes)
}

func BuildFactory(factoryType FactoryType) (VehicleFactory, error) {
	switch factoryType {
	case SuzukiFactoryType:
		return new(SuzukiFactory), nil
	case MercedesFactoryType:
		return new(MercedesFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", factoryType)
	}
}
