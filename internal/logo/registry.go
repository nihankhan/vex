package logo

import "strings"

var providers = []Provider{
	&Ubuntu{},
}

func Get(name string) Provider {
	name = strings.ToLower(name)

	for _, p := range providers {
		if p.Name() == name {
			return p
		}
	}

	return &Ubuntu{}
}
