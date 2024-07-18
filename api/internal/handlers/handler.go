package handlers

import (
	"net/http"

	"github.com/JoaoPauloFontana/data-import/api/internal/auth"
	"github.com/JoaoPauloFontana/data-import/api/internal/db"
	"github.com/JoaoPauloFontana/data-import/api/internal/models"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Username == "test" && u.Password == "password" {
		token, err := auth.GenerateJWT(u.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Erro ao gerar o token",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Credenciais inválidas",
	})
}

func GetRecords(c echo.Context) error {
	records := []models.Record{}
	err := db.DB.Select(&records, "SELECT * FROM records")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Erro ao buscar os registros",
		})
	}
	return c.JSON(http.StatusOK, records)
}
