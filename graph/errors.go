package graph

import "errors"

var (
	ErrAccountUsed              = errors.New("account already in used")
	ErrNameNotLongEnough        = errors.New("name not long enough")
	ErrDescriptionNotLongEnough = errors.New("description not long enough")
	ErrWrongPassword            = errors.New("wrong password")
	ErrUnkown                   = errors.New("something went wrong")
	ErrBadCredentials           = errors.New("email/password combination don't work")
)
