package accounts

import (
	"fmt"
	"github.com/savaki/stormpath-go/auth"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestAccounts(t *testing.T) {
	if _, err := auth.EnvAuth(); err != nil {
		return
	}
	if os.Getenv("STORMPATH_URL") == "" {
		return
	}

	var apiKey auth.ApiKey
	var accounts *Accounts
	var err error
	var account *Account
	var saved *Account
	var email string

	Convey("Given a StormPath client", t, func() {
		apiKey, _ = auth.EnvAuth()
		accounts = FromUrl(apiKey, os.Getenv("STORMPATH_URL"))

		So(err, ShouldBeNil)
		So(accounts, ShouldNotBeNil)

		Convey("When I create an account", func() {
			email = fmt.Sprintf("joe.public-%d@gmail.com", time.Now().Unix())
			account = &Account{
				Username:  email,
				Email:     email,
				GivenName: "Joe",
				Surname:   "Public",
				Password:  "Pass123!",
			}
			saved, err = accounts.Create(*account)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect a valid saved account back", func() {
				So(saved, ShouldNotBeNil)
				So(saved.Href, ShouldNotEqual, "")
			})
		})

		Reset(func() {
			if accounts != nil && saved != nil {
				err = accounts.Delete(saved)
				if err != nil {
					fmt.Println(err)
				}
			}
		})
	})
}
