package cache

import "log"

type HandleErrFunc func(error)

func IgnoreErr(_ error) {}

func LogErr(err error) {
	log.Println(err)
}

func PanicErr(err error) {
	log.Panicln(err)
}

func FatalErr(err error) {
	log.Fatalln(err)
}
