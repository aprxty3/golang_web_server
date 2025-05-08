package helper

import "github.com/jmoiron/sqlx"

func CommitOrRollback(tx *sqlx.Tx) {
	err := recover()
	if err != nil {
		// Rollback the transaction if a panic occurred
		errorRollback := tx.Rollback()
		// Panic if the rollback itself fails
		PanicIfError(errorRollback)
		// Re-panic with the original error to propagate it up the call stack
		// This is necessary to ensure the original error is handled by the error handler
		panic(err)
	} else {
		// Commit the transaction if no panic occurred
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
