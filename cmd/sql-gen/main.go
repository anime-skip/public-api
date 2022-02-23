// Generate all basic methods that perform database operations. This way they're consistent and can
// be easily updated

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"anime-skip.com/timestamps-service/internal"
	. "github.com/dave/jennifer/jen"
	"github.com/ettle/strcase"
)

const (
	SQL_GEN_TAG = "sql_gen"
	PRIMARY_KEY = "primary_key"
	GET_ONE     = "get_one"
	GET_MANY    = "get_many"
	SOFT_DELETE = "soft_delete"
)

var models = []interface{}{
	internal.APIClient{},
	internal.EpisodeURL{},
	internal.Episode{},
	internal.Preferences{},
	internal.ShowAdmin{},
	internal.Show{},
	internal.TemplateTimestamp{},
	internal.Template{},
	internal.TimestampType{},
	internal.Timestamp{},
	internal.User{},
}

var internalPkg = "anime-skip.com/timestamps-service/internal"
var asContextPkg = "anime-skip.com/timestamps-service/internal/context"
var asErrorsPkg = "anime-skip.com/timestamps-service/internal/errors"

func main() {
	println("SQL Generation")
	for _, model := range models {
		generateModelSqlMethods(model)
	}
}

type ModelDetails struct {
	softDelete     *reflect.StructField
	primaryKeys    []reflect.StructField
	getOneColumns  []reflect.StructField
	getManyColumns []reflect.StructField
}

func generateModelSqlMethods(model interface{}) {
	modelType := reflect.TypeOf(model)
	modelName := modelType.Name()
	filename := fmt.Sprintf("%s_repo.generated.go", strcase.ToSnake(modelName))
	fmt.Printf(" - %s -> %s\n", modelName, filename)
	modelDetails := getModelDetails(modelType)

	file := NewFile("postgres")
	file.HeaderComment("Code generated by cmd/sqlgen/main.go, DO NOT EDIT.")

	if len(modelDetails.primaryKeys) == 1 {
		// Primary key is just a get one, but ignoreing deleted at
		hardDeleteGetOne(file, modelType, modelDetails.primaryKeys[0])
	}

	for _, field := range modelDetails.getOneColumns {
		if modelDetails.softDelete != nil {
			softDeleteGetOneScoped(file, modelType, field, *modelDetails.softDelete)
			softDeleteGetOneUnscoped(file, modelType, field)
		} else {
			hardDeleteGetOne(file, modelType, field)
		}
	}

	for _, field := range modelDetails.getManyColumns {
		if modelDetails.softDelete != nil {
			softDeleteGetManyScoped(file, modelType, field, *modelDetails.softDelete)
			softDeleteGetManyUnscoped(file, modelType, field)
		} else {
			hardDeleteGetMany(file, modelType, field)
		}
	}

	insertInTx(file, modelType)
	insert(file, modelType)

	updateInTx(file, modelType)
	update(file, modelType)

	if len(modelDetails.primaryKeys) > 0 {
		deleteInTx(file, modelType, modelDetails.primaryKeys)
		delete(file, modelType)
	} else {
	}

	writeFile(filename, file)
}

// Utils

func writeFile(filename string, generatedFile *File) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%#v", generatedFile))
}

func getModelDetails(model reflect.Type) ModelDetails {
	fields := getModelFields(model)
	getOneColumns := []reflect.StructField{}
	getManyColumns := []reflect.StructField{}
	primaryKeys := []reflect.StructField{}
	var softDelete *reflect.StructField
	for _, field := range fields {
		tag := field.Tag.Get(SQL_GEN_TAG)
		if strings.Contains(tag, PRIMARY_KEY) {
			primaryKeys = append(primaryKeys, field)
		}
		if strings.Contains(tag, GET_ONE) {
			getOneColumns = append(getOneColumns, field)
		}
		if strings.Contains(tag, GET_MANY) {
			getManyColumns = append(getManyColumns, field)
		}
		if strings.Contains(tag, SOFT_DELETE) {
			fieldCopy := field
			softDelete = &fieldCopy
		}
	}
	return ModelDetails{
		softDelete:     softDelete,
		primaryKeys:    primaryKeys,
		getOneColumns:  getOneColumns,
		getManyColumns: getManyColumns,
	}
}

func getModelFields(t reflect.Type) (fields []reflect.StructField) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Name() == reflect.TypeOf(internal.BaseEntity{}).Name() {
			baseFields := getModelFields(field.Type)
			for _, baseField := range baseFields {
				fields = append(fields, baseField)
			}
		} else {
			fields = append(fields, field)
		}
	}
	return
}

