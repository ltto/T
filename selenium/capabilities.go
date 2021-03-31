package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/firefox"
	"github.com/tebeka/selenium/log"
)

type Capabilities struct {
	d *selenium.Capabilities
}

// AddChrome adds Chrome-specific capabilities.
func (c *Capabilities) AddChrome(f chrome.Capabilities) {
	(*c.d)[chrome.CapabilitiesKey] = f
	(*c.d)[chrome.DeprecatedCapabilitiesKey] = f
}

// AddFirefox adds Firefox-specific capabilities.
func (c *Capabilities) AddFirefox(f firefox.Capabilities) {
	(*c.d)[firefox.CapabilitiesKey] = f
}

// AddProxy adds proxy configuration to the capabilities.
func (c *Capabilities) AddProxy(p selenium.Proxy) {
	(*c.d)["proxy"] = p
}

// AddLogging adds logging configuration to the capabilities.
func (c *Capabilities) AddLogging(l log.Capabilities) {
	(*c.d)[log.CapabilitiesKey] = l
}

// SetLogLevel sets the logging level of a component. It is a shortcut for
// passing a log.Capabilities instance to AddLogging.
func (c *Capabilities) SetLogLevel(typ log.Type, level log.Level) {
	if _, ok := (*c.d)[log.CapabilitiesKey]; !ok {
		(*c.d)[log.CapabilitiesKey] = make(log.Capabilities)
	}
	m := (*c.d)[log.CapabilitiesKey].(log.Capabilities)
	m[typ] = level
}
