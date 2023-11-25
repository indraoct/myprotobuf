package basic

import (
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"myprotobuf/protogen/basic"
	"myprotobuf/protogen/news"
	"time"
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
		LastKnownLocation: &latlng.LatLng{
			Latitude:  40.23842478349234,
			Longitude: -74.1238198237192,
		},
		LastLoginTimestamp: timestamppb.Now(),
		BirthDate: &date.Date{
			Year:  1988,
			Month: 10,
			Day:   21,
		},
	}

	log.Println(&user)
}

func BasicUserWithNews() {

	addr := basic.Address{
		Street:  "a",
		City:    "b",
		Country: "c",
		Coordinate: &basic.Address_Coordinate{
			Langitude: 40.23842478349234,
			Logitude:  -74.1238198237192,
		},
	}

	var anyNews anypb.Any
	anyNews = randomNews()

	user := basic.User{
		Id:       69,
		Username: "indra",
		IsActive: true,
		Password: []byte("cemplon"),
		Emails:   []string{"indra123@gmal.com", "indra234@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
		News:     &anyNews,
	}

	log.Println(&user)
}

func BasicUserWithNews2() {

	social := news.SocialMedia{
		SocialId:   99,
		SocialNews: "pew pew pews",
	}

	//msg := news.InstantMessaging{
	//	MsgId:   99,
	//	MsgNews: "ki ki ki",
	//}

	news2 := basic.User_SocialMedia{SocialMedia: &social}

	user := basic.User{
		Id:       69,
		Username: "indra",
		IsActive: true,
		Password: []byte("cemplon"),
		Emails:   []string{"indra123@gmal.com", "indra234@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		News2:    &news2,
	}

	log.Println(&user)
}

func randomNews() (anyNews anypb.Any) {
	rand.NewSource(time.Now().UnixNano())

	paper := news.PaperMail{
		PaperId:   1,
		PaperNews: "hai hai hai",
	}

	social := news.SocialMedia{
		SocialId:   99,
		SocialNews: "pew pew pews",
	}

	msg := news.InstantMessaging{
		MsgId:   99,
		MsgNews: "ki ki ki",
	}

	switch r := rand.Intn(10) % 3; r {

	case 0:
		anypb.MarshalFrom(&anyNews, &paper, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&anyNews, &social, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&anyNews, &msg, proto.MarshalOptions{})
	}

	return
}
