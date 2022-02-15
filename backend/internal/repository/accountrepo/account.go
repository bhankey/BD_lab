package accountrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bhankey/BD_lab/backend/internal/entities/accountentities"
	"github.com/bhankey/BD_lab/backend/internal/repository"
	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	*repository.BaseRepository

	db *sqlx.DB
}

func NewAccountRepo(baseRepo *repository.BaseRepository, db *sqlx.DB) *AccountRepo {
	return &AccountRepo{
		BaseRepository: baseRepo,
		db:             db,
	}
}

func (r *AccountRepo) Create(ctx context.Context, name string, userID int) error {
	const query = `
		INSERT INTO account (
				name, 
				user_id
		)
		VALUES (
				$1,
				$2
		)
`

	if _, err := r.db.ExecContext(ctx, query, name, userID); err != nil {
		r.Logger.Error("accountrepo.Create.QueryError", err)

		return err
	}

	return nil
}

func (r *AccountRepo) Update(ctx context.Context, accountID int, name string) error {
	const query = `
		UPDATE account 
			SET name = $1
		WHERE id = $2
`

	if _, err := r.db.ExecContext(ctx, query, name, accountID); err != nil {
		r.Logger.Error("accountrepo.Update.QueryError")

		return err
	}

	return nil
}

func (r *AccountRepo) Delete(ctx context.Context, accountID int) error {
	const query = `
		UPDATE account 
			SET is_show = false
		WHERE id = $1
`

	if _, err := r.db.ExecContext(ctx, query, accountID); err != nil {
		r.Logger.Error("accountrepo.Delete.QueryError")

		return err
	}

	return nil
}

func (r *AccountRepo) GetOne(ctx context.Context, id int) (accountentities.Account, error) {
	const query = `
		SELECT 
			id,
		    name,
		    user_id,
		    is_show,
		    sum
		FROM account
		WHERE id = $1 AND is_show = true
`

	row := account{}
	if err := r.db.GetContext(ctx, &row, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accountentities.Account{}, repository.ErrNoEntity
		}
		r.Logger.Error("accountrepo.GetOne.QueryError")

		return accountentities.Account{}, err
	}

	return accountentities.Account{
		ID:     row.ID,
		Name:   row.Name,
		Sum:    row.Sum,
		UserID: row.UserID,
		IsShow: row.IsShow,
	}, nil
}

func (r *AccountRepo) GetAll(ctx context.Context) ([]accountentities.Account, error) {
	const query = `
		SELECT 
			id,
		    name,
		    user_id,
		    is_show,
		    sum
		FROM account
		WHERE is_show = true
`

	rows := make([]account, 0)
	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		r.Logger.Error("accountrepo.getAll.QueryError")

		return nil, err
	}

	result := make([]accountentities.Account, 0, len(rows))
	for _, row := range rows {
		result = append(result, accountentities.Account{
			ID:     row.ID,
			Name:   row.Name,
			Sum:    row.Sum,
			UserID: row.UserID,
			IsShow: row.IsShow,
		})
	}

	return result, nil
}

func (r *AccountRepo) ChangeSum(ctx context.Context, accountID int, sum float64) error {
	const query = `
		UPDATE account 
			SET sum = $2
		WHERE id = $1
`

	if _, err := r.db.ExecContext(ctx, query, accountID, sum); err != nil {
		r.Logger.Error("accountrepo.ChangeSum.QueryError")

		return err
	}

	return nil
}
