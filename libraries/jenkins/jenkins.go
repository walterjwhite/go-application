package jenkins

import (
	"github.com/walterjwhite/go-application/libraries/logging"

	"gopkg.in/bndr/gojenkins.v1"
	//"github.com/pushyzheng/gojenkins"
	"time"
)

/*
type JenkinsCredentials struct {
	Username string
	Password string
}
*/

func (c *JenkinsInstance) EncryptedFields() []string {
	//return []string{"JenkinsCredentials.Username", "JenkinsCredentials.Password"}
	return []string{"Username", "Password"}
}

type JenkinsInstance struct {
	Url string
	//JenkinsCredentials *JenkinsCredentials
	Username string
	Password string

	buildTimeout       time.Duration
	buildCheckInterval time.Duration

	jenkins *gojenkins.Jenkins
}

func (i *JenkinsInstance) setup() {
	if i.jenkins != nil {
		return
	}

	i.jenkins = gojenkins.CreateJenkins(nil, i.Url, i.Username, i.Password)

	_, err := i.jenkins.Init()
	logging.Panic(err)
}
