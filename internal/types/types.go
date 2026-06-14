package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
