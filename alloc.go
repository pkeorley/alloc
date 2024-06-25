package alloc

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
)

type Supplier interface {
	Alloc(uint64)
}

type Allocator struct {
	to *[]uint64
}

func NewAllocator() *Allocator {
	return &Allocator{to: new([]uint64)}
}

func (a Allocator) Alloc(count uint64) {
	for i := uint64(0); i < count/32; i++ {
		*a.to = append(*a.to, 1)
	}
}
