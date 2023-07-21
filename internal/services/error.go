package services

import "errors"

//TODO move this file somewhere else (?)

var ErrorUserNotFound = errors.New("User Not Found") //FIXME need to fix all errors in project, error string should not be capitalized
var ErrorProductNotFound = errors.New("Product Not Found")