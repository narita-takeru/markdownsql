package main

import (
	"fmt"
	"github.com/narita-takeru/markdownsql/sam"
	"github.com/narita-takeru/markdownsql/sqldef"
	"os"
	"strings"
)

func onTableColumns(tbl *sqldef.TableDefinition, columns map[string]string) error {

	nullable := false
	if columns[`null`] == `YES` {
		nullable = true
	}

	colDef := sqldef.ColumnDefinition{
		Name:    columns[`name`],
		Type:    columns[`type`],
		Null:    nullable,
		Default: columns[`default`],
		Key:     columns[`key`],
		Comment: columns[`comment`],
	}

	tbl.Columns = append(tbl.Columns, colDef)
	return nil
}

func onTableIndexes(tbl *sqldef.TableDefinition, columns map[string]string) error {
	idxDef := sqldef.IndexDefinition{
		Columns:  strings.Split(columns[`columns`], `,`),
		IsUnique: columns[`unique`] == `YES`,
	}

	tbl.Indexes = append(tbl.Indexes, idxDef)
	return nil
}

func main() {
	mdPath := os.Args[1]

	sqlDef := sqldef.SQLDefinition{}
	currentTableIdx := -1

	isColumnMode := false
	isIndexesMode := false

	sm := sam.SamParser{
		OnOneLines: map[string]func(line string) error{
			"#": func(line string) error {
				sqlDef.DatabaseName = line
				return nil
			},
			"##": func(line string) error {
				tblDef := sqldef.TableDefinition{Name: line}
				sqlDef.Tables = append(sqlDef.Tables, tblDef)
				currentTableIdx = currentTableIdx + 1
				return nil
			},
			"###": func(line string) error {
				isColumnMode = false
				isIndexesMode = false
				if line == "columns" {
					isColumnMode = true
					return nil
				}

				if line == "indexes" {
					isIndexesMode = true
					return nil
				}

				return fmt.Errorf("Unkown ### [%s]", line)
			},
		},
		OnTable: func(columns map[string]string) error {
			tblDef := &sqlDef.Tables[currentTableIdx]

			if isColumnMode {
				return onTableColumns(tblDef, columns)
			}

			if isIndexesMode {
				return onTableIndexes(tblDef, columns)
			}

			return nil
		},
	}

	if err := sm.Start(mdPath); err != nil {
		panic(err)
	}

	fmt.Println(sqlDef.ToSQLStmt())
}
