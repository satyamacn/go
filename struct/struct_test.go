package main

import "testing"
import "fmt"

func TestPerimeter(t *testing.T){
	got := Perimeter(100.0,200.0)
	want := 600.0

	if got != want {
		t.Errorf("Got %.2f want %.2f",got,want)
	}
}

func TestArea(t *testing.T){
	got := Area(20.0,40.0)
	want := 800.0

	if got != want {
		t.Errorf("Got %.2f want %.2f",got,want)
	}
}

func Area(length, width float64)float64 {
	
	area := length*width
	fmt.Println(area)
	return area
}

func Perimeter(length, width float64)float64 {
	perimeter := 2*(length+width)
	fmt.Println(perimeter)
	return perimeter
}

