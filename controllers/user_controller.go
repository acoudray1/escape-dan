package controllers

import (
    "database/sql"
    
	"github.com/aicyp/escape-dan-back/models"
)

// GetAllUsers returns all the users registered in the database
// @return UsersList[] | error
func (db Database) GetAllUsers() (*models.UsersList, error) {
    list := &models.UsersList{}
    rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID DESC")
    if err != nil {
        return list, err
    }
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Mail)
        if err != nil {
            return list, err
        }
        list.Users = append(list.Users, user)
    }
    return list, nil
}

// AddUser inserts a new user in the database
// @return nil | error
func (db Database) AddUser(user *models.User) error {
    var id int
    query := `INSERT INTO users (name, phone, mail) VALUES ($1, $2, $3) RETURNING id`
    err := db.Conn.QueryRow(query, user.Name, user.Phone, user.Mail).Scan(&id)
    if err != nil {
        return err
    }
    user.Id = id
    return err
}

// GetUserById returns the user associated to the id
// @return User | error
func (db Database) GetUserById(userId int) (models.User, error) {
    user := models.User{}
    query := `SELECT * FROM users WHERE id = $1;`
    row := db.Conn.QueryRow(query, userId)
    switch err := row.Scan(&user.Id, &user.Name, &user.Phone, &user.Mail); err {
    case sql.ErrNoRows:
        return user, ErrNoMatch
    default:
        return user, err
    }
}

// DeleteUser deletes a user from the database by id
// @return nil | error
func (db Database) DeleteUser(userId int) (bool, error) {
    query := `DELETE FROM users WHERE id = $1;`
    _, err := db.Conn.Exec(query, userId)
    switch err {
    case sql.ErrNoRows:
        return false, ErrNoMatch
    default:
        return false, err
    }
    return true, nil
}

// UpdateUser updates a user from the database by id
// @return User | error
func (db Database) UpdateUser(userId int, userData models.User) (models.User, error) {
    user := models.User{}
    query := `UPDATE users SET name=$2, phone=$3, mail=$4 WHERE id=$1 RETURNING id, name, phone, mail;`
    err := db.Conn.QueryRow(query, userId, userData.Name, userData.Phone, userData.Mail).Scan(&user.Id, &user.Name, &user.Phone, &user.Mail)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, ErrNoMatch
        }
        return user, err
    }
    return user, nil
}