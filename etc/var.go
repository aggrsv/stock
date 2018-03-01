package etc

import (
	"os"
)

var (
	Service string
)

func ServiceName() string {
	if Service != "" {
		return Service
	}
	return os.Args[0]
}
