package database

import (
	"github.com/aklinker1/anime-skip-backend/pkg/utils/constants"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
	"github.com/jinzhu/gorm"
)

func updateColumn(columnName string) func(scope *gorm.Scope) {
	return func(scope *gorm.Scope) {
		if !scope.HasError() {
			if scope.HasColumn(columnName) {
				if userId, ok := scope.DB().Get(constants.USER_ID_FROM_TOKEN); ok {
					scope.SetColumn(columnName, userId)
				} else {
					log.E("Could not update '%s' because userId is not present on the scope's database values")
				}
			}
		}
	}
}
