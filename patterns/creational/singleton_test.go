package creational

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	singletonCounter := GetInstance()
	singletonCounter.AddOne()

	singletonCounter2 := GetInstance()
	singletonCounter2.AddOne()
	singletonCounter2.AddOne()

	if singletonCounter2.GetCount() != singletonCounter.GetCount() {
		t.Errorf("counter should have exactly the same count")
	}

	if singletonCounter2 != singletonCounter {
		t.Errorf("instance object should be the same")
	}
}
