package main

import (
	"context"
	"database/sql"
	"github.com/DaoVuDat/graphql/sqlStore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strings"
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
	// Create DB
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/login", func(c echo.Context) error {
		var loginReq LoginReq

		if err := c.Bind(&loginReq); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		// Create Store
		store := sqlStore.NewStore(db)

		user, err := store.FindUserByEmail(loginReq.Email)
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
	},
		sqlStore.Middleware(db),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// Get header
				authorizationStr := c.Request().Header.Get("Authorization")
				chunks := strings.Split(authorizationStr, " ")
				if len(chunks) < 2 {
					return c.String(http.StatusUnauthorized, "unauthorized")
				}

				if strings.ToLower(chunks[0]) != "bearer" {
					return c.String(http.StatusBadRequest, "bad request")
				}

				token := chunks[1]
				parseToken, err := ParseToken(token)
				if err != nil {
					c.Error(err)
				}

				// Create Store
				store := sqlStore.NewStore(db)
				user, err := store.FindUserById(parseToken.Subject())
				if err != nil {
					c.Error(err)
				}

				// Copy ctx from parent context + additional key value
				ctx := context.WithValue(c.Request().Context(), "user", user)

				// set new request with new ctx from above
				c.SetRequest(c.Request().WithContext(ctx))

				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}
