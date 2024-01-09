package model

type User struct {
	Id       string
	Name     string
	Email    string
	Banner   string
	Profile  string
	Category string
}

type UserPassword struct {
	User
	Password string
}

type UMKM struct {
	User
	Verified bool
}

type Mahasiswa struct {
	User
	Badge bool
}

type EmailPassword struct {
	Id       string
	Email    string
	Password string
	Category string
}
