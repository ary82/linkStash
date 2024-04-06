package middleware

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/idtoken"
)

func GetData(tokenStr []byte) (*idtoken.Payload, error) {
	pay, err := idtoken.Validate(context.Background(), string(tokenStr), os.Getenv("AUDIENCE"))
	if err != nil {
		log.Println("err", err)
	}
	return pay, nil
}
