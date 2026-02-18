package test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/dvvnFrtn/sisima/internal/config"
	handler "github.com/dvvnFrtn/sisima/internal/handlers"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
)

func TestStudentHandler(t *testing.T) {
	db, err := config.ConnectDatabase()
	if err != nil {
		panic("can't connect to database")
	}

	service := service.NewStudentService(db)
	handler := handler.NewStudentHandler(service)

	app := fiber.New()
	app.Get("/student", handler.FindAllPaginated)

	req, _ := http.NewRequest("GET", "/student", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("app test failed: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, but got %d", res.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := buf.String()
	t.Logf("Response body : %s", body)
}
