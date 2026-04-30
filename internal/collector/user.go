package collector

import (
	"fmt"
	"os"
)

type IdentityCollector struct{}

func (i *IdentityCollector) Name() string { return "Identity" }

func (i *IdentityCollector) Collect() (*Info, error) {
	host, _ := os.Hostname()
	user := os.Getenv("USER")

	return &Info{
		Label: "Identity",
		Value: fmt.Sprintf("%s@%s", user, host),
	}, nil
}
