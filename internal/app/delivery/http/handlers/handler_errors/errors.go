package handler_errors

import "errors"

var (
	UserAlreadyExist           = errors.New("user already exist")
	NicknameAlreadyExist       = errors.New("nickname already exist")
	IncorrectEmailOrPassword   = errors.New("incorrect email or password")
	UserNotFound               = errors.New("user with this id not found")
	GetProfileFail             = errors.New("can not get user from db")
	DeleteCookieFail           = errors.New("can not delete cookie from session store")
	InvalidBody                = errors.New("invalid body in request")
	ErrorCreateUser            = errors.New("can not create user")
	ErrorPrepareUser           = errors.New("can not prepare user info")
	ContextError               = errors.New("can not get info from context")
	ErrorCreateSession         = errors.New("can not create session")
	BDError                    = errors.New("can not do bd operation")
	InvalidParameters          = errors.New("invalid parameters")
	NotAllowedMethod           = errors.New("method not allowed")
	ProfileAlreadyExist        = errors.New("user already exist")
	InvalidNickname            = errors.New("invalid creator nickname")
	InvalidCategory            = errors.New("invalid creator category")
	InvalidCategoryDescription = errors.New("invalid creator category-description")
	InternalError              = errors.New("server error")
)