func getColumnName(field reflect.StructField) string {
	dbTag := field.Tag.Get("db")
	if dbTag != "" {
		return dbTag
	}
	return strcase.ToSnake(field.Name)
}

func pluralize(str string) string {
	if strings.HasSuffix(str, "s") {
		return str
	}
	return str + "s"
}

func getSqlColumns(model reflect.Type) []string {
	fields := getModelFields(model)
	columns := []string{}
	for _, field := range fields {
		columns = append(columns, strcase.ToSnake(field.Name))
	}
	return columns
}

func getSqlPlaceholders(model reflect.Type) []string {
	fields := getModelFields(model)
	placeholders := []string{}
	for i := 1; i <= len(fields); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return placeholders
}

func getSqlParamValues(varName string, model reflect.Type) []Code {
	fields := getModelFields(model)
	values := []Code{}
	for _, field := range fields {

		values = append(values, Id(varName).Dot(field.Name))
	}
	return values
}

// Generated Functions

// Get One

func _getOne(file *File, funcName string, sql string, model reflect.Type, field reflect.StructField) {
	modelName := model.Name()
	varName := strcase.ToGoCamel(modelName)
	fieldName := strcase.ToGoCamel(field.Name)
	errMessage := fmt.Sprintf("%s.%s=", modelName, fieldName) + "%s"

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("db").Qual(internalPkg, "Database"),
		Id(fieldName).Qual(field.Type.PkgPath(), field.Type.Name()),
	).Params(
		Qual(internalPkg, modelName),
		Error(),
	).Block(
		Var().Id(varName).Qual(internalPkg, modelName),
		Err().Op(":=").Id("db").Dot("GetContext").Call(
			Id("ctx"),
			Op("&").Id(varName),
			Lit(sql),
			Id(fieldName),
		),
		If(Qual("errors", "Is").Call(Err(), Qual("database/sql", "ErrNoRows"))).Block(
			Return(
				Qual(internalPkg, modelName).Block(),
				Qual(asErrorsPkg, "NewRecordNotFound").Call(
					Qual("fmt", "Sprintf").Call(Lit(errMessage), Id(fieldName)),
				),
			),
		),
		Return(
			Id(varName),
			Err(),
		),
	).Line()
}

func _getOneNoSoftDelete(file *File, funcName string, model reflect.Type, field reflect.StructField) {
	modelName := model.Name()
	columnName := getColumnName(field)
	tableName := pluralize(strcase.ToSnake(modelName))
	sql := fmt.Sprintf(`SELECT * FROM %s WHERE %s=$1`, tableName, columnName)
	_getOne(file, funcName, sql, model, field)
}

func hardDeleteGetOne(file *File, model reflect.Type, field reflect.StructField) {
	funcName := fmt.Sprintf("get%sBy%s", model.Name(), strcase.ToGoPascal(field.Name))
	_getOneNoSoftDelete(file, funcName, model, field)
}

func softDeleteGetOneUnscoped(file *File, model reflect.Type, field reflect.StructField) {
	funcName := fmt.Sprintf("getUnscoped%sBy%s", model.Name(), strcase.ToGoPascal(field.Name))
	_getOneNoSoftDelete(file, funcName, model, field)
}

func softDeleteGetOneScoped(
	file *File,
	model reflect.Type,
	field reflect.StructField,
	softDeleteField reflect.StructField,
) {
	modelName := model.Name()
	columnName := getColumnName(field)
	softDeleteColumnName := getColumnName(softDeleteField)
	funcName := fmt.Sprintf("get%sBy%s", modelName, strcase.ToGoPascal(field.Name))
	tableName := pluralize(strcase.ToSnake(modelName))
	sql := fmt.Sprintf(
		`SELECT * FROM %s WHERE %s=$1 AND %s IS NULL`,
		tableName,
		columnName,
		softDeleteColumnName,
	)

	_getOne(file, funcName, sql, model, field)
}

// Get Many

