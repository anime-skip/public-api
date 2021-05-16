package utils

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(constants.CTX_GIN_CONTEXT)
	if ginContext == nil {
		log.E("ctx[\"CTX_GIN_CONTEXT\"] is missing")
		return nil, fmt.Errorf("500 Internal Error [001]")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		log.E("ctx[\"CTX_GIN_CONTEXT\"] is not *gin.Context")
		return nil, fmt.Errorf("500 Internal Error [002]")
	}
	return gc, nil
}

func UserIDFromContext(ctx context.Context) (string, error) {
	context, err := GinContext(ctx)
	if err == nil {
		if userID, hasUserID := context.Get(constants.CTX_USER_ID); hasUserID {
			return userID.(string), nil
		}
	}
	return "", fmt.Errorf("500 Internal Error [003]")
}

func StartTransaction2(db *gorm.DB, err *error) (tx *gorm.DB, commitOrRollback func() interface{}) {
	tx = db.Begin()
	var txID int
	if env.IS_DEV {
		txID = rand.New(rand.NewSource(time.Now().Unix())).Int()
		log.V("Begin transaction %d", txID)
	}
	commitOrRollback = func() interface{} {
		if r := recover(); r != nil {
			tx.Rollback()
			log.V("Rollback %d", txID)
			log.E("Rollback due to panicked: %v", r)
			return r
		}
		if *err != nil {
			tx.Rollback()
			log.V("Rollback %d", txID)
			log.V("Rollback due to known error: %v", err)
			return err
		}

		tx.Commit()
		if env.IS_DEV {
			log.V("Commit %d", txID)
		}
		return nil
	}
	return tx, commitOrRollback
}

func RandomProfileURL() string {
	return "https://avatars3.githubusercontent.com/u/1400247?s=460&v=4"
}

func GetIP(ctx context.Context) (string, error) {
	ginCtx, err := GinContext(ctx)
	if err != nil {
		return "", err
	}
	forwarded := ginCtx.Request.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded, nil
	}
	return ginCtx.Request.RemoteAddr, nil
}

func StringArrayIncludes(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func CleanURL(url string) string {
	queryRegex := regexp.MustCompile(`\?.*$`)
	withoutQuery := queryRegex.ReplaceAllString(url, "")
	slashRegex := regexp.MustCompile(`\/$`)
	return slashRegex.ReplaceAllString(withoutQuery, "")
}

func ArrayOrNil(array []string) []string {
	if len(array) == 0 {
		return nil
	}
	return array
}
