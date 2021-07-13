package db_client

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Query(statement *Statement) ([]map[string]interface{}, error)
}

type db struct {
	connection *sql.DB
}

func NewDBClient(connection *sql.DB) DB {
	return &db{
		connection: connection,
	}
}

func (db *db) Query(statement *Statement) ([]map[string]interface{}, error) {
	var rows *sql.Rows
	var err error

	rows, err = db.connection.Query("INSERT INTO `estate_agency`.`person` (`doc_number`, `name`, `last_name`) VALUES ('12345', 'Emi', 'Bent');\n")
	if err != nil {
		println(statement)
		return nil, errors.New("error making query to db")
	}
	defer rows.Close()
	return db.parseResults(rows)
}

func (db *db) parseResults(rows *sql.Rows) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, 0)
	if rows != nil {
		cols, _ := rows.Columns()
		for rows.Next() {
			dest := make([]sql.NullString, len(cols))
			destPtrs := make([]interface{}, len(cols))
			for i := range dest {
				destPtrs[i] = &dest[i]
			}

			err := rows.Scan(destPtrs...)
			if err != nil {
				return nil, errors.New("unexpect error parsing db results")
			}

			row := make(map[string]interface{})

			for i, colName := range cols {
				if dest[i].Valid {
					row[colName] = dest[i].String
				} else {
					row[colName] = nil
				}
			}

			results = append(results, row)
		}
		if rows.Err() != nil {
			return nil, errors.New("unexpect error iterating result set")
		}
	}

	return results, nil
}