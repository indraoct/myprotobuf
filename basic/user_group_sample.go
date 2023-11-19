package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"myprotobuf/protogen/basic"
)

func BasicUserGroup() {

	addr := basic.Address{
		Street:  "dfwd",
		City:    "wef",
		Country: "wefwf",
	}

	batman := basic.User{
		Id:       97,
		Username: "bobo",
		IsActive: true,
		Password: []byte("hdbfkjwebfbwefbw"),
		Emails:   []string{"sjkbfksbdf", "hjdsfwbejfw"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	robin := basic.User{
		Id:       96,
		Username: "robin",
		IsActive: true,
		Password: []byte("hdbfkjwebfbwefbw"),
		Emails:   []string{"sjkbfksbdf", "hjdsfwbejfw"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	batFamily := basic.UserGroup{
		GroupId:     999,
		GroupName:   "Batman Family",
		Roles:       []string{"admin", "editor", "viewer"},
		Users:       []*basic.User{&batman, &robin},
		Description: "ewer ewer",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)

	log.Println(string(jsonBytes))

}
