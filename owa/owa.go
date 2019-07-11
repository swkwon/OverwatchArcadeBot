package owa

// WebHook ...
type WebHook interface {
	Send(text string) error
}

var webHookers []WebHook
func init() {
	webHookers = append(webHookers, &Discord{})
}

func Send(text string) error {
	for _, v := range webHookers {
		if e := v.Send(text); e != nil {
			return e
		}
	}
	return nil
}