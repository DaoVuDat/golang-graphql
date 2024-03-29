package main

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/sqlStore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Token interface{} `json:"token"`
}

const filepath = "./data/db.sqlite3"

func main() {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/login", func(c echo.Context) error {
		var loginReq LoginReq

		if err := c.Bind(&loginReq); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		user, err := sqlStore.FindUserByEmail(db, loginReq.Email)
		if err != nil {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}

		if user.Password != loginReq.Password {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}

		token, err := NewClaim(user.ID, user.Email)
		if user.Password != loginReq.Password {
			return c.String(http.StatusInternalServerError, "internal error")
		}

		claimToken, err := SignClaimToken(token)
		if err != nil {
			return c.String(http.StatusInternalServerError, "internal error claim")
		}

		// return token instead
		return c.JSON(http.StatusOK, Response{
			Token: string(claimToken),
		})
	})

	// Setup GraphQL Server Handler
	graphqlServerHandler, playgroundHandler := setupGraphqlHandler(db)

	e.POST("/graphql", func(c echo.Context) error {
		graphqlServerHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}
