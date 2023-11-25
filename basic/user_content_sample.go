package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"myprotobuf/protogen/basic"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/this-is-v1",
		//Title:         "whdqjwdbqkjwhdbqkhjwdbq",
		HtmlContent: "<p> just dummy </p>",
		//AuthorId:      99,
	}

	WriteProtoFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")

	uc := basic.UserContent{}
	ReadProtofromFile("user_content_v1.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))

}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1002,
		Slug:          "/this-is-v2",
		//Title:         "whdqjwdbqkjwhdbqkhjwdbq V2",
		HtmlContent: "<p> just dummy V2</p>",
		//AuthorId:      99,
		//Category: "NEWS",
	}

	WriteProtoFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")

	uc := basic.UserContent{}
	ReadProtofromFile("user_content_v2.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))

}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1003,
		Slug:          "/this-is-v3",
		HtmlContent:   "<p> just dummy V3</p>",
		//Category:      "NEWS",
	}

	WriteProtoFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read V3")

	uc := basic.UserContent{}
	ReadProtofromFile("user_content_v3.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))

}
