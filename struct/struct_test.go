package main

import (
	"fmt"
	"math"
	"testing"
)

type Rectangle struct{
	width float64
	length float64
}

type Circle struct{
	radius float64
}

type Shape interface {
	Area() float64
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
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10},314.1592653589793},
	}

	for _, ta := range areaTests{
		got := ta.shape.Area()
		if got != ta.want {
			t.Errorf("got %f want %f",got,ta.want)
		}
	}
}

// 	checkArea := func(t testing.TB,shape Shape, want float64){
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want{
// 			t.Errorf("got %.2f want %.2f",got,want)
// 		}
// 	}
	
// 	t.Run("rectangle",func(t *testing.T) {
// 	rectangle := Rectangle{100,200}
// 	checkArea(t,rectangle,20000)
// 	})

// 	t.Run("circles",func(t *testing.T){
// 		circle := Circle{10}
// 		checkArea(t,circle,314.1592653589793)	
// 	})

	
// }

func (rectangle Rectangle) Area() float64 {	
	area := rectangle.length*rectangle.width
	fmt.Println(area)
	return area
}

func (circle Circle) Area() float64{
	area := math.Pi*circle.radius*circle.radius
	return area
}

func Perimeter(rectangle Rectangle)float64 {
	perimeter := 2*(rectangle.length+rectangle.width)
	fmt.Println(perimeter)
	return perimeter
}



