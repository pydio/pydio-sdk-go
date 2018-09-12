package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	sampleFilePath = "config.sample.json"
)

func TestSetUpProcess(t *testing.T) {

	Convey("Sample files should be found", t, func() {
		_, e := ioutil.ReadFile(sampleFilePath)
		So(e, ShouldBeNil)
	})

	// We run subsequent tests only in basic builds:
	// It breaks other tests when in EnvAware mode because initialisation of the env is kept by convey
	if RunEnvAwareTests {
		return
	}

	Convey("Environment Config must be loaded", t, func() {
		s, _ := ioutil.ReadFile(sampleFilePath)
		var c SdkConfig
		e := json.Unmarshal(s, &c)
		So(e, ShouldBeNil)
		So(c.User, ShouldEqual, "admin")
	})

	Convey("SDK should be configured", t, func() {
		s, _ := ioutil.ReadFile(sampleFilePath)
		var c SdkConfig
		json.Unmarshal(s, &c)
		So(c.User, ShouldEqual, "admin")

		DefaultConfig = &c
		So(DefaultConfig.User, ShouldEqual, "admin")
	})

}

func TestEnvConfiguration(t *testing.T) {

	Convey("Environment must be configured", t, func() {

		Convey("Pydio Cells SDK must be configured", func() {
			var c SdkConfig
			// check presence of Env variable
			protocol := os.Getenv(KeyProtocol)
			url := os.Getenv(KeyURL)
			user := os.Getenv(KeyUser)
			password := os.Getenv(KeyPassword)
			skipVerifyStr := os.Getenv(KeySkipVerify)
			if skipVerifyStr == "" {
				skipVerifyStr = "false"
			}
			skipVerify, err := strconv.ParseBool(skipVerifyStr)
			So(err, ShouldBeNil)
			if len(protocol) > 0 && len(url) > 0 && len(user) > 0 && len(password) > 0 {
				c.Protocol = protocol
				c.Url = url
				c.User = user
				c.Password = password
				c.SkipVerify = skipVerify
			} else { // fallback to config.json file
				currPath, err := filepath.Abs(".")
				So(err, ShouldBeNil)
				configFilePath := filepath.Join(currPath, sampleFilePath)
				s, err := ioutil.ReadFile(configFilePath)
				e := json.Unmarshal(s, &c)
				So(e, ShouldBeNil)
			}

			if RunEnvAwareTests {
				So(c.User, ShouldEqual, "admin")
			} else {

				DefaultConfig = &c
				So(DefaultConfig.User, ShouldEqual, "admin")
				// err := SetUpEnvironment()
				// Convey("Setup must be complete", t, func() {
				// 	So(err, ShouldBeNil)
				// 	So(DefaultConfig.User, ShouldEqual, "admin")
				// 	So(DefaultS3Config.Bucket, ShouldEqual, "io")
				// })
			}
		})
	})
}