func _getMany(file *File, funcName string, sql string, model reflect.Type, field reflect.StructField) {
	modelName := model.Name()
	itemName := strcase.ToGoCamel(modelName)
	varName := pluralize(itemName)
	fieldName := strcase.ToGoCamel(field.Name)
	internalQual := model.PkgPath()
	returnErr := If(Err().Op("!=").Nil()).Block(
		Return().List(Nil(), Err()),
	)

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("db").Qual(internalQual, "Database"),
		Id(fieldName).Qual(field.Type.PkgPath(), field.Type.Name()),
	).Params(
		Index().Qual(internalQual, modelName),
		Error(),
	).Block(
		List(Id("rows"), Err()).Op(":=").Id("db").Dot("QueryxContext").Call(Id("ctx"), Lit(sql)),
		returnErr,
		Defer().Id("rows").Dot("Close").Call(),
		Line(),

		Id(varName).Op(":=").Index().Qual(internalQual, modelName).Values(),
		For(Id("rows").Dot("Next").Call()).Block(
			Var().Id(itemName).Qual(internalQual, modelName),
			Err().Op("=").Id("rows").Dot("StructScan").Call(Op("&").Id(itemName)),
			returnErr,
			Id(varName).Op("=").Append(Id(varName), Id(itemName)),
		),
		Return(
			Id(varName),
			Nil(),
		),
	).Line()
}

func _getManyIgnoreSoftDelete(file *File, funcName string, model reflect.Type, field reflect.StructField) {
	modelName := model.Name()
	tableName := pluralize(strcase.ToSnake(modelName))
	sql := fmt.Sprintf(`SELECT * FROM %s`, tableName)

	_getMany(file, funcName, sql, model, field)
}

func hardDeleteGetMany(file *File, model reflect.Type, field reflect.StructField) {
	funcName := fmt.Sprintf("get%sBy%s", pluralize(model.Name()), strcase.ToGoPascal(field.Name))

	_getManyIgnoreSoftDelete(file, funcName, model, field)
}

func softDeleteGetManyUnscoped(file *File, model reflect.Type, field reflect.StructField) {
	funcName := fmt.Sprintf("getUnscoped%sBy%s", pluralize(model.Name()), strcase.ToGoPascal(field.Name))

	_getManyIgnoreSoftDelete(file, funcName, model, field)
}

func softDeleteGetManyScoped(
	file *File,
	model reflect.Type,
	field reflect.StructField,
	softDeleteField reflect.StructField,
) {
	modelName := model.Name()
	funcName := fmt.Sprintf("get%sBy%s", pluralize(modelName), strcase.ToGoPascal(field.Name))
	tableName := pluralize(strcase.ToSnake(modelName))
	sql := fmt.Sprintf(
		`SELECT * FROM %s WHERE %s IS NULL`,
		tableName,
		strcase.ToSnake(softDeleteField.Name),
	)

	_getMany(file, funcName, sql, model, field)
}

// Insert

func updateMetadataTime(id string, timeFieldName string) *Statement {
	return Id(id).Dot(timeFieldName).Op("=").Qual("time", "Now").Call()
}
func updateMetadataUserID(id string, userIDFieldName string) *Statement {
	return Id(id).Dot(userIDFieldName).Op("=").Id("claims").Dot("UserID")
}
func updateToNil(id string, nilFieldName string) *Statement {
	return Id(id).Dot(nilFieldName).Op("=").Nil()
}

func unwrapInTxFunc(file *File, model reflect.Type, inTxFuncName string) {
	modelName := model.Name()
	funcName := strings.Replace(inTxFuncName, "InTx", "", 1)
	argName := strcase.ToGoCamel(modelName)
	resultName := "result"
	emptyModel := Qual(internalPkg, modelName).Block()
	ifErrReturn := If(Err().Op("!=").Nil()).Block(
		Return().List(emptyModel, Err()),
	)

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("db").Qual(internalPkg, "Database"),
		Id(argName).Qual(internalPkg, modelName),
	).Params(
		Qual(internalPkg, modelName),
		Error(),
	).Block(
		List(Id("tx"), Err()).Op(":=").Id("db").Dot("BeginTxx").Call(Id("ctx"), Nil()),
		ifErrReturn,
		Defer().Id("tx").Dot("Rollback").Call(),
		Empty(),
		List(Id(resultName), Err()).Op(":=").Id(inTxFuncName).Call(Id("ctx"), Id("tx"), Id(argName)),
		ifErrReturn,
		Empty(),
		Id("tx").Dot("Commit").Call(),
		Return().List(Id(resultName), Nil()),
	).Line()
}

