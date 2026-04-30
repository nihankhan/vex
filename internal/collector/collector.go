package collector

type Info struct {
	Label string
	Value string
}

// var _ Collector = (*Info)(nil)

type Collector interface {
	Name() string
	Collect() (*Info, error)
}
