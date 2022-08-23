package model

import (
	"2k22sirka/views"
	"fmt"
	"log"
	// "github.com/davecgh/go-spew/spew"
	// "database/sql"
)

func DisplayAllUsers() ([]views.Users, error) {
	rows, err := con.Query("SELECT userid, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	people := []views.Users{}
	for rows.Next() {
		users := views.Users{}
		
		rows.Scan(&users.Userid, &users.Name)
		people = append(people, users)
	}
	
	return people, nil
}

func DisplayUser(userid string) ([]views.Users, error) {
	rows, err := con.Query("SELECT type2, quantity FROM users WHERE userid = ($1)", userid)
	
	if rows == nil {
		fmt.Println("No rows returned")
	}
	if err != nil {
		return nil, err
	} 
	
	people := []views.Users{}
	
	for rows.Next() {
		data := views.Users{}
		
		rows.Scan(&data.Userid, &data.Name)
		people = append(people, data)
	}
	
	return people, nil
}