/*
MIT License

Copyright (c) 2020 ScriptWorks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package position

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const localItems int = 3

const (
	iLocalX int = 0
	iLocalY int = 1
	iLocalZ int = 2
)

// LocalPosition represents local coordinates within a region
type LocalPosition struct {
	x float32
	y float32
	z float32
}

// NewLocalPosition returns local 3D coordinates, given an x, y and z
func NewLocalPosition(
	x float32, y float32, z float32) (
	*LocalPosition, error) {

	// Compare the floats as uint32s
	maxRegionXY := math.Float32bits(255.0)
	maxRegionZ := math.Float32bits(4096.0)

	if math.Float32bits(x) > maxRegionXY {
		err := errors.New("invalid argument")
		return nil, err
	}

	if math.Float32bits(y) > maxRegionXY {
		err := errors.New("invalid argument")
		return nil, err
	}

	if math.Float32bits(z) > maxRegionZ {
		err := errors.New("invalid argument")
		return nil, err
	}

	l := new(LocalPosition)

	l.x = x
	l.y = y
	l.z = z

	return l, nil
}

// NewLocalPositionFromVector returns local 3D coordinates, given a vector passed as a string
func NewLocalPositionFromVector(vector string) (*LocalPosition, error) {
	if vector == "" {
		err := errors.New("invalid argument")
		return nil, err
	}

	trimmed := strings.TrimSpace(vector)

	if !strings.HasPrefix(trimmed, "<") || !strings.HasSuffix(trimmed, ">") {
		err := errors.New("invalid argument")
		return nil, err
	}

	items := strings.Split(strings.Trim(trimmed, "<>"), ",")
	if len(items) != localItems {
		err := errors.New("invalid argument")
		return nil, err
	}

	posX, err := strconv.ParseFloat(strings.TrimSpace(items[iLocalX]), 32)
	if err != nil {
		return nil, err
	}

	posY, err := strconv.ParseFloat(strings.TrimSpace(items[iLocalY]), 32)
	if err != nil {
		return nil, err
	}

	posZ, err := strconv.ParseFloat(strings.TrimSpace(items[iLocalZ]), 32)
	if err != nil {
		return nil, err
	}

	l, err := NewLocalPosition(float32(posX), float32(posY), float32(posZ))
	if err != nil {
		return nil, err
	}

	return l, nil
}
