package xorshift

const (
	defaultX = 123456789
	defaultY = 362436069
	defaultZ = 521288629
	defaultW = 88675123
)

// Rand is a source of random numbers
type Rand struct {
	x, y, z, w uint64
}

// New returns a new Rand object
func New(seed uint64) *Rand {
	return &Rand{}
}

func (r *Rand) Int() uint64 {
	t := r.x ^ (r.x << 11)
	r.x = r.y
	r.y = r.z
	r.z = r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))
	return r.w
}
