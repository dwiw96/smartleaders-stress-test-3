package repository

import "fmt"

func (r *rentRepository) ExecDbTx(fn func(*rentRepository) error) error {
	tx, err := r.pool.Begin(r.ctx)
	if err != nil {
		return err
	}

	err = fn(r)
	if err != nil {
		if rbErr := tx.Rollback(r.ctx); rbErr != nil {
			return fmt.Errorf("--- [error](execTx) fn err = %s, and rb err = %s", err, rbErr)
		}
		return err
	}

	return tx.Commit(r.ctx)
}
