package basic

import (
	"log"
	"myprotobuf/protogen/basic"
)

func BasicUser() {

	addr := basic.Address{
		Street:  "a",
		City:    "b",
		Country: "c",
		Coordinate: &basic.Address_Coordinate{
			Langitude: 40.23842478349234,
			Logitude:  -74.1238198237192,
		},
	}

	user := basic.User{
		Id:       69,
		Username: "indra",
		IsActive: true,
		Password: []byte("cemplon"),
		Emails:   []string{"indra123@gmal.com", "indra234@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	log.Println(&user)
}
