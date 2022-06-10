package models

type Dummy struct {
	tableName struct{} `pg:"dummy,discard_unknown_columns"`
	ID        int      `pg:"id"`
}
