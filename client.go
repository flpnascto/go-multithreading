package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	mainCtx := context.Background()
	ctx, cancel := context.WithTimeout(mainCtx, 1*time.Second)
	defer cancel()

	c1 := make(chan string)

	select {
	case <-ctx.Done():
		println("Tempo de 1 segundo excedido")
		return
	default:
		go fetchBrasilApi("13330250", c1)
		go fetchViaCepApi("13330250", c1)
	}

	readChannel(c1)
}

func fetchBrasilApi(cep string, c chan string) {
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
	result := "BrasilApi:\n" + string(body)
	channelReciver(result, c)
}

func fetchViaCepApi(cep string, c chan string) {
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
	result := "ViaCepApi:\n" + string(body)
	channelReciver(result, c)
}

func channelReciver(data string, c chan<- string) {
	c <- data
}

func readChannel(data chan string) {
	fmt.Println(<-data)
}
