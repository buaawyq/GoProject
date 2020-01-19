package Domain

import "database/sql"

type Card struct {
	Id       int
	Number   sql.NullString
	User     sql.NullString
	Password sql.NullString
}
