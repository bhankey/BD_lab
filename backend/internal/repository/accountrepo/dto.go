package accountrepo

type account struct {
	ID     int     `db:"id"`
	Name   string  `db:"name"`
	UserID int     `db:"user_id"`
	IsShow bool    `db:"is_show"`
	Sum    float64 `db:"sum"`
}
