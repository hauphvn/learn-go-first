package main

import (
	"server/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8009") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
