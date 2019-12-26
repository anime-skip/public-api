package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
)

func findOne(orm *database.ORM, modelName string, output interface{}, whereArg string, whereValues ...interface{}) error {
	db := orm.DB.New()
	var count int64
	db.Where(whereArg, whereValues...).Find(output).Count(&count)
	if count == 0 {
		return fmt.Errorf("Failed to find %s where %s %+v", modelName, whereArg, whereValues)
	}
	return nil
}
