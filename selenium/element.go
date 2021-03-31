package selenium

import (
	"github.com/tebeka/selenium"
)

type WebElement struct {
	d selenium.WebElement
}

func (w WebElement) Click() {
	err := w.d.Click()
	if err != nil {
		panic(err)
	}
}

func (w WebElement) SendKeys(keys string) {
	err := w.d.SendKeys(keys)
	if err != nil {
		panic(err)
	}
}

func (w WebElement) Submit() {
	err := w.d.Submit()
	if err != nil {
		panic(err)
	}
}

func (w WebElement) Clear() {
	err := w.d.Clear()
	if err != nil {
		panic(err)
	}
}

func (w WebElement) MoveTo(xOffset, yOffset int) {
	err := w.d.MoveTo(xOffset, yOffset)
	if err != nil {
		panic(err)
	}
}

func (w WebElement) FindElement(by, value string) WebElement {
	element, err := w.d.FindElement(by, value)
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (w WebElement) FindElements(by, value string) (list []WebElement) {
	elements, err := w.d.FindElements(by, value)
	if err != nil {
		panic(err)
	}
	for i := range elements {
		list = append(list, WebElement{elements[i]})
	}
	return
}

func (w WebElement) TagName() string {
	name, err := w.d.TagName()
	if err != nil {
		panic(err)
	}
	return name
}

func (w WebElement) Text() string {
	text, err := w.d.Text()
	if err != nil {
		panic(err)
	}
	return text
}

func (w WebElement) IsSelected() bool {
	enabled, err := w.d.IsSelected()
	if err != nil {
		panic(err)
	}
	return enabled
}

func (w WebElement) IsEnabled() bool {
	enabled, err := w.d.IsEnabled()
	if err != nil {
		panic(err)
	}
	return enabled
}

func (w WebElement) IsDisplayed() bool {
	displayed, err := w.d.IsDisplayed()
	if err != nil {
		panic(err)
	}
	return displayed
}

func (w WebElement) GetAttribute(name string) string {
	attribute, err := w.d.GetAttribute(name)
	if err != nil {
		panic(err)
	}
	return attribute
}

func (w WebElement) Location() *selenium.Point {
	location, err := w.d.Location()
	if err != nil {
		panic(err)
	}
	return location
}

func (w WebElement) LocationInView() *selenium.Point {
	view, err := w.d.LocationInView()
	if err != nil {
		panic(err)
	}
	return view
}

func (w WebElement) Size() *selenium.Size {
	size, err := w.d.Size()
	if err != nil {
		panic(err)
	}
	return size
}

func (w WebElement) CSSProperty(name string) string {
	property, err := w.d.CSSProperty(name)
	if err != nil {
		panic(err)
	}
	return property
}

func (w WebElement) Screenshot(scroll bool) []byte {
	screenshot, err := w.d.Screenshot(scroll)
	if err != nil {
		panic(err)
	}
	return screenshot
}
