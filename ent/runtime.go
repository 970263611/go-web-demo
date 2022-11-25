// Code generated by ent, DO NOT EDIT.

package ent

import (
	"project/ent/schema"
	"project/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserCode is the schema descriptor for userCode field.
	userDescUserCode := userFields[0].Descriptor()
	// user.UserCodeValidator is a validator for the "userCode" field. It is called by the builders before save.
	user.UserCodeValidator = func() func(string) error {
		validators := userDescUserCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(userCode string) error {
			for _, fn := range fns {
				if err := fn(userCode); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescDefaultPassword is the schema descriptor for defaultPassword field.
	userDescDefaultPassword := userFields[3].Descriptor()
	// user.DefaultPasswordValidator is a validator for the "defaultPassword" field. It is called by the builders before save.
	user.DefaultPasswordValidator = func() func(string) error {
		validators := userDescDefaultPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(defaultPassword string) error {
			for _, fn := range fns {
				if err := fn(defaultPassword); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreateTime is the schema descriptor for createTime field.
	userDescCreateTime := userFields[5].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the createTime field.
	user.DefaultCreateTime = userDescCreateTime.Default.(int64)
}