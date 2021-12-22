package accountrepo

import (
	"context"
	"finance/internal/entities/accountentities"
	"github.com/jmoiron/sqlx"
)

func (r *AccountRepo) GetDebtors(ctx context.Context, accountIDs []int) ([]accountentities.Account, error) {
	const query = `
		SELECT 
			id,
		    name,
			user_id,
		    sum 
		FROM account
		WHERE is_show = true AND sum < 0 AND id IN (?);
`

	resultingQuery, params, err := sqlx.In(query, accountIDs)
	if err != nil {
		return nil, err
	}

	resultingQuery = r.db.Rebind(resultingQuery)
	rows := make([]account, 0)
	if err := r.db.SelectContext(ctx, &rows, resultingQuery, params...); err != nil {
		r.Logger.Error("accountrepo.GetDebtors.QueryError")

		return nil, err
	}

	result := make([]accountentities.Account, 0, len(rows))
	for _, row := range rows {
		result = append(result, accountentities.Account{
			ID:     row.ID,
			Name:   row.Name,
			Sum:    row.Sum,
			UserID: row.UserID,
			IsShow: true,
		})
	}

	return result, nil
}

func (r *AccountRepo) GetAllDebtors(ctx context.Context) ([]accountentities.Account, error) {
	const query = `
		SELECT 
			id,
		    name,
			user_id,
		    sum 
		FROM account
		WHERE is_show = true AND sum < 0
`

	rows := make([]account, 0)
	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		r.Logger.Error("accountrepo.GetAllDebtors.QueryError")

		return nil, err
	}

	result := make([]accountentities.Account, 0, len(rows))
	for _, row := range rows {
		result = append(result, accountentities.Account{
			ID:     row.ID,
			Name:   row.Name,
			Sum:    row.Sum,
			UserID: row.UserID,
			IsShow: true,
		})
	}

	return result, nil
}
