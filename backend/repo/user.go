package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(email, pass string) (*User, error)
	List() ([]*User, error)
	Delete(id int) error
	Update(id int, u User) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (user *userRepo) Create(u User) (*User, error) {
	// if u.ID !=0{
	// 	return &u,nil
	// }

	usr, _ := user.Get(u.Email, u.Password)

	// if err != nil {
	// 	return nil, err
	// }

	if usr != nil {
		return nil, nil
	}

	u.ID = len(user.userList) + 1
	user.userList = append(user.userList, &u)
	return &u, nil
}

func (user *userRepo) Get(email, pass string) (*User, error) {

	for _, user := range user.userList {
		if user.Email == email && user.Password == pass {
			return user, nil
		}
	}

	return nil, nil
}

func (user *userRepo) List() ([]*User, error) {
	return user.userList, nil

}

func (user *userRepo) Delete(id int) error {
	var temList []*User
	for _, u := range user.userList {
		if u.ID != id {
			temList = append(temList, u)
		}
	}
	user.userList = temList

	return nil
}
func (user *userRepo) Update(id int, u User) (*User, error) {
	u.ID = id
	for idx, usr := range user.userList {
		if usr.ID == u.ID {
			user.userList[idx] = &u
		}
	}

	return &u, nil
}
