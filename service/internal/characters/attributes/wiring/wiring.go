package wiring

import (
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/dao"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes/http"
	"github.com/sazzer/wyrdwest/service/internal/database"
	"github.com/sazzer/wyrdwest/service/internal/server"
)

// AttributesWiring builds the components that we want to wire in for our application
func AttributesWiring(db database.DB) (dao.AttributesDao, server.HandlerRegistrationFunc) {
	dao := dao.New(db)

	registrationFunc := http.RegisterAttributes(dao)

	return dao, registrationFunc
}
