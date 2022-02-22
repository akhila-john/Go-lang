package models

type User struct {
	ID        int
	Firstname string
	Lastnsme  string
}

var (
	users  []User
	nextID = 1
)

func getUser() []User {
	return users
}

func addUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, u)

	return u, nil

}
