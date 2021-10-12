package main

import (
	"fmt"
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

func trace(s string) (string, time.Time) {
	fmt.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func (c1 Cordinate) Trilaterate(c2, c3 Cordinate) (x, y float64) {

	return Trilaterate(c1, c2, c3)
}

func (c *Cordinate) solveY(x float64) (y1, y2 float64) {
	tmp := math.Sqrt(math.Pow(c.d, 2) - math.Pow(x-c.x, 2))
	y1 = c.y + tmp
	y2 = c.y - tmp

	return y1, y2
}

// SolveY array 296 &{100 100 50}
func (c *Cordinate) solveYArray(x float64) (y [2]float64, skipped bool) {

	tmp := math.Sqrt(math.Pow(c.d, 2) - math.Pow(x-c.x, 2))

	if math.IsNaN(tmp) {
		return y, true
	}

	y[0] = c.y + tmp
	y[1] = c.y - tmp

	return y, false
}

func Trilaterate(c1, c2, c3 Cordinate) (x, y float64) {
	defer un(trace("Trilaterate 1"))
	// https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method
	// Possiblility to divide by zero. Rearranging c1,c2,c3 may result in correct answer when that happens

	d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.d, &c2.d, &c3.d, &c1.x, &c2.x, &c3.x, &c1.y, &c2.y, &c3.y

	x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2)))

	y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / (2**j2 - 2**j1)

	return x, y
}

func (c1 Cordinate) TrilaterateRast(c2, c3 Cordinate) (x, y float64) {

	return TrilaterateRast(c1, c2, c3)
}

func TrilaterateRast(c1, c2, c3 Cordinate) (x, y float64) {

	defer un(trace("Trilaterate 2"))
	width, height := calcBounds(c1, c2, c3)

	// Create a matrix to roughly calculate all points of each circle equation on
	matrix := make([][]uint64, width)
	for i := 0; i < width; i++ {
		matrix[i] = make([]uint64, height)
	}

	done := make(chan string)

	fillMatrix := func(c Cordinate) {
		for x := 0; x < width; x++ {
			yValues, skipped := c.solveYArray(float64(x))
			if !skipped {
				for _, y := range yValues {

					atomic.AddUint64(&matrix[x][int(y+0.5)], 1)

				}
			}
			runtime.Gosched()
		}
		done <- "done"
	}

	go fillMatrix(c1)
	go fillMatrix(c2)
	go fillMatrix(c3)

	for j := 0; j < 3; j++ {
		select {
		case <-done:
		}
	}

	for l := range matrix {
		for k := range matrix[l] {
			if matrix[l][k] > 0 {
				if matrix[l][k] == 3 {
					x, y = float64(l), float64(k)
					return x, y
				}
			}
		}
	}

	x = 120
	y = 140
	return x, y
}

func calcBounds(c1, c2, c3 Cordinate) (width, height int) {
	// Stub for now always return a certain bounding
	width = 300
	height = 300
	return
}
