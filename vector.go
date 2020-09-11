package main

import "math"

type Vector struct {
	x float64
	y float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
	}
}

func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
	}
}

func (v1 Vector) AddByValue(value float64) Vector {
	return Vector{
		x: v1.x + value,
		y: v1.y + value,
	}
}

func (v1 Vector) MultiplyByValue(value float64) Vector {
	return Vector{
		x: v1.x * value,
		y: v1.y * value,
	}
}

func (v1 Vector) DivideByValue(value float64) Vector {
	return Vector{
		x: v1.x / value,
		y: v1.y / value,
	}
}

func (v1 Vector) Limit(lowerValue, upperValue float64) Vector {
	return Vector{
		x: math.Min(math.Max(v1.x, lowerValue), upperValue),
		y: math.Min(math.Max(v1.y, lowerValue), upperValue),
	}
}

func (v1 Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
