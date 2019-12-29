package database

import (
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func updateColumn(columnName string) func(scope *gorm.Scope) {
	return func(scope *gorm.Scope) {
		if !scope.HasError() {
			if scope.HasColumn(columnName) {
				if userId, ok := scope.DB().Get(constants.CTX_USER_ID); ok {
					scope.SetColumn(columnName, userId)
				} else {
					log.E("Could not update '%s' because userId is not present on the scope's database values")
				}
			}
		}
	}
}
