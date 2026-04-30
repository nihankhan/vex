package logo

type Provider interface {
	Name() string
	Lines() []string
}
