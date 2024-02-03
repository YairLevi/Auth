package standard

import (
	"encoding/base64"
	"fmt"
	"time"
)

var initTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

type Lockout struct {
	Time     time.Time `json:"time"`
	Attempts int       `json:"attempts"`
}

type LockoutGuard struct {
	Lockouts  map[string]*Lockout
	Threshold int
	Duration  time.Duration
}

func (g *LockoutGuard) IsLocked(appID, email string) bool {
	key := g.getKey(appID, email)
	source, exists := g.Lockouts[key]
	if !exists {
		return false
	}
	if time.Now().Sub(source.Time) > g.Duration {
		source.Attempts = 0
		return false
	}
	return true
}

func (g *LockoutGuard) Fail(appID, email string) {
	key := g.getKey(appID, email)
	if _, exists := g.Lockouts[key]; !exists {
		g.create(key)
	}
	source := g.Lockouts[key]
	source.Attempts += 1
	if source.Attempts == g.Threshold {
		source.Time = time.Now()
	}
}

func (g *LockoutGuard) create(key string) {
	g.Lockouts[key] = &Lockout{
		Attempts: 0,
		Time:     initTime,
	}
}

func (g *LockoutGuard) getKey(appID, email string) string {
	key := fmt.Sprint(appID, email)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// implement a clean function.

//func (g *LockoutGuard) clean() {
//	for key, source := range g.Lockouts {
//		if
//		delete(g.Lockouts, key)
//	}
//}

/*

AppID and Email Address together make a key.
Keep counter for each key if the counter reaches the threshold, lockout.
the lockout will contain last lockout date.
on login, if the current date - lockout date is larger than duration, allow. otherwise, deny.
struct:

key string: {
	LastLockout time
	Attempts 	int
}
*/
