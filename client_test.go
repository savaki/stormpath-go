package stormpath

import (
	"github.com/savaki/stormpath-go/auth"
	. "github.com/savaki/stormpath-go/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestClient(t *testing.T) {
	if _, err := auth.EnvAuth(); err != nil {
		return
	}

	var client *Client
	var apiKey auth.ApiKey
	var err error
	var tenant *Tenant
	var applications []Application
	var directories []Map

	Convey("Given a StormPath client connected to the default tenant", t, func() {
		apiKey, _ = auth.EnvAuth()
		client, err = WithCurrentTenant(apiKey)

		So(err, ShouldBeNil)
		So(client.tenant, ShouldNotBeNil)

		Convey("When I find #CurrentTenant", func() {
			tenant, err = client.CurrentTenant()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect tenant to be assigned", func() {
				So(tenant, ShouldNotBeNil)
				So(tenant.Href, ShouldNotEqual, "")
				So(tenant.Name, ShouldNotEqual, "")
				So(tenant.Key, ShouldNotEqual, "")
				So(tenant.CustomData.Href, ShouldNotEqual, "")
				So(tenant.Applications.Href, ShouldNotEqual, "")
				So(tenant.Directories.Href, ShouldNotEqual, "")
			})
		})

		Convey("When I find #Applications", func() {
			applications, err = client.Applications()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect applications to be returned", func() {
				So(applications, ShouldNotBeNil)
				So(applications[0].Accounts, ShouldNotBeNil)
				So(applications[0].AccountStoreMappings, ShouldBeNil)
				So(applications[0].ApiKeys, ShouldNotBeNil)
				So(applications[0].CustomData, ShouldNotBeNil)
				So(applications[0].DefaultAccountStoreMapping, ShouldBeNil)
				So(applications[0].DefaultGroupStoreMapping, ShouldBeNil)
				So(applications[0].Description, ShouldNotBeEmpty)
				So(applications[0].Groups, ShouldNotBeNil)
				So(applications[0].Href, ShouldNotBeEmpty)
				So(applications[0].LoginAttempts, ShouldNotBeNil)
				So(applications[0].Name, ShouldNotBeEmpty)
				So(applications[0].PasswordResetTokens, ShouldNotBeNil)
				So(applications[0].Status, ShouldNotBeEmpty)
				So(applications[0].Tenant, ShouldNotBeNil)
			})
		})

		Convey("When I find #Directories", func() {
			directories, err = client.Directories()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect applications to be returned", func() {
				So(directories, ShouldNotBeNil)
				So(len(directories), ShouldBeGreaterThan, 0)
				for _, directory := range directories {
					So(directory.Href(), ShouldNotEqual, "")
				}
			})
		})
	})
}
