package standard

import (
	"auth/service/database/types"
	"fmt"
	"time"
)

var oldTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

func IsLocked(email string) bool {
	security := types.SecurityConfig{}
	if err := db.First(&security).Error; err != nil {
		fmt.Println(err)
		return false
	}
	lockout := types.Lockout{Email: email}
	if err := db.Where(&lockout).First(&lockout).Error; err != nil {
		fmt.Println(err)
		return false
	}
	if time.Since(lockout.LastLockout).Seconds() < float64(security.LockoutDuration) {
		return false
	}
	return true
}

func Succeed(email string) error {
	return db.Where("email = ?", email).Delete(&types.Lockout{}).Error
}

func Fail(email string) error {
	security := types.SecurityConfig{}
	if err := db.First(&security).Error; err != nil {
		return err
	}
	lockout := types.Lockout{
		Email:       email,
		Attempts:    0,
		LastLockout: oldTime,
	}
	if err := db.Where("email = ?", email).FirstOrCreate(&lockout).Error; err != nil {
		return err
	}

	lockout.Attempts++
	if lockout.Attempts >= security.LockoutThreshold {
		lockout.LastLockout = time.Now()
	}

	if err := db.Save(&lockout).Error; err != nil {
		return err
	}

	return nil
}
