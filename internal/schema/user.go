package schema

type User struct {
	ID           int64  `bun:",pk,autoincriment"`
	Username     string `bun:",notnull,unique"`
	passwordhash string `bun:",notnull"`
}