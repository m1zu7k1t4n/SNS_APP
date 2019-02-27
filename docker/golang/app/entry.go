package main

import (
	"github.com/arasan01/app/api"
)

func init() {
	api.Router_init()
}

func main() {
	api.Router()
}
