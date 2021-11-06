package service

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
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
	err = model.Validate()
	if err != nil {
		return
	}
	v := reflect.ValueOf(model)
	typeOfS := v.Type()
	modelTypes, modelValues, updateStmt := "", "", ""
	for i := 0; i < v.NumField(); i++ {
		modelTypes = fmt.Sprintln(modelTypes + typeOfS.Field(i).Name + ",")
		modelValues = fmt.Sprintln(modelValues + fmt.Sprintln(v.Field(i).Interface()) + "','")
		updateStmt = fmt.Sprintln(updateStmt + typeOfS.Field(i).Name + "= '" + fmt.Sprintln(v.Field(i).Interface()) + "', ")
	}
	stmt := fmt.Sprintln("INSERT INTO " + fmt.Sprintln(reflect.TypeOf(model)) + " (" + strings.TrimSuffix(modelTypes, ",") + ") VALUES ('" + strings.TrimSuffix(modelValues, ",'") + ") ON DUPLICATE KEY UPDATE " + strings.TrimSuffix(updateStmt, ","))
	err = ExecuteStatements(db, []string{stmt})
	return
}
