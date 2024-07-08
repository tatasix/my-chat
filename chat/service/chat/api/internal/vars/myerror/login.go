package myerror

import "errors"

var (
	LoginMobileError = errors.New(" mobile error")

	LoginVerifyCodeError = errors.New(" verify code error")

	LoginAccountNotExistError = errors.New(" account is not exist")
)
