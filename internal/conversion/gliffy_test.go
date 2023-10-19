package conversion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPointsOffset(t *testing.T) {
	points := [][]float64{
		{0, 0},
		{32, 2},
		{34, -4},
		{-14, -8},
		{-20, -10},
		{-35, -13},
		{-49, -16},
		{-73, -16},
		{-109, -12},
		{-120, -7},
		{-123, -5},
	}

	expected := [][]float64{
		{123, 16}, {155, 18}, {157, 12}, {109, 8}, {103, 6}, {88, 3}, {74, 0}, {50, 0}, {14, 4}, {3, 9}, {0, 11},
	}

	result := AddPointsOffset(points)
	fmt.Println(result)

	svg := ConvertPointsToSvgPath(result, 297, 415)
	fmt.Println(svg)

	assert.Equal(t, expected, result)
}
