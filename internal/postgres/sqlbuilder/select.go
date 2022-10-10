package sqlbuilder

import (
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"github.com/samber/lo"
)

type OrderBy struct {
	column    string
	direction string
}

type selectBuilder struct {
	table             string
	columns           []string
	scanDest          []any
	args              []any
	where             []string
	limit             string
	offset            string
	order             *OrderBy
	includeSoftDelete bool
	pagination        *internal.Pagination
}

func Select(table string, columns map[string]any) *selectBuilder {
	selectColumns := []string{}
	scanDest := []any{}
	for name, target := range columns {
		selectColumns = append(selectColumns, name)
		scanDest = append(scanDest, target)
	}
	return &selectBuilder{
		table:    table,
		columns:  selectColumns,
		scanDest: scanDest,
	}
}

func (b *selectBuilder) Where(condition string, args ...any) *selectBuilder {
	b.where = append(b.where, condition)
	b.args = append(b.args, args...)
	return b
}

func (b *selectBuilder) IncludeSoftDeleted() *selectBuilder {
	b.includeSoftDelete = true
	return b
}

func (b *selectBuilder) OrderBy(column string, direction string) *selectBuilder {
	b.order = &OrderBy{
		column:    column,
		direction: direction,
	}
	return b
}

func (b *selectBuilder) Paginate(pagination internal.Pagination) *selectBuilder {
	b.pagination = &pagination
	return b
}

func (b *selectBuilder) ToSQL() (sql string, args []any) {
	args = b.args
	columns := strings.Join(b.columns, ", ")

	var where string
	wheres := b.where
	if !b.includeSoftDelete && lo.Contains(b.columns, "deleted_at") {
		wheres = append(wheres, "deleted_at IS NULL")
	}
	if len(wheres) > 0 {
		where = fmt.Sprintf(" WHERE %s", strings.Join(wheres, " AND "))
	}

	var order string
	if b.order != nil {
		dir := "ASC"
		if strings.ToUpper(b.order.direction) == "DESC" {
			dir = "DESC"
		}
		order = fmt.Sprintf(" ORDER BY %s %s", b.order.column, dir)
	}

	var limitOffset string
	if b.pagination != nil {
		limitOffset = " LIMIT ? OFFSET ?"
		args = append(args, b.pagination.Limit, b.pagination.Offset)
	}
	sql = fmt.Sprintf(`SELECT %s FROM %s%s%s%s`, columns, b.table, where, order, limitOffset)
	numberedSQL := replaceWithOrderedParams(sql)
	log.V("SQL Query: %s %v", numberedSQL, args)
	return numberedSQL, args
}

func (b *selectBuilder) ScanDest() []any {
	return b.scanDest
}
