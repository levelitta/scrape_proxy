package api

type CaptchaError struct{}

func NewCaptchaError() CaptchaError {
	return CaptchaError{}
}

func (e CaptchaError) Error() string {
	return "captcha error"
}
