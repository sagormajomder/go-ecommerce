package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func GetUsers() []User {
	return users

}

func Find(email, pass string) *User {

	for _, user := range users {
		if user.Email == email && user.Password == pass {
			return &user
		}
	}

	return nil
}

func (u User) Store() *User {
	usr := Find(u.Email, u.Password)
	if usr != nil {
		return nil
	}

	u.ID = len(users) + 1
	users = append(users, u)
	return &u
}
