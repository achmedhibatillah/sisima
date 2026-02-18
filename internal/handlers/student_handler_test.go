package handler

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func getServer() string {
	return "http://" + os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT/")
}
func TestGetStudent(t *testing.T) {
	res, err := http.Get(getServer() + "/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
