// defines person struct and methods
// for creating, updating, deleting, and

package models

import "log"

type Person struct { 
	ID   uint 		`gorm:"primary_key auto_increment" json:"id"`
	Name string 	`json:"name"`
	Age  int 		`json:"age"`
	Birthday string `json:"birthday"`
	Email string 	`json:"email"`
}

func (p *Person) CelebrateBirthday() {
	log.Printf("age of person with id %d increased", p.ID)
	p.Age++
}

func (p *Person) ChangeMail(newMail string) {
	p.Email = newMail
}

