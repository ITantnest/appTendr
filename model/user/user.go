// Package user provides access to the user table in the MySQL database.
package user

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"archive/zip"
)

var (
	// table is the table name.
	table = "user"
)

// Item defines the model.
type Item struct {
	ID        			uint32	`db:"id"`
	Country	  			string	`db:"country"`
	Registr_country		string	`db:"registr_country"`
	Comp_name			string	`db:"comp_name"`
	Edrp_ipn			string	`db:"edrp_ipn"`
	Cash_manager		string	`db:"cash_manager"`
	Customer_type		string	`db:"customer_type"`
	Fio					string	`db:"fio"`
	Position			string	`db:"position"`
	Document			string	`db:"document"`
	Region				string	`db:"region"`
	City				string  `db:"city"`
	Adress				string  `db:"adress"`
	Zip					string 	`db:"zip"`
	LastName  			string  `db:"last_name"`
	FirstName 			string  `db:"first_name"`
	Middle_name			string  `db:"middle_name"`
	Email     			string  `db:"email"`
	Phone				string	`db:"phone"`
	Fax					string	`db:"fax"`
	Mobile				string 	`db:"mobile"`
	Site				string	`db:"site"`
	Password  			string  `db:"password"`
	StatusID  			uint8   `db:"status_id"`
	CreatedAt mysql.NullTime 	`db:"created_at"`
	UpdatedAt mysql.NullTime 	`db:"updated_at"`
	DeletedAt mysql.NullTime 	`db:"deleted_at"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByEmail gets user information from email.
func ByEmail(db Connection, email string) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
		SELECT id, password, status_id, first_name
		FROM %v
		WHERE email = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		email)
	return result, err == sql.ErrNoRows, err
}

// Create creates user.
func Create(db Connection,registr_country , comp_name , edrp_ipn , cash_manager , customer_type , fio , position , document , region , city , adress , zip , last_name , first_name ,  middle_name ,  email , phone , fax , mobile , site, password string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(registr_country , comp_name , edrp_ipn , cash_manager , customer_type , fio , position , document , region , city , adress , zip , last_name , first_name ,  middle_name ,  email , phone , fax , mobile , site, password)
		VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
		`, table),
		registr_country , comp_name , edrp_ipn , cash_manager , customer_type , fio , position , document , region , city , adress , zip , last_name , first_name ,  middle_name ,  email , phone , fax , mobile , site, password)
	return result, err
}
