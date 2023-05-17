package db

import (
	"github.com/const-tmp/rotation"
)

type DB struct {
	Host, Port, User, Password rotation.ValueGetter
}
