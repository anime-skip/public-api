package utils

import (
	"math/rand"
	"regexp"
	"time"

	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func StartTransaction(db *gorm.DB, inTransaction bool) *gorm.DB {
	if inTransaction {
		return db
	} else {
		tx := db.Begin()
		return tx
	}
}

func StartTransaction2(db *gorm.DB, err *error) (tx *gorm.DB, commitOrRollback func()) {
	tx = db.Begin()
	var txID int
	if env.IS_DEV {
		txID = rand.New(rand.NewSource(time.Now().Unix())).Int()
		log.V("Begin transaction %d", txID)
	}
	commitOrRollback = func() {
		if r := recover(); r != nil || *err != nil {
			tx.Rollback()
			if env.IS_DEV {
				log.V("Rollback %d", txID)
			}
			if r != nil {
				log.E("Rollback, panicked: %v", r)
			} else {
				log.V("Rollback, expected error: %v", err)
			}
		} else {
			tx.Commit()
			if env.IS_DEV {
				log.V("Commit %d", txID)
			}
		}
	}
	return tx, commitOrRollback
}

func CommitTransaction(tx *gorm.DB, wasInTransaction bool) *gorm.DB {
	if wasInTransaction {
		return tx
	} else {
		return tx.Commit()
	}
}

func RandomProfileURL() string {
	return "https://avatars3.githubusercontent.com/u/1400247?s=460&v=4"
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
