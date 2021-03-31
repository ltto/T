package selenium

import (
	"github.com/tebeka/selenium"
)

// Methods by which to find elements.
const (
	ByID              = "id"
	ByXPATH           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name"
	ByTagName         = "tag name"
	ByClassName       = "class name"
	ByCSSSelector     = "css selector"
)

// Mouse buttons.
const (
	LeftButton = iota
	MiddleButton
	RightButton
)

// Special keyboard keys, for SendKeys.
const (
	NullKey       = string('\ue000')
	CancelKey     = string('\ue001')
	HelpKey       = string('\ue002')
	BackspaceKey  = string('\ue003')
	TabKey        = string('\ue004')
	ClearKey      = string('\ue005')
	ReturnKey     = string('\ue006')
	EnterKey      = string('\ue007')
	ShiftKey      = string('\ue008')
	ControlKey    = string('\ue009')
	AltKey        = string('\ue00a')
	PauseKey      = string('\ue00b')
	EscapeKey     = string('\ue00c')
	SpaceKey      = string('\ue00d')
	PageUpKey     = string('\ue00e')
	PageDownKey   = string('\ue00f')
	EndKey        = string('\ue010')
	HomeKey       = string('\ue011')
	LeftArrowKey  = string('\ue012')
	UpArrowKey    = string('\ue013')
	RightArrowKey = string('\ue014')
	DownArrowKey  = string('\ue015')
	InsertKey     = string('\ue016')
	DeleteKey     = string('\ue017')
	SemicolonKey  = string('\ue018')
	EqualsKey     = string('\ue019')
	Numpad0Key    = string('\ue01a')
	Numpad1Key    = string('\ue01b')
	Numpad2Key    = string('\ue01c')
	Numpad3Key    = string('\ue01d')
	Numpad4Key    = string('\ue01e')
	Numpad5Key    = string('\ue01f')
	Numpad6Key    = string('\ue020')
	Numpad7Key    = string('\ue021')
	Numpad8Key    = string('\ue022')
	Numpad9Key    = string('\ue023')
	MultiplyKey   = string('\ue024')
	AddKey        = string('\ue025')
	SeparatorKey  = string('\ue026')
	SubstractKey  = string('\ue027')
	DecimalKey    = string('\ue028')
	DivideKey     = string('\ue029')
	F1Key         = string('\ue031')
	F2Key         = string('\ue032')
	F3Key         = string('\ue033')
	F4Key         = string('\ue034')
	F5Key         = string('\ue035')
	F6Key         = string('\ue036')
	F7Key         = string('\ue037')
	F8Key         = string('\ue038')
	F9Key         = string('\ue039')
	F10Key        = string('\ue03a')
	F11Key        = string('\ue03b')
	F12Key        = string('\ue03c')
	MetaKey       = string('\ue03d')
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
