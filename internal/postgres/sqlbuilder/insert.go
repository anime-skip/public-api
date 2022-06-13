package sqlbuilder

import (
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal/log"
)

type insertBuilder struct {
	table   string
	columns []string
	values  []string
	args    []any
}

func Insert(table string, record map[string]any) *insertBuilder {
	columns := []string{}
	values := []string{}
	args := []any{}
	i := 1
	for column, value := range record {
		columns = append(columns, column)
		args = append(args, value)
		values = append(values, fmt.Sprintf("$%d", i))
		i++
	}
	return &insertBuilder{
		table:   table,
		columns: columns,
		args:    args,
		values:  values,
	}
}

func (b *insertBuilder) ToSQL() (sql string, args []any) {
	args = b.args
	columns := strings.Join(b.columns, ", ")
	values := strings.Join(b.values, ", ")
	sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, b.table, columns, values)
	log.V("SQL Insert: %s %v", sql, args)
	return sql, args
}
