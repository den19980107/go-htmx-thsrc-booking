// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/ordervalidation"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/schema"
	"github.com/mikestefanello/pagoda/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescDepartureStation is the schema descriptor for departure_station field.
	orderDescDepartureStation := orderFields[2].Descriptor()
	// order.DepartureStationValidator is a validator for the "departure_station" field. It is called by the builders before save.
	order.DepartureStationValidator = orderDescDepartureStation.Validators[0].(func(string) error)
	// orderDescArrivalStation is the schema descriptor for arrival_station field.
	orderDescArrivalStation := orderFields[3].Descriptor()
	// order.ArrivalStationValidator is a validator for the "arrival_station" field. It is called by the builders before save.
	order.ArrivalStationValidator = orderDescArrivalStation.Validators[0].(func(string) error)
	// orderDescIDNumber is the schema descriptor for id_number field.
	orderDescIDNumber := orderFields[4].Descriptor()
	// order.IDNumberValidator is a validator for the "id_number" field. It is called by the builders before save.
	order.IDNumberValidator = orderDescIDNumber.Validators[0].(func(string) error)
	// orderDescPhoneNumber is the schema descriptor for phone_number field.
	orderDescPhoneNumber := orderFields[5].Descriptor()
	// order.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	order.PhoneNumberValidator = orderDescPhoneNumber.Validators[0].(func(string) error)
	// orderDescEmail is the schema descriptor for email field.
	orderDescEmail := orderFields[6].Descriptor()
	// order.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	order.EmailValidator = orderDescEmail.Validators[0].(func(string) error)
	// orderDescStatus is the schema descriptor for status field.
	orderDescStatus := orderFields[7].Descriptor()
	// order.DefaultStatus holds the default value on creation for the status field.
	order.DefaultStatus = orderDescStatus.Default.(string)
	// order.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	order.StatusValidator = orderDescStatus.Validators[0].(func(string) error)
	// orderDescErrorMessage is the schema descriptor for error_message field.
	orderDescErrorMessage := orderFields[8].Descriptor()
	// order.DefaultErrorMessage holds the default value on creation for the error_message field.
	order.DefaultErrorMessage = orderDescErrorMessage.Default.(string)
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderFields[9].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	ordervalidationFields := schema.OrderValidation{}.Fields()
	_ = ordervalidationFields
	// ordervalidationDescJessionID is the schema descriptor for jession_id field.
	ordervalidationDescJessionID := ordervalidationFields[0].Descriptor()
	// ordervalidation.JessionIDValidator is a validator for the "jession_id" field. It is called by the builders before save.
	ordervalidation.JessionIDValidator = ordervalidationDescJessionID.Validators[0].(func(string) error)
	// ordervalidationDescCaptchaImage is the schema descriptor for captcha_image field.
	ordervalidationDescCaptchaImage := ordervalidationFields[1].Descriptor()
	// ordervalidation.CaptchaImageValidator is a validator for the "captcha_image" field. It is called by the builders before save.
	ordervalidation.CaptchaImageValidator = ordervalidationDescCaptchaImage.Validators[0].(func(string) error)
	// ordervalidationDescCookies is the schema descriptor for cookies field.
	ordervalidationDescCookies := ordervalidationFields[2].Descriptor()
	// ordervalidation.CookiesValidator is a validator for the "cookies" field. It is called by the builders before save.
	ordervalidation.CookiesValidator = ordervalidationDescCookies.Validators[0].(func(string) error)
	// ordervalidationDescCreatedAt is the schema descriptor for created_at field.
	ordervalidationDescCreatedAt := ordervalidationFields[3].Descriptor()
	// ordervalidation.DefaultCreatedAt holds the default value on creation for the created_at field.
	ordervalidation.DefaultCreatedAt = ordervalidationDescCreatedAt.Default.(func() time.Time)
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[3].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)
