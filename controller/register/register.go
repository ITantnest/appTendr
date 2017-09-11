// Package register handles the user creation.
package register

import (
	"errors"
	"net/http"

	"github.com/blue-jay/blueprint/lib/flight"
	"github.com/blue-jay/blueprint/middleware/acl"
	"github.com/blue-jay/blueprint/model/user"

	"github.com/blue-jay/core/form"
	"github.com/blue-jay/core/passhash"
	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/register", Index, acl.DisallowAuth)
	router.Post("/register", Store, acl.DisallowAuth)
}

// Index displays the register page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("register/index")
	form.Repopulate(r.Form, v.Vars, "country","registr_country","comp_name","edrp_ipn","cash_manager","customer_type","fio","position","document","region","city","adress","zip","last_name","first_name", "middle_name", "email","phone","fax","mobile","site")
	v.Render(w, r)
}

// Store handles the registration form submission.
func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Validate with required fields
	if !c.FormValid("country","registr_country","comp_name","edrp_ipn","cash_manager","customer_type","fio","position","document","region","city","adress","zip","last_name","first_name", "middle_name", "email","phone","fax","mobile","site", "password", "password_verify") {
		Index(w, r)
		return
	}

	// Get form values
	country := r.FormValue("country")
	registr_country := r.FormValue("registr_country")
	comp_name := r.FormValue("comp_name")
	edrp_ipn := r.FormValue("edrp_ipn")
	cash_manager := r.FormValue("cash_manager")
	customer_type := r.FormValue("customer_type")
	fio := r.FormValue("fio")
	position := r.FormValue("position")
	document := r.FormValue("document")
	region := r.FormValue("region")
	city := r.FormValue("city")
	adress := r.FormValue("adress")
	zip := r.FormValue("zip")
	lastName := r.FormValue("last_name")
	firstName := r.FormValue("first_name")
	middle_name :=r.FormValue("middle_name")
	email := r.FormValue("email")
	phone :=r.FormValue("phone")
	fax := r.FormValue("fax")
	mobile := r.FormValue("mobile")
	site := r.FormValue("site")

	// Validate passwords
	if r.FormValue("password") != r.FormValue("password_verify") {
		c.FlashError(errors.New("Passwords do not match."))
		Index(w, r)
		return
	}

	// Hash password
	password, errp := passhash.HashString(r.FormValue("password"))

	// If password hashing failed
	if errp != nil {
		c.FlashErrorGeneric(errp)
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// Get database result
	_, noRows, err := user.ByEmail(c.DB, email)

	if noRows { // If success (no user exists with that email)
		_, err = user.Create(c.DB,  edrp_ipn,last_name, first_name, email, password)
		// Will only error if there is a problem with the query
		if err != nil {
			c.FlashErrorGeneric(err)
		} else {
			c.FlashSuccess("Account created successfully for: " + email)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	} else if err != nil { // Catch all other errors
		c.FlashErrorGeneric(err)
	} else { // Else the user already exists
		c.FlashError(errors.New("Account already exists for: " + email))
	}

	// Display the page
	Index(w, r)
}
