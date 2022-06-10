package models

type User struct {
	tableName struct{} `pg:"user,discard_unknown_columns"`
	ID        int      `pg:"id"`
	Name      string   `pg:"name"`
	Role      string   `pg:"role"`
}
