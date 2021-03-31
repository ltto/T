package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
	"time"
)

type WebDriver struct {
	Ser  *selenium.Service
	Data selenium.WebDriver
}

func (d *WebDriver) Status() *selenium.Status {
	status, err := d.Data.Status()
	if err != nil {
		panic(err)
	}
	return status
}

func (d *WebDriver) NewSession() string {
	session, err := d.Data.NewSession()
	if err != nil {
		panic(err)
	}
	return session
}

func (d *WebDriver) SessionId() string {
	return d.Data.SessionId()
}

func (d *WebDriver) SessionID() string {
	return d.Data.SessionID()
}

func (d *WebDriver) SwitchSession(sessionID string) {
	if err := d.Data.SwitchSession(sessionID); err != nil {
		panic(err)
	}
}

func (d *WebDriver) Capabilities() Capabilities {
	capabilities, err := d.Data.Capabilities()
	if err != nil {
		panic(err)
	}
	return Capabilities{d: &capabilities}
}

func (d *WebDriver) SetAsyncScriptTimeout(timeout time.Duration) {
	err := d.Data.SetAsyncScriptTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) SetImplicitWaitTimeout(timeout time.Duration) {
	err := d.Data.SetImplicitWaitTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) SetPageLoadTimeout(timeout time.Duration) {
	err := d.Data.SetPageLoadTimeout(timeout)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Quit() {
	err := d.Data.Quit()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) CurrentWindowHandle() string {
	url, err := d.Data.CurrentURL()
	if err != nil {
		panic(err)
	}
	return url
}

func (d *WebDriver) WindowHandles() []string {
	handles, err := d.Data.WindowHandles()
	if err != nil {
		panic(err)
	}
	return handles
}

func (d *WebDriver) CurrentURL() string {
	url, err := d.Data.CurrentURL()
	if err != nil {
		panic(err)
	}
	return url
}

func (d *WebDriver) Title() string {
	title, err := d.Data.Title()
	if err != nil {
		panic(err)
	}
	return title
}

func (d *WebDriver) PageSource() string {
	source, err := d.Data.PageSource()
	if err != nil {
		panic(err)
	}
	return source
}

func (d *WebDriver) Close() {
	err := d.Data.Close()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) SwitchFrame(frame interface{}) {
	if element, ok := frame.(WebElement); ok {
		frame = element.d
	}
	err := d.Data.SwitchFrame(frame)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) SwitchWindow(name string) {
	err := d.Data.SwitchWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) CloseWindow(name string) {
	err := d.Data.CloseWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) MaximizeWindow(name string) {
	err := d.Data.MaximizeWindow(name)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) ResizeWindow(name string, width, height int) {
	err := d.Data.ResizeWindow(name, width, height)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Get(url string) {
	err := d.Data.Get(url)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Forward() {
	err := d.Data.Forward()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Back() {
	err := d.Data.Back()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Refresh() {
	err := d.Data.Refresh()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) FindElement(by, value string) WebElement {
	element, err := d.Data.FindElement(by, value)
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *WebDriver) FindElements(by, value string) (list []WebElement) {
	elements, err := d.Data.FindElements(by, value)
	if err != nil {
		panic(err)
	}
	for i := range elements {
		list = append(list, WebElement{elements[i]})
	}
	return list
}

func (d *WebDriver) ActiveElement() WebElement {
	element, err := d.Data.ActiveElement()
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *WebDriver) DecodeElement(bytes []byte) WebElement {
	element, err := d.Data.DecodeElement(bytes)
	if err != nil {
		panic(err)
	}
	return WebElement{element}
}

func (d *WebDriver) DecodeElements(bytes []byte) (list []WebElement) {
	elements, err := d.Data.DecodeElements(bytes)
	if err != nil {
		panic(err)
	}
	for i := range elements {
		list = append(list, WebElement{elements[i]})
	}

	return list
}

func (d *WebDriver) GetCookies() []selenium.Cookie {
	cookies, err := d.Data.GetCookies()
	if err != nil {
		panic(err)
	}
	return cookies
}

func (d *WebDriver) GetCookie(name string) (selenium.Cookie, ) {
	cookie, err := d.Data.GetCookie(name)
	if err != nil {
		panic(err)
	}
	return cookie
}

func (d *WebDriver) AddCookie(cookie *selenium.Cookie) {
	err := d.Data.AddCookie(cookie)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) DeleteAllCookies() {
	err := d.Data.DeleteAllCookies()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) DeleteCookie(name string) {
	err := d.Data.DeleteCookie(name)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Click(button int) {
	err := d.Data.Click(button)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) DoubleClick() {
	err := d.Data.DoubleClick()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) ButtonDown() {
	err := d.Data.ButtonDown()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) ButtonUp() {
	err := d.Data.ButtonUp()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) SendModifier(modifier string, isDown bool) {
	err := d.Data.SendModifier(modifier, isDown)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) KeyDown(keys string) {
	err := d.Data.KeyDown(keys)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) KeyUp(keys string) {
	err := d.Data.KeyUp(keys)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Screenshot() []byte {
	screenshot, err := d.Data.Screenshot()
	if err != nil {
		panic(err)
	}
	return screenshot
}

func (d *WebDriver) Log(typ log.Type) []log.Message {
	messages, err := d.Data.Log(typ)
	if err != nil {
		panic(err)
	}
	return messages
}

func (d *WebDriver) DismissAlert() {
	err := d.Data.DismissAlert()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) AcceptAlert() {
	err := d.Data.AcceptAlert()
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) AlertText() string {
	text, err := d.Data.AlertText()
	if err != nil {
		panic(err)
	}
	return text
}

func (d *WebDriver) SetAlertText(text string) {
	err := d.Data.SetAlertText(text)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) ExecuteScript(script string, args []interface{}) interface{} {
	executeScript, err := d.Data.ExecuteScript(script, args)
	if err != nil {
		panic(err)
	}
	return executeScript
}

func (d *WebDriver) ExecuteScriptAsync(script string, args []interface{}) interface{} {
	async, err := d.Data.ExecuteScriptAsync(script, args)
	if err != nil {
		panic(err)
	}
	return async
}

func (d *WebDriver) ExecuteScriptRaw(script string, args []interface{}) []byte {
	raw, err := d.Data.ExecuteScriptRaw(script, args)
	if err != nil {
		panic(err)
	}
	return raw
}

func (d *WebDriver) ExecuteScriptAsyncRaw(script string, args []interface{}) []byte {
	raw, err := d.Data.ExecuteScriptAsyncRaw(script, args)
	if err != nil {
		panic(err)
	}
	return raw
}

func (d *WebDriver) WaitWithTimeoutAndInterval(condition selenium.Condition, timeout, interval time.Duration) {
	err := d.Data.WaitWithTimeoutAndInterval(condition, timeout, interval)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) {
	err := d.Data.WaitWithTimeout(condition, timeout)
	if err != nil {
		panic(err)
	}
}

func (d *WebDriver) Wait(condition selenium.Condition) {
	err := d.Data.Wait(condition)
	if err != nil {
		panic(err)
	}
}
