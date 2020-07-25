package db

import (
  "context"
  "fmt"
  "github.com/jackc/pgx/v4"
)

type User struct {
  ID uint64
  Username string
  Salt string
  Password string
}

func (u *User) String() string {
  return fmt.Sprintf("User<%d, %s>", u.ID, u.Username)
}

func CreateUser(username, password string) (bool, error) {
  _, exists, err := GetUserByName(username)
  if err != nil {
    return false, err
  }

  if exists {
    return false, nil
  }

  return true, nil
}

func GetUserByName(username string) (User, bool, error) {
  u := User{}
  row := db.QueryRow(context.Background(), "SELECT * FROM users WHERE username=$1", username)

  err := row.Scan(&u.ID,
    &u.Username,
    &u.Salt,
    &u.Password)

  if err != nil {
    if err == pgx.ErrNoRows {
      return User{}, false, nil
    }

    return User{}, false, err
  }

  return u, true, nil
}

func GetUserByID(id uint64) (User, bool, error) {
  u := User{}
  row := db.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", id)

  err := row.Scan(&u.ID,
    &u.Username,
    &u.Salt,
    &u.Password)

  if err != nil {
    if err == pgx.ErrNoRows {
      return User{}, false, nil
    }

    return User{}, false, err
  }

  return u, true, nil
}