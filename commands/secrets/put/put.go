package main

import (
	"flag"
	"io/ioutil"

	"github.com/walterjwhite/go-application/libraries/application"
	"github.com/walterjwhite/go-application/libraries/logging"
	"github.com/walterjwhite/go-application/libraries/secrets"
)

type NoSecretDataError struct{}
type NoCommitMessageError struct{}
type NoNameError struct{}

func (e NoSecretDataError) Error() string {
	return "No secret data was provided."
}

func (e NoCommitMessageError) Error() string {
	return "No commit message was provided."
}

func (e NoNameError) Error() string {
	return "No name was provided."
}

var name = flag.String("name", "", "Secret key name (hierarchy to key, excluding trailing /value, ie. /email/gmail.com/personal/email-address)")
var message = flag.String("message", "", "Commit message")
var source = flag.String("source", "", "source file")

// TODO: add support for flags
// instead of specifying the key type (email, user, pass), use a flag instead (-e, -u, -p)
func main() {
	_ = application.Configure()

	validatePut(name, message, source)

	data, err := ioutil.ReadFile(*source)
	logging.Panic(err)

	secrets.Encrypt(name, message, data)
}

func validatePut(name *string, message *string, source *string) {
	if len(*name) == 0 {
		logging.Panic(&NoNameError{})
	}

	if len(*message) == 0 {
		logging.Panic(&NoCommitMessageError{})
	}

	if len(*source) == 0 {
		logging.Panic(&NoSecretDataError{})
	}
}