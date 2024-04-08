package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	mainCtx := context.Background()
	ctx, cancel := context.WithTimeout(mainCtx, 1*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		println("Tempo de 1 segundo excedido")
		return
	default:
		fetchBrasilApi("13330250")
		fetchViaCepApi("13330250")
	}
}

func fetchBrasilApi(cep string) {
	endpoint := "https://brasilapi.com.br/api/cep/v1/" + cep
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
	// TODO escrever responsta no email.
}

func fetchViaCepApi(cep string) {
	endpoint := "http://viacep.com.br/ws/" + cep + "/json"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
	// TODO escrever responsta no email.
}
