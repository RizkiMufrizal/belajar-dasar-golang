package helper

import "database/sql"

func PanifIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanifIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanifIfError(errorCommit)
	}
}
