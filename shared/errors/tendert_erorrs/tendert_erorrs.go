package tendert_erorrs

import "errors"

var (
	ErrTenderNotFound        = errors.New("tender not found")
	ErrUnauthorized          = errors.New("user not authorized")
	ErrTenderVersionNotFound = errors.New("tender version not found")
	ErrInvalidStatus         = errors.New("not valid status")
)
