package utils

import "errors"

// Controller errors
var (
	// ErrBadRequestBody is used when the request body is not valid format
	ErrBadRequestBody = errors.New("bad request body")

	// ErrInvalidCredentials is used when the user's credentials are invalid
	ErrInvalidCredentials = errors.New("invalid username or password")

	// ErrDidntHavePermission is used when the user doesn't have permission to access or modify the resource
	ErrDidntHavePermission = errors.New("you didn't have permission to do this action")

	// ErrInvalidNumber is used when the number covertion is invalid
	ErrInvalidNumber = errors.New("invalid number")
)

// Service errors
var (
	// ErrFieldNotMatch is used when the field in the request body is not match with the field in the template that saved in the database
	ErrFieldNotMatch = errors.New("document fields doesn't match with template fields")

	// ErrAlreadyVerified is used when the document is already verified
	ErrAlreadyVerified = errors.New("already verified")

	// ErrNotVerifiedYet is used when the document is not verified yet
	ErrNotVerifiedYet = errors.New("not verified yet")

	// ErrAlreadySigned is used when the document is already signed
	ErrAlreadySigned = errors.New("already signed")
)

// Repository errors
var (
	// ErrUsernameAlreadyExist is used when the username is already exist in the database
	ErrUsernameAlreadyExist = errors.New("user with provided username already exist")

	// ErrUserNotFound is used when the user is not found in the database
	ErrUserNotFound = errors.New("user not found")

	// ErrItemAlreadyExist is used when the item name is already exist in the database
	ErrEncryptionAlreadyExist = errors.New("item with provided name already exist")

	// ErrSupplierAlreadyExist is used when the supplier name is already exist in the database
	ErrDecryptionAlreadyExist = errors.New("supplier with provided name already exist")

	// ErrSupplierAlreadyExist is used when the supplier name is already exist in the database
	ErrKeyAlreadyExist = errors.New("supplier with provided name already exist")

	// ErrTemplateFieldNotFound is used when the template field is not found in the database
	ErrTemplateFieldNotFound = errors.New("template field not found")

	// ErrDuplicateRegister is used when the document register is already exist in the database
	ErrDuplicateRegister = errors.New("document with provided register already exist")

	// ErrDocumentNotFound is used when the document is not found in the database
	ErrDocumentNotFound = errors.New("document not found")

	// ErrItemNotFound is used when the item is not found in the database
	ErrEncryptionNotFound = errors.New("item not found")

	// ErrSupplierNotFound is used when the supplier is not found in the database
	ErrDecryptionNotFound = errors.New("supplier not found")

	// ErrPurchaseNotFound is used when the purchase is not found in the database
	ErrKeyNotFound = errors.New("purchase not found")

	// ErrSaleNotFound is used when the purchase is not found in the database
	ErrSaleNotFound = errors.New("sale not found")
)
