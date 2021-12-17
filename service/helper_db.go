package service

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	model "github.com/GuillaumeBergeronGeoffroy/chacra-api/model"
)

// ExecuteStatements exportable
func ExecuteStatements(db *sql.DB, stmts []string) (err error) {
	for _, stmt := range stmts {
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		_, err = db.ExecContext(ctx, stmt)
		if err != nil {
			fmt.Println("ResStatus:", err)
			return
		}
	}
	return
}

// SaveModel
func SaveModel(model model.Model, db *sql.DB) (err error) {
	ctx := context.TODO()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()
	if err = model.Validate(); err != nil {
		return
	}
	if err = model.BeforeSave(); err != nil {
		return
	}
	v := reflect.ValueOf(model)
	typeOfS := reflect.Indirect(v).Type()
	modelValues, modelValuesStmt, modelTypes, updateStmt := []string{}, "", "", ""
	for i := 0; i < reflect.Indirect(v).NumField(); i++ {
		modelValues = append(modelValues, fmt.Sprintln(reflect.Indirect(v).Field(i).Interface()))
		modelValuesStmt = fmt.Sprintln(modelValuesStmt + "?")
		modelTypes = fmt.Sprintln(modelTypes + typeOfS.Field(i).Name)
		updateStmt = fmt.Sprintln(updateStmt + typeOfS.Field(i).Name + " = ?")
		if i+1 < reflect.Indirect(v).NumField() {
			modelValuesStmt += " ,"
			modelTypes += ","
			updateStmt += ", "
		}
	}
	stmt := fmt.Sprintln("INSERT INTO " + fmt.Sprintln(reflect.TypeOf(model).Elem().Name()) + " (" + modelTypes + ") VALUES (" + modelValuesStmt + ") ON DUPLICATE KEY UPDATE " + updateStmt)
	modelValues = append(modelValues, modelValues...)
	modelValuesInterface := make([]interface{}, len(modelValues))
	for i, v := range modelValues {
		modelValuesInterface[i] = v
	}
	if _, err = tx.ExecContext(ctx, stmt, modelValuesInterface...); err != nil {
		fmt.Println(err)
		return
	}
	if err = model.AfterSave(); err != nil {
		return
	}
	err = tx.Commit()
	return
}
