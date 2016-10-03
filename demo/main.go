package main

import (
	"log"

	_ "github.com/itpkg/magnolia/engines/auth"
	_ "github.com/itpkg/magnolia/engines/forum"
	_ "github.com/itpkg/magnolia/engines/mail"
	_ "github.com/itpkg/magnolia/engines/reading"
	_ "github.com/itpkg/magnolia/engines/shop"
	_ "github.com/itpkg/magnolia/engines/vpn"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"github.com/itpkg/magnolia/web"
)

var version string

func main() {
	if err := web.Main(version); err != nil {
		log.Fatal(err)
	}
}