func insertInTx(file *File, model reflect.Type) {
	modelName := model.Name()
	funcName := fmt.Sprintf("insert%sInTx", modelName)
	argName := strcase.ToGoCamel(modelName)
	newModelName := strcase.ToGoCamel("new" + modelName)
	tableName := pluralize(strcase.ToSnake(modelName))
	columns := getSqlColumns(model)
	placeholders := getSqlPlaceholders(model)
	paramValues := getSqlParamValues(newModelName, model)
	sql := fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	execParams := []Code{Line().Id("ctx"), Line().Lit(sql)}
	execParams = append(execParams, Line().List(paramValues...))
	execParams = append(execParams, Line())
	emptyModel := Qual(internalPkg, modelName).Block()
	ifErrReturn := If(Err().Op("!=").Nil()).Block(
		Return().List(emptyModel, Err()),
	)
	_, hasCreatedAt := model.FieldByName("CreatedAt")
	_, hasCreatedBy := model.FieldByName("CreatedByUserID")
	_, hasUpdatedAt := model.FieldByName("UpdatedAt")
	_, hasUpdatedBy := model.FieldByName("UpdatedByUserID")
	_, hasDeletedAt := model.FieldByName("DeletedAt")
	_, hasDeletedBy := model.FieldByName("DeletedByUserID")

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("tx").Qual(internalPkg, "Tx"),
		Id(argName).Qual(internalPkg, modelName),
	).Params(
		Qual(internalPkg, modelName),
		Error(),
	).BlockFunc(func(g *Group) {
		g.Id(newModelName).Op(":=").Id(argName)
		if hasCreatedBy || hasUpdatedBy {
			g.List(Id("claims"), Err()).Op(":=").Qual(asContextPkg, "GetAuthClaims").Call(Id("ctx"))
			g.Add(ifErrReturn)
		}
		if hasCreatedAt {
			g.Add(updateMetadataTime(newModelName, "CreatedAt"))
		}
		if hasCreatedBy {
			g.Add(updateMetadataUserID(newModelName, "CreatedByUserID"))
		}
		if hasUpdatedAt {
			g.Add(updateMetadataTime(newModelName, "UpdatedAt"))
		}
		if hasUpdatedBy {
			g.Add(updateMetadataUserID(newModelName, "UpdatedByUserID"))
		}
		if hasDeletedAt {
			g.Add(updateToNil(newModelName, "DeletedAt"))
		}
		if hasDeletedBy {
			g.Add(updateToNil(newModelName, "DeletedByUserID"))
		}
		g.List(Id("result"), Err()).Op(":=").Id("tx").Dot("ExecContext").Call(execParams...)
		g.Add(ifErrReturn)
		g.List(Id("changedRows"), Err()).Op(":=").Id("result").Dot("RowsAffected").Call()
		g.Add(ifErrReturn)
		g.If(Id("changedRows").Op("!=").Lit(1)).Block(
			Return().List(emptyModel, Qual("fmt", "Errorf").Call(Lit("Inserted more than 1 row (%d)"), Id("changedRows"))),
		)
		g.Return().List(Id(newModelName), Err())
	}).Line()
}

func insert(file *File, model reflect.Type) {
	inTxFuncName := fmt.Sprintf("insert%sInTx", model.Name())
	unwrapInTxFunc(file, model, inTxFuncName)
}

// Update

func updateInTx(file *File, model reflect.Type) {
	modelName := model.Name()
	funcName := fmt.Sprintf("update%sInTx", modelName)
	argName := strcase.ToGoCamel("new" + modelName)
	updatedModelName := strcase.ToGoCamel("updated" + modelName)
	tableName := pluralize(strcase.ToSnake(modelName))
	columns := getSqlColumns(model)
	placeholders := getSqlPlaceholders(model)
	sqlSets := []string{}
	for i := 0; i < len(columns); i++ {
		sqlSets = append(sqlSets, fmt.Sprintf("%s=%s", columns[i], placeholders[i]))
	}
	paramValues := getSqlParamValues(updatedModelName, model)
	sql := fmt.Sprintf(
		"UPDATE %s SET %s",
		tableName,
		strings.Join(sqlSets, ", "),
	)
	execParams := []Code{Line().Id("ctx"), Line().Lit(sql)}
	execParams = append(execParams, Line().List(paramValues...))
	execParams = append(execParams, Line())
	emptyModel := Qual(internalPkg, modelName).Block()
	ifErrReturn := If(Err().Op("!=").Nil()).Block(
		Return().List(emptyModel, Err()),
	)
	_, hasUpdatedAt := model.FieldByName("UpdatedAt")
	_, hasUpdatedBy := model.FieldByName("UpdatedByUserID")

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("tx").Qual(internalPkg, "Tx"),
		Id(argName).Qual(internalPkg, modelName),
	).Params(
		Qual(internalPkg, modelName),
		Error(),
	).BlockFunc(func(g *Group) {
		g.Id(updatedModelName).Op(":=").Id(argName)
		if hasUpdatedBy {
			g.List(Id("claims"), Err()).Op(":=").Qual(asContextPkg, "GetAuthClaims").Call(Id("ctx"))
			g.Add(ifErrReturn)
		}
		if hasUpdatedAt {
			g.Add(updateMetadataTime(updatedModelName, "UpdatedAt"))
		}
		if hasUpdatedBy {
			g.Add(updateMetadataUserID(updatedModelName, "UpdatedByUserID"))
		}
		g.List(Id("result"), Err()).Op(":=").Id("tx").Dot("ExecContext").Call(execParams...)
		g.Add(ifErrReturn)
		g.List(Id("changedRows"), Err()).Op(":=").Id("result").Dot("RowsAffected").Call()
		g.Add(ifErrReturn)
		g.If(Id("changedRows").Op("!=").Lit(1)).Block(
			Return().List(emptyModel, Qual("fmt", "Errorf").Call(Lit("Updated more than 1 row (%d)"), Id("changedRows"))),
		)
		g.Return().List(Id(updatedModelName), Err())
	}).Line()
}

