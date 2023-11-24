package boot

import (
	"reflect"

	"github.com/backend-timedoor/gtimekeeper/base/contracts"
	"github.com/backend-timedoor/gtimekeeper/providers"
)

func Boot(pvds []contracts.ServiceProvider) {
	pvds = append(pvds, []contracts.ServiceProvider{
		&providers.CacheServiceProvider{},
		&providers.DatabaseServiceProvider{},
	}...)

	for _, provider := range pvds {
		provider.Boot()
		r := reflect.TypeOf(provider).Elem()
		if (r.Name() == "ConfigServiceProvider") {
			log := providers.LogServiceProvider{}
			log.Register()
			log.Boot()
		}

		provider.Register()
	}
}
