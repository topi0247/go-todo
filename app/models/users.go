package models

import (
    "log"
    "time"
)

type User struct {
    ID int
    UUID string
    Name string
    Email string
    Password string
    CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
    cmd := `insert into users (
        uuid,
        name,
        email,
        password,
        created_at) values ($1, $2, $3, $4, $5)`

    _, err = Db.Exec(cmd,
        createUUID(),
        u.Name,
        u.Email,
        Encrypt(u.Password),
        time.Now())

    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func GetUser(id int) (user User, err error) {
    user = User{}
    cmd := `select id, uuid, name, email, password, created_at
        from users
        where id = $1`
    err = Db.QueryRow(cmd, id).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt)
    return user, err
}

func (u *User) UpdateUser() (err error) {
    cmd := `update users
        set name = $1,
        email = $2
        where id = $3`
    _, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func (u *User) DeleteUser() (err error) {
    cmd := `delete from users where id = $1`
    _, err = Db.Exec(cmd, u.ID)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}