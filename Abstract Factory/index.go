package main

import "fmt"

// Abstract Factory interface
type ShapeFactory interface {
    CreateCircle() Circle
    CreateSquare() Square
}

// Concrete Factory 1
type RedShapeFactory struct{}

func (r RedShapeFactory) CreateCircle() Circle {
    return Circle{color: "red"}
}

func (r RedShapeFactory) CreateSquare() Square {
    return Square{color: "red"}
}

// Concrete Factory 2
type BlueShapeFactory struct{}

func (b BlueShapeFactory) CreateCircle() Circle {
    return Circle{color: "blue"}
}

func (b BlueShapeFactory) CreateSquare() Square {
    return Square{color: "blue"}
}

// Abstract Product 1
type Circle struct {
    color string
}

func (c Circle) Draw() {
    fmt.Printf("Drawing %s circle\n", c.color)
}

// Abstract Product 2
type Square struct {
    color string
}

func (s Square) Draw() {
    fmt.Printf("Drawing %s square\n", s.color)
}

func main() {
    // Create a RedShapeFactory
    redFactory := RedShapeFactory{}

    // Use the RedShapeFactory to create a red circle and square
    redCircle := redFactory.CreateCircle()
    redSquare := redFactory.CreateSquare()

    // Draw the red circle and square
    redCircle.Draw()
    redSquare.Draw()

    // Create a BlueShapeFactory
    blueFactory := BlueShapeFactory{}

    // Use the BlueShapeFactory to create a blue circle and square
    blueCircle := blueFactory.CreateCircle()
    blueSquare := blueFactory.CreateSquare()

    // Draw the blue circle and square
    blueCircle.Draw()
    blueSquare.Draw()
}