package standard

import (
	"auth/service/database"
)

var (
	db             = database.DB
	lockoutManager = NewLockoutManager()
)
