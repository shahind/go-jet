package module

import (
	"github.com/RobyFerro/go-web/service"
	"github.com/shahind/go-jet-framework/register"
)

var MainModule = register.DIModule{
	Name: "auth",
	Provides: []interface{}{
		service.ConnectDB,
		service.ConnectRedis,
		service.ConnectElastic,
		service.ConnectMongo,
	},
}
