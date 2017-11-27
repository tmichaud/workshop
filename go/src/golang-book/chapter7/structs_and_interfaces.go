package main

import ("fmt"; "math")



func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1,y1,x1,y2)
	w := distance(x1,y1,x2,y1)
	return l*w
}

func circleArea(x,y,r float64) float64 {
	return math.Pi * r * r
}

type Circle struct {
	x,y,r float64
}

type Rectangle struct {
	x1,y1,x2,y2 float64;
}

func circleAreaC(c Circle ) float64 {
	return math.Pi * c.r * c.r
}

func circleAreapC(c *Circle ) float64 {
	return math.Pi * c.r * c.r
}

//func (c *Circle) area() float64 {
//	return math.Pi * c.r * c.r
//}

//func (r *Rectangle) area() float64 {
//	l := distance(r.x1, r.y1, r.x1, r.y2)
//	w := distance(r.x1, r.y1, r.x2, r.y1)
//	return l*w
//}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l*w
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l + l + w + w
}

//func totalArea(circles ...Circle) float64 {
//	var total float64
//	for _, c := range circles {
//		total += c.area()
//	}
//	return total
//}

// Invalid - can't have two variadic parameters
//func totalArea(circles ...Circle, rectangles ...Rectangle) float64 {
//	var total float64
//	for _, c := range circles {
//		total += c.area()
//	}
//	for _, r := range rectangles {
//		total += r.area()
//	}
//	return total
//}

// Fixed -- but we have to alter it EVERY time we add a new shape...sucky.
func totalAreaBad(circles []Circle, rectangles []Rectangle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	for _, r := range rectangles {
		total += r.area()
	}
	return total
}

type Shape interface {
	area() float64
	perimeter() float64
}

// Now using Shapes
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

type MultiShape struct {
	shapes []Shape
}

func (m MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}




func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Structs ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	var rx1, ry1 float64 = 0,0
	var rx2, ry2 float64 = 10,10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println("Area of Rectangle: :", rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println("Area of circle:", circleArea(cx, cy, cr))

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Using Struct ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	c := Circle{0, 0, 5}
	fmt.Println("Area of circle:", circleAreaC(c))
	fmt.Println("Area of circle:", circleAreapC(&c))

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Adding method ...  with a receiver  ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println("Area of circle:", c.area())
	r := Rectangle{ 0, 0, 10, 10}
	fmt.Println("Area of rectangle: ", r.area())

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" To is-a  ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	//  weird.

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Interfaces  ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println("Total area is ", totalArea(&c, &r))

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Interfaces and fields  ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	multiShape := MultiShape {
		shapes: []Shape {
			Circle{0, 0, 5},
			Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println("Total area is ", totalArea(multiShape))
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Exercises ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	fmt.Println("Total area is ", totalPerimeter(multiShape))
}

