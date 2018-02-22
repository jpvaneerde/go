package jp

import ("github.com/mediocregopher/radix.v2/redis"
	"fmt"
       )
type ContactDetails struct {
	name      string
	age       string
	favDrink1 string
	favDrink2 string
}

func addContact(newContact ContactDetails) {
	conn, err := redis.Dial("tcp", "10.1.1.21:32769")
	if err != nil {
		// log.Fatal(err)
	}
	defer conn.Close()
	resp := conn.Cmd("HMSET", "friend:", "name", newContact.name, "Age", newContact.age, "favDrink1", newContact.favDrink1, "favDrink2", newContact.favDrink2)
	if resp.Err != nil {
		// log.Fatal(resp.Err)
	}
	fmt.Println(newContact.name)
}
