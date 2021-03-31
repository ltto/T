package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
	"time"
)

type DriverDecorator struct {
	Data selenium.WebDriver
}

func (d *DriverDecorator) Status() *selenium.Status {
	status, err := d.Data.Status()
	if err != nil {
		panic(err)
	}
	return status
}

func (d *DriverDecorator) NewSession() string {
	session, err := d.Data.NewSession()
	if err != nil {
		panic(err)
	}
	return session
}

func (d *DriverDecorator) SessionId() string {
	return d.Data.SessionId()
}

func (d *DriverDecorator) SessionID() string {
	return d.Data.SessionID()
}

func (d *DriverDecorator) SwitchSession(sessionID string) {
	if err := d.Data.SwitchSession(sessionID); err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Capabilities() Capabilities {
	capabilities, err := d.Data.Capabilities()
	if err != nil {
		panic(err)
	}
	return Capabilities{d: &capabilities}
}

func (d *DriverDecorator) SetAsyncScriptTimeout(timeout time.Duration) {
	err := d.Data.SetAsyncScriptTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) SetImplicitWaitTimeout(timeout time.Duration) {
	err := d.Data.SetImplicitWaitTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) SetPageLoadTimeout(timeout time.Duration) {
	err := d.Data.SetPageLoadTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Quit() {
	err := d.Data.Quit()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) CurrentWindowHandle() string {
	url, err := d.Data.CurrentURL()
	if err != nil {
		panic(err)
	}
	return url
}

func (d *DriverDecorator) WindowHandles() []string {
	handles, err := d.Data.WindowHandles()
	if err != nil {
		panic(err)
	}
	return handles
}

func (d *DriverDecorator) CurrentURL() string {
	url, err := d.Data.CurrentURL()
	if err != nil {
		panic(err)
	}
	return url
}

func (d *DriverDecorator) Title() string {
	title, err := d.Data.Title()
	if err != nil {
		panic(err)
	}
	return title
}

func (d *DriverDecorator) PageSource() string {
	source, err := d.Data.PageSource()
	if err != nil {
		panic(err)
	}
	return source
}

func (d *DriverDecorator) Close() {
	err := d.Data.Close()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) SwitchFrame(frame interface{}) {
	err := d.Data.SwitchFrame(frame)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) SwitchWindow(name string) {
	err := d.Data.SwitchWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) CloseWindow(name string) {
	err := d.Data.CloseWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) MaximizeWindow(name string) {
	err := d.Data.MaximizeWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) ResizeWindow(name string, width, height int) {
	err := d.Data.ResizeWindow(name, width, height)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Get(url string) {
	err := d.Data.Get(url)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Forward() {
	err := d.Data.Forward()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Back() {
	err := d.Data.Back()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Refresh() {
	err := d.Data.Refresh()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) FindElement(by, value string) WebElement {
	element, err := d.Data.FindElement(by, value)
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *DriverDecorator) FindElements(by, value string) (list []WebElement) {
	elements, err := d.Data.FindElements(by, value)
	if err != nil {
		panic(err)
	}
	for i := range elements {
		list = append(list, WebElement{elements[i]})
	}
	return list
}

func (d *DriverDecorator) ActiveElement() WebElement {
	element, err := d.Data.ActiveElement()
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *DriverDecorator) DecodeElement(bytes []byte) WebElement {
	element, err := d.Data.DecodeElement(bytes)
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *DriverDecorator) DecodeElements(bytes []byte) (list []WebElement) {
	elements, err := d.Data.DecodeElements(bytes)
	if err != nil {
		panic(err)
	}
	for i := range elements {
		list = append(list, WebElement{elements[i]})
	}

	return list
}

func (d *DriverDecorator) GetCookies() []selenium.Cookie {
	cookies, err := d.Data.GetCookies()
	if err != nil {
		panic(err)
	}
	return cookies
}

func (d *DriverDecorator) GetCookie(name string) (selenium.Cookie, ) {
	cookie, err := d.Data.GetCookie(name)
	if err != nil {
		panic(err)
	}
	return cookie
}

func (d *DriverDecorator) AddCookie(cookie *selenium.Cookie) {
	err := d.Data.AddCookie(cookie)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) DeleteAllCookies() {
	err := d.Data.DeleteAllCookies()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) DeleteCookie(name string) {
	err := d.Data.DeleteCookie(name)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Click(button int) {
	err := d.Data.Click(button)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) DoubleClick() {
	err := d.Data.DoubleClick()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) ButtonDown() {
	err := d.Data.ButtonDown()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) ButtonUp() {
	err := d.Data.ButtonUp()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) SendModifier(modifier string, isDown bool) {
	err := d.Data.SendModifier(modifier, isDown)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) KeyDown(keys string) {
	err := d.Data.KeyDown(keys)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) KeyUp(keys string) {
	err := d.Data.KeyUp(keys)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Screenshot() []byte {
	screenshot, err := d.Data.Screenshot()
	if err != nil {
		panic(err)
	}
	return screenshot
}

func (d *DriverDecorator) Log(typ log.Type) []log.Message {
	messages, err := d.Data.Log(typ)
	if err != nil {
		panic(err)
	}
	return messages
}

func (d *DriverDecorator) DismissAlert() {
	err := d.Data.DismissAlert()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) AcceptAlert() {
	err := d.Data.AcceptAlert()
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) AlertText() string {
	text, err := d.Data.AlertText()
	if err != nil {
		panic(err)
	}
	return text
}

func (d *DriverDecorator) SetAlertText(text string) {
	err := d.Data.SetAlertText(text)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) ExecuteScript(script string, args []interface{}) interface{} {
	executeScript, err := d.Data.ExecuteScript(script, args)
	if err != nil {
		panic(err)
	}
	return executeScript
}

func (d *DriverDecorator) ExecuteScriptAsync(script string, args []interface{}) interface{} {
	async, err := d.Data.ExecuteScriptAsync(script, args)
	if err != nil {
		panic(err)
	}
	return async
}

func (d *DriverDecorator) ExecuteScriptRaw(script string, args []interface{}) []byte {
	raw, err := d.Data.ExecuteScriptRaw(script, args)
	if err != nil {
		panic(err)
	}
	return raw
}

func (d *DriverDecorator) ExecuteScriptAsyncRaw(script string, args []interface{}) ([]byte, ) {
	raw, err := d.Data.ExecuteScriptAsyncRaw(script, args)
	if err != nil {
		panic(err)
	}
	return raw
}

func (d *DriverDecorator) WaitWithTimeoutAndInterval(condition selenium.Condition, timeout, interval time.Duration) {
	err := d.Data.WaitWithTimeoutAndInterval(condition, timeout, interval)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) {
	err := d.Data.WaitWithTimeout(condition, timeout)
	if err != nil {
		panic(err)
	}
}

func (d *DriverDecorator) Wait(condition selenium.Condition) {
	err := d.Data.Wait(condition)
	if err != nil {
		panic(err)
	}
}
