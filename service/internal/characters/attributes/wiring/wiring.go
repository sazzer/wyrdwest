package wiring

import (
	"github.com/go-chi/chi"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/dao"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/http"
	"github.com/sazzer/wyrdwest/service/internal/database"
)

// AttributesWiring builds the components that we want to wire in for our application
func AttributesWiring(db database.DB) (dao.AttributesDao, *chi.Mux) {
	dao := dao.New(db)

	router := http.NewRouter(dao)

	return dao, router
}
