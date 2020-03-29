package database

import (
	"fmt"
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func updateColumn(columnName string) func(scope *gorm.Scope) {
	return func(scope *gorm.Scope) {

		if !scope.HasError() {
			updateColumnField, hasUpdateColumnField := scope.FieldByName(columnName)
			if hasUpdateColumnField {
				userID, ok := scope.DB().Get(constants.CTX_USER_ID)
				log.V("Update %s to %v", columnName, userID)
				if !ok {
					log.V("CTX_USER_ID does not exist on database values, skipping update")
					return
				}
				scope.SetColumn(updateColumnField.DBName, userID)
			}
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")
		deletedByField, hasDeletedByField := scope.FieldByName("DeletedByUserID")

		userID, hasUserID := scope.DB().Get(constants.CTX_USER_ID)
		if !hasUserID {
			log.V("CTX_USER_ID does not exist on database values, skipping update")
		}

		if !scope.Search.Unscoped && hasDeletedAtField && hasDeletedByField && hasUserID {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now()),
				scope.Quote(deletedByField.DBName),
				scope.AddToVars(userID),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
