package utils

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
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
		log.E("Failed getting userID from context")
	}
	return "", fmt.Errorf("500 Internal Error [003]")
}

func StartTransaction(db *gorm.DB, inTransaction bool) *gorm.DB {
	if inTransaction {
		return db
	} else {
		tx := db.Begin()
		// userID, hasUserID := db.Get(constants.CTX_USER_ID)
		// if !hasUserID {
		// 	log.W("Failed to pass CTX_USER_ID into transaction")
		// 	return tx
		// }
		// tx.Set(constants.CTX_USER_ID, userID)
		return tx
	}
}
