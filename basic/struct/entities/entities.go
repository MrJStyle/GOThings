package entities

type Person1 struct {
	Name  string
	email string
}

type person2 struct {
	Name  string
	Email string
}

type Admin struct {
	person2
	Level string
}
