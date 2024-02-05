package standard

import "auth/database"

var (
	db             = database.DB
	lockoutManager = NewLockoutManager()
)
