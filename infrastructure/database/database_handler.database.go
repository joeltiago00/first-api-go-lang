package database

import "github.com/joeltiago00/first-api-go-lang/config"

func Handler() {
	DatabaseFactoryHandle(config.DbConfig().Connection)
}
