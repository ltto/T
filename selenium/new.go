package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func NewChromeDriver() (d *WebDriver, err error) {
	var (
		service *selenium.Service
		port    int
	)
	if port, err = FindFreePort(); err != nil {
		return nil, err
	}
	if service, err = selenium.NewChromeDriverService(FindPATH("chromedriver"), port); nil != err {
		return nil, fmt.Errorf("start a chromedriver service falid %v", err)
	}
	if driver, err := selenium.NewRemote(map[string]interface{}{"browserName": "chrome"}, fmt.Sprintf("http://localhost:%d/wd/hub", port)); err != nil {
		service.Stop()
		return nil, fmt.Errorf("failed to open session: %v", err)
	} else {
		return &WebDriver{Data: driver, Ser: service}, nil
	}
}
