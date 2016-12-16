package helper

type User struct
{
	Name string
	Email string
	Password string `json:_`
}

func GetUserList() []User {
	return []User{
		{"Rob Pike", "r@google.com", "123456"},
		{"Toto", "toto@google.com", "AZERTY"},
		{"Jérémy MOULIN", "jaymoulin@gmail.com", "azerty"}}
}