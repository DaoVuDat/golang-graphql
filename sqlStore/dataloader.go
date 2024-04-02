package sqlStore

import (
	"context"
	"database/sql"
	"github.com/DaoVuDat/graphql/.gen/model"
	"github.com/labstack/echo/v4"
	"github.com/vikstrous/dataloadgen"
	"time"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	CompanyLoader *dataloadgen.Loader[string, *model.Company]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(db *sql.DB) *Loaders {
	// define the data loader
	storeReader := NewStore(db)
	return &Loaders{
		CompanyLoader: dataloadgen.NewLoader(storeReader.FindCompaniesByIds, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(db *sql.DB) func(next echo.HandlerFunc) echo.HandlerFunc {
	// return a middleware that injects the loader to the request context
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			loader := NewLoaders(db)
			ctx := context.WithValue(c.Request().Context(), loadersKey, loader)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
