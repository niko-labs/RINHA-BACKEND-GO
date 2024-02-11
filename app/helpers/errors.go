package helpers

import "log"

func ErroLancaPanico(err error) {
	if err != nil {
		panic(err)
	}
}

func ErroLancaLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
