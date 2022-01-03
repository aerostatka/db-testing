package domains

import "database/sql"

type Process struct {
	Id           int            `db:"ID"`
	User         sql.NullString `db:"USER"`
	Host         sql.NullString `db:"HOST"`
	Db           sql.NullString `db:"DB"`
	Command      sql.NullString `db:"COMMAND"`
	Time         int            `db:"TIME"`
	State        sql.NullString `db:"STATE"`
	Info         sql.NullString `db:"INFO"`
	TimeMs       int            `db:"TIME_MS"`
	RowsSent     int            `db:"ROWS_SENT"`
	RowsExamined int            `db:"ROWS_EXAMINED"`
}
