package sqlbuilder

import (
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal/log"
)

type updateBuilder struct {
	table   string
	columns []string
	args    []any
}

func Update(table string, id any, record map[string]any) *updateBuilder {
	columns := []string{}
	args := []any{}
	for column, value := range record {
		columns = append(columns, column)
		args = append(args, value)
	}
	args = append(args, id)
	return &updateBuilder{
		table:   table,
		columns: columns,
		args:    args,
	}
}

func (b *updateBuilder) ToSQL() (sql string, args []any) {
	args = b.args
	setters := []string{}
	for i, column := range b.columns {
		setters = append(setters, fmt.Sprintf("%s = $%d", column, i+1))
	}
	sql = fmt.Sprintf(
		`UPDATE %s SET %s WHERE id = $%d`,
		b.table,
		strings.Join(setters, ", "),
		len(setters)+1,
	)
	log.V("SQL Update: %s %v", sql, args)
	return sql, args
}
