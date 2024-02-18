package standard

//
//import (
//	"auth/service/database/types"
//	"encoding/base64"
//	"fmt"
//	"github.com/labstack/gommon/log"
//	"os"
//	"time"
//)
//
//var initTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
//
//// LockoutManager manages the lockout guards of all apps in the service.
//type LockoutManager struct {
//	AppGuards map[string]*AppLockout
//}
//
//func NewLockoutManager() *LockoutManager {
//	var apps []types.App
//	if err := db.Find(&apps).Error; err != nil {
//		log.Errorf("Failed to load apps - lockout guard, %v", err)
//		os.Exit(1)
//	}
//
//	guard := &LockoutManager{
//		AppGuards: make(map[string]*AppLockout),
//	}
//
//	for _, app := range apps {
//		guard.AppGuards[app.ID] = &AppLockout{
//			Status:    make(map[string]*Lockout),
//			Threshold: app.LockoutThreshold,
//			Duration:  time.Duration(app.LockoutDuration) * time.Second,
//		}
//	}
//
//	go func() {
//		for {
//			time.Sleep(time.Hour)
//			for _, appGuard := range guard.AppGuards {
//				appGuard.Status = make(map[string]*Lockout)
//			}
//		}
//	}()
//
//	return guard
//}
//
//// AppLockout manages the lockout status of email addresses in a singular app.
//type AppLockout struct {
//	Status    map[string]*Lockout
//	Threshold int
//	Duration  time.Duration
//}
//
//type Lockout struct {
//	Time     time.Time `json:"time"`
//	Attempts int       `json:"attempts"`
//}
//
//func (g *AppLockout) IsLocked(email string) bool {
//	key := g.getKey(email)
//	source, exists := g.Status[key]
//	if !exists {
//		return false
//	}
//	if time.Now().Sub(source.Time) > g.Duration {
//		source.Attempts = 0
//		return false
//	}
//	return true
//}
//
//func (g *AppLockout) Fail(email string) {
//	key := g.getKey(email)
//	source, exists := g.Status[key]
//	if !exists {
//		g.create(key)
//		source = g.Status[key]
//	}
//	source.Attempts += 1
//	if source.Attempts == g.Threshold {
//		source.Time = time.Now()
//	}
//}
//
//func (g *AppLockout) create(key string) {
//	g.Status[key] = &Lockout{
//		Attempts: 0,
//		Time:     initTime,
//	}
//}
//
//func (g *AppLockout) getKey(email string) string {
//	key := fmt.Sprint(email)
//	return base64.StdEncoding.EncodeToString([]byte(key))
//}
//
//// implement a clean function.
//
//func (g *AppLockout) clean() {
//	g.Status = make(map[string]*Lockout)
//}
