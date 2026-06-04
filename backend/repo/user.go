package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(email, pass string) (*User, error)
	List() ([]*User, error)
	Delete(id int) error
	Update(id int, u User) (*User, error)
}

type userRepo struct {
	// userList []*User
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(u User) (*User, error) {

	usr, err := r.Get(u.Email, u.Password)

	if err != nil {
		return nil, err
	}

	if usr != nil {
		return nil, nil
	}

	query := `
	INSERT INTO users (
	first_name, 
	last_name, 
	email, 
	password, 
	is_shop_owner
	) 
	VALUES (
	:first_name,
	:last_name, 
	:email, 
	:password, 
	:is_shop_owner
	) 
	RETURNING id
`

	rows, err := r.db.NamedQuery(query, u)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&u.ID)
	}

	return &u, nil
}

func (r *userRepo) Get(email, pass string) (*User, error) {
	var user User

	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`
	err := r.db.Get(&user, query, email, pass)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) List() ([]*User, error) {

	var usrs []*User

	query := `
	SELECT 
	id,
	first_name,
	last_name,
	email,
	is_shop_owner 
	FROM users
	`
	err := r.db.Select(&usrs, query)

	if err != nil {
		return nil, err
	}

	return usrs, nil

}

func (r *userRepo) Delete(id int) error {
	query := `
	DELETE FROM users 
	WHERE id=$1
	`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
func (r *userRepo) Update(id int, u User) (*User, error) {

	query := `
	UPDATE users SET 
	first_name=$1,
	last_name=$2,
	email=$3,
	is_shop_owner=$4
	WHERE id=$5
	`
	row := r.db.QueryRow(query, u.FirstName, u.LastName, u.Email, u.IsShopOwner, id)

	err := row.Err()

	if err != nil {
		return nil, err
	}

	return &u, nil
}
