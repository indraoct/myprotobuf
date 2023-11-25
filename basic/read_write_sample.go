package basic

import (
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func WriteProtoFile(msg proto.Message, fname string) {

	b, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Can't marshal message", err)
	}

	if err = os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can't write message", err)
	}
}

func ReadProtofromFile(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	if err = proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal file", err)
	}

}
