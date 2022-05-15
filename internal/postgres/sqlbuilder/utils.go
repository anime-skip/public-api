package sqlbuilder

import (
	"fmt"
	"strings"
)

func replaceWithOrderedParams(sql string) string {
	i := 1
	for strings.ContainsRune(sql, '?') {
		sql = strings.Replace(sql, "?", fmt.Sprintf("$%d", i), 1)
		i++
	}
	return sql
}
