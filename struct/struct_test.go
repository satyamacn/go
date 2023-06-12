package main

import "testing"
import "fmt"

type Rectangle struct{
	width float64
	length float64
}

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{100.0,200.0}
	got := Perimeter(rectangle)
	want := 600.0

	if got != want {
		t.Errorf("Got %.2f want %.2f",got,want)
	}
}

func TestArea(t *testing.T){
	area := Rectangle{100,200}
	got := Area(area)
	want := 20000.0

	if got != want {
		t.Errorf("Got %.2f want %.2f",got,want)
	}
}

func Area(rectangle Rectangle)float64 {	
	area := rectangle.length*rectangle.width
	fmt.Println(area)
	return area
}

func Perimeter(rectangle Rectangle)float64 {
	perimeter := 2*(rectangle.length+rectangle.width)
	fmt.Println(perimeter)
	return perimeter
}

