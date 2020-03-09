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

import "errors"

// RegionPosition represents the map coordinates of a region
type RegionPosition struct {
	x uint32
	y uint32
}

// NewRegionPosition returns the region's map coordinates, given an x and y
func NewRegionPosition(x uint32, y uint32) (*RegionPosition, error) {

	maxRegionXY := uint32(9999)

	if x > maxRegionXY {
		err := errors.New("invalid arguments")
		return nil, err
	}

	if y > maxRegionXY {
		err := errors.New("invalid arguments")
		return nil, err
	}

	r := new(RegionPosition)

	r.x = x
	r.y = y

	return r, nil
}
