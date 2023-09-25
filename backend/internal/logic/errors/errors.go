package errors

import "errors"

var (
	ErrExist         = errors.New("does not exist")
	ErrFilter        = errors.New("invalid filter")
	ErrGet           = errors.New("empty context")
	ErrAdmin         = errors.New("admin")
	ErrNotAdmin      = errors.New("not an admin")
	ErrBadLogin      = errors.New("bad login")
	ErrExistingLogin = errors.New("existing login")
	ErrBadPassword   = errors.New("bad password")
	ErrUserNotFound  = errors.New("user not found")
	ErrNotFound      = errors.New("data not found")
	ErrAutoDelete    = errors.New("cannot delete yourself")
)
