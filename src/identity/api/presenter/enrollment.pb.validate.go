// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: identity/api/presenter/enrollment.proto

package presenter

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// define the regex for a UUID once up-front
var _enrollment_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RegistrationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegistrationRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Firstname

	// no validation rules for Lastname

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return RegistrationRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	return nil
}

func (m *RegistrationRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *RegistrationRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// RegistrationRequestValidationError is the validation error returned by
// RegistrationRequest.Validate if the designated constraints aren't met.
type RegistrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistrationRequestValidationError) ErrorName() string {
	return "RegistrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegistrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistrationRequestValidationError{}

// Validate checks the field values on EnrollmentRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EnrollmentRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetUserID()); err != nil {
		return EnrollmentRequestValidationError{
			field:  "UserID",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Role

	return nil
}

func (m *EnrollmentRequest) _validateUuid(uuid string) error {
	if matched := _enrollment_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// EnrollmentRequestValidationError is the validation error returned by
// EnrollmentRequest.Validate if the designated constraints aren't met.
type EnrollmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrollmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrollmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrollmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrollmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrollmentRequestValidationError) ErrorName() string {
	return "EnrollmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EnrollmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrollmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrollmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrollmentRequestValidationError{}

// Validate checks the field values on RegistrationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegistrationResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistrationResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AccessToken

	return nil
}

// RegistrationResponseValidationError is the validation error returned by
// RegistrationResponse.Validate if the designated constraints aren't met.
type RegistrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistrationResponseValidationError) ErrorName() string {
	return "RegistrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegistrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistrationResponseValidationError{}
