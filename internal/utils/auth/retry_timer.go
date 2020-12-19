package auth

import (
	"time"

	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/math"
)

type loginAttemptInfo struct {
	lastLogInAt    time.Time
	failedAttempts int
}

type loginRetryTimer struct {
	cache map[string]loginAttemptInfo
}

var LoginRetryTimer = loginRetryTimer{
	cache: map[string]loginAttemptInfo{},
}

func (timer loginRetryTimer) Failure(usernameEmail string) {
	now := time.Now()
	var newFailedAttempts int
	if value, ok := timer.cache[usernameEmail]; ok {
		if now.After(value.lastLogInAt.Add(3 * time.Hour)) {
			newFailedAttempts = 1
		} else {
			newFailedAttempts = value.failedAttempts + 1
		}
	} else {
		newFailedAttempts = 1
	}
	timer.cache[usernameEmail] = loginAttemptInfo{
		lastLogInAt:    now,
		failedAttempts: newFailedAttempts,
	}

	msSleep := math.BoundedInt(
		constants.LOGIN_RETRY_INCREMENT*(newFailedAttempts-constants.LOGIN_RETRY_FREEBEES),
		0, constants.LOGIN_RETRY_MAX_SLEEP,
	)
	if msSleep > 0 {
		time.Sleep(time.Duration(msSleep) * time.Millisecond)
	}
}

func (timer loginRetryTimer) Success(usernameEmail string) {
	delete(timer.cache, usernameEmail)
}
