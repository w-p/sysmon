package pkg

// Bounds holds a min and max value
type Bounds struct {
	min float64
	max float64
}

// ScaleFn defines the scaling function signature
type ScaleFn func(n float64) float64

// NewScale creates a new scaling function based on a domain and range
func NewScale(d Bounds, r Bounds) ScaleFn {
	return func(n float64) float64 {
		// https://stats.stackexchange.com/a/281165
		return (r.max-r.min)*((n-d.min)/(d.max-d.min)) + r.min
	}
}
