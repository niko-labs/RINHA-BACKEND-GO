package helpers

import "log"

func VerificaErro(err error) {
	if err != nil {
		panic(err)
	}
}

func VerificaErroComMsgLog(err error, msg string) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
