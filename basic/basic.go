package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"myprotobuf/protogen/basic"
)

func BasicHello() {
	whatever := basic.Hello{Name: "indra octama"}
	log.Printf(whatever.Name)
}

func ProtoToJsonUser() {
	proto := basic.User{
		Id:       69,
		Username: "indra",
		IsActive: true,
		Password: []byte("cemplon"),
		Emails:   []string{"1@gmail.com", "2@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&proto)
	log.Println(string(jsonBytes))
}

func JsonToProto() {

	var (
		protoku basic.User
	)

	json := `
			{
			  "id": 69,
			  "username": "indra",
			  "is_active": true,
			  "password": "Y2VtcGxvbg==",
			  "emails": [
				"1@gmail.com",
				"2@gmail.com"
			  ],
			  "gender": "GENDER_MALE"
			}`

	err := protojson.Unmarshal([]byte(json), &protoku)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(&protoku)
}
