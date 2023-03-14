package vsqlx

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// ReleaseTx clean db transaction by commit if no error, or rollback if an error occurred
func ReleaseTx(tx *sqlx.Tx, err *error) {
	if *err != nil {
		// If an error occurred, rollback transaction
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic(fmt.Errorf("failed to rollback database transaction.\n  > %w", errRollback))
		}
		return
	}

	// Else, commit transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		panic(fmt.Errorf("failed to commit database transaction\n  > %w", errCommit))
	}
}

func HandleErrorRepository(errRepo error, errMsg error) error {
	if errors.Is(errRepo, sql.ErrNoRows) {
		return errMsg
	}
	return errRepo
}
