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

	c := new(RegionPosition)

	c.x = x
	c.y = y

	return c, nil
}
