package tenants

import (
	"github.com/savaki/stormpath-go/stormpath"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCurrentTenant(t *testing.T) {
	if _, err := stormpath.EnvAuth(); err != nil {
		return
	}

	var tenants *Tenants
	var apiKey stormpath.ApiKey
	var tenant *Tenant
	var err error

	Convey("Given a Tenants instance", t, func() {
		apiKey, _ = stormpath.EnvAuth()
		tenants = New(apiKey)

		Convey("When I find #CurrentTenant", func() {
			tenant, err = tenants.CurrentTenant()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect tenant to be assigned", func() {
				So(tenant, ShouldNotBeNil)
			})
		})
	})
}
