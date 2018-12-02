package newsletter

import (
	"os"
)

func main() {
	languages := os.Args[1:]
	newsletter(languages)
}