func update(file *File, model reflect.Type) {
	inTxFuncName := fmt.Sprintf("update%sInTx", model.Name())
	unwrapInTxFunc(file, model, inTxFuncName)
}

// Delete

func deleteInTx(file *File, model reflect.Type, primaryKeys []reflect.StructField) {
	modelName := model.Name()
	funcName := fmt.Sprintf("delete%sInTx", modelName)
	argName := strcase.ToGoCamel("new" + modelName)
	deletedModelName := strcase.ToGoCamel("deleted" + modelName)
	tableName := pluralize(strcase.ToSnake(modelName))
	sqlWheres := []string{}
	for index, primaryKey := range primaryKeys {
		sqlWheres = append(sqlWheres, fmt.Sprintf("%s=$%d", getColumnName(primaryKey), index+1))
	}
	sql := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		tableName,
		strings.Join(sqlWheres, " AND "),
	)
	execArgs := []Code{Id("ctx"), Lit(sql)}
	for _, primaryKey := range primaryKeys {
		execArgs = append(execArgs, Id(deletedModelName).Dot(primaryKey.Name))
	}
	emptyModel := Qual(internalPkg, modelName).Block()
	ifErrReturn := If(Err().Op("!=").Nil()).Block(
		Return().List(emptyModel, Err()),
	)
	_, hasUpdatedAt := model.FieldByName("UpdatedAt")
	_, hasUpdatedBy := model.FieldByName("UpdatedByUserID")
	_, hasDeletedAt := model.FieldByName("DeletedAt")
	_, hasDeletedBy := model.FieldByName("DeletedByUserID")

	file.Func().Id(funcName).Params(
		Id("ctx").Qual("context", "Context"),
		Id("tx").Qual(internalPkg, "Tx"),
		Id(argName).Qual(internalPkg, modelName),
	).Params(
		Qual(internalPkg, modelName),
		Error(),
	).BlockFunc(func(g *Group) {
		g.Id(deletedModelName).Op(":=").Id(argName)
		if hasUpdatedBy {
			g.List(Id("claims"), Err()).Op(":=").Qual(asContextPkg, "GetAuthClaims").Call(Id("ctx"))
			g.Add(ifErrReturn)
		}
		if hasUpdatedAt {
			g.Add(updateMetadataTime(deletedModelName, "UpdatedAt"))
		}
		if hasUpdatedBy {
			g.Add(updateMetadataUserID(deletedModelName, "UpdatedByUserID"))
		}
		if hasDeletedAt {
			g.Id("now").Op(":=").Qual("time", "Now").Call()
			g.Add(Id(deletedModelName).Dot("DeletedAt").Op("=").Op("&").Id("now"))
		}
		if hasDeletedBy {
			g.Id(deletedModelName).Dot("DeletedByUserID").Op("=").Op("&").Id("claims").Dot("UserID")
		}
		g.List(Id("result"), Err()).Op(":=").Id("tx").Dot("ExecContext").Call(execArgs...)
		g.Add(ifErrReturn)
		g.List(Id("changedRows"), Err()).Op(":=").Id("result").Dot("RowsAffected").Call()
		g.Add(ifErrReturn)
		g.If(Id("changedRows").Op("!=").Lit(1)).Block(
			Return().List(emptyModel, Qual("fmt", "Errorf").Call(Lit("Deleted more than 1 row (%d)"), Id("changedRows"))),
		)
		g.Return().List(Id(deletedModelName), Err())
	}).Line()
}

func delete(file *File, model reflect.Type) {
	inTxFuncName := fmt.Sprintf("delete%sInTx", model.Name())
	unwrapInTxFunc(file, model, inTxFuncName)
}
