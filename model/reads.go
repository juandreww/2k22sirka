package model

import (
	"2k22sirka/views"
	"fmt"
	"log"
	// "github.com/davecgh/go-spew/spew"
	// "database/sql"
)

func ReadAll() ([]views.Users, error) {
	rows, err := con.Query("SELECT userid, name FROM trnkelapabakar")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	people := []views.Users{}
	for rows.Next() {
		users := views.Users{}
		
		if err := rows.Scan(&users.Userid, &users.Name); err != nil {
            log.Fatal(err)
        }

		fmt.Printf("hey %s you %.2f\n", users.Userid, users.Name)
	}
	
	return people, nil
}

func ReadSelected(uid string) ([]views.Users, error) {
	rows, err := con.Query("SELECT type2, quantity FROM trnkelapabakar WHERE uid = ($1)::uuid", uid)
	
	if rows == nil {
		fmt.Println("No rows returned")
	}
	if err != nil {
		return nil, err
	} 
	
	coconut := []views.Users{}
	// spew.Dump(coconut)
	
	for rows.Next() {
		data := views.Users{}
		
		rows.Scan(&data.Type2, &data.Quantity)
		coconut = append(coconut, data)
	}
	
	fmt.Println("CreateKelapa here...")
	
	return coconut, nil
}