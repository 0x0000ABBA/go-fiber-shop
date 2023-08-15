package utils

import "errors"

//TODO:move this file somewhere else (?)

var ErrorUserNotFound = errors.New("user not found") //FIXME:need to fix all errors in project, error string should not be capitalized
var ErrorProductNotFound = errors.New("product not found")

var ErrorUserAlreadyExists = errors.New("user already exists")
var ErrorInvalidPassword = errors.New("invalid password")
var ErrorInvalidEmail = errors.New("invalid email")
