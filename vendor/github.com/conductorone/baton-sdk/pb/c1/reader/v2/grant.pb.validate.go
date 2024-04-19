// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: c1/reader/v2/grant.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on GrantsReaderServiceGetGrantRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GrantsReaderServiceGetGrantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantsReaderServiceGetGrantRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GrantsReaderServiceGetGrantRequestMultiError, or nil if none found.
func (m *GrantsReaderServiceGetGrantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceGetGrantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := len(m.GetGrantId()); l < 1 || l > 1024 {
		err := GrantsReaderServiceGetGrantRequestValidationError{
			field:  "GrantId",
			reason: "value length must be between 1 and 1024 bytes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GrantsReaderServiceGetGrantRequestMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceGetGrantRequestMultiError is an error wrapping multiple
// validation errors returned by
// GrantsReaderServiceGetGrantRequest.ValidateAll() if the designated
// constraints aren't met.
type GrantsReaderServiceGetGrantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceGetGrantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceGetGrantRequestMultiError) AllErrors() []error { return m }

// GrantsReaderServiceGetGrantRequestValidationError is the validation error
// returned by GrantsReaderServiceGetGrantRequest.Validate if the designated
// constraints aren't met.
type GrantsReaderServiceGetGrantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceGetGrantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantsReaderServiceGetGrantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantsReaderServiceGetGrantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantsReaderServiceGetGrantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceGetGrantRequestValidationError) ErrorName() string {
	return "GrantsReaderServiceGetGrantRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceGetGrantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceGetGrantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceGetGrantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceGetGrantRequestValidationError{}

// Validate checks the field values on GrantsReaderServiceGetGrantResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GrantsReaderServiceGetGrantResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantsReaderServiceGetGrantResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GrantsReaderServiceGetGrantResponseMultiError, or nil if none found.
func (m *GrantsReaderServiceGetGrantResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceGetGrantResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetGrant()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrantsReaderServiceGetGrantResponseValidationError{
					field:  "Grant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrantsReaderServiceGetGrantResponseValidationError{
					field:  "Grant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrantsReaderServiceGetGrantResponseValidationError{
				field:  "Grant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GrantsReaderServiceGetGrantResponseMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceGetGrantResponseMultiError is an error wrapping multiple
// validation errors returned by
// GrantsReaderServiceGetGrantResponse.ValidateAll() if the designated
// constraints aren't met.
type GrantsReaderServiceGetGrantResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceGetGrantResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceGetGrantResponseMultiError) AllErrors() []error { return m }

// GrantsReaderServiceGetGrantResponseValidationError is the validation error
// returned by GrantsReaderServiceGetGrantResponse.Validate if the designated
// constraints aren't met.
type GrantsReaderServiceGetGrantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceGetGrantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantsReaderServiceGetGrantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantsReaderServiceGetGrantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantsReaderServiceGetGrantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceGetGrantResponseValidationError) ErrorName() string {
	return "GrantsReaderServiceGetGrantResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceGetGrantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceGetGrantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceGetGrantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceGetGrantResponseValidationError{}

// Validate checks the field values on
// GrantsReaderServiceListGrantsForEntitlementRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrantsReaderServiceListGrantsForEntitlementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrantsReaderServiceListGrantsForEntitlementRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// GrantsReaderServiceListGrantsForEntitlementRequestMultiError, or nil if
// none found.
func (m *GrantsReaderServiceListGrantsForEntitlementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceListGrantsForEntitlementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEntitlement() == nil {
		err := GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
			field:  "Entitlement",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEntitlement()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
					field:  "Entitlement",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
					field:  "Entitlement",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEntitlement()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
				field:  "Entitlement",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPrincipalId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
					field:  "PrincipalId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
					field:  "PrincipalId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipalId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
				field:  "PrincipalId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetPageSize() != 0 {

		if m.GetPageSize() > 250 {
			err := GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
				field:  "PageSize",
				reason: "value must be less than or equal to 250",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPageToken() != "" {

		if l := len(m.GetPageToken()); l < 1 || l > 2048 {
			err := GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
				field:  "PageToken",
				reason: "value length must be between 1 and 2048 bytes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	for idx, item := range m.GetAnnotations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrantsReaderServiceListGrantsForEntitlementRequestValidationError{
					field:  fmt.Sprintf("Annotations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GrantsReaderServiceListGrantsForEntitlementRequestMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceListGrantsForEntitlementRequestMultiError is an error
// wrapping multiple validation errors returned by
// GrantsReaderServiceListGrantsForEntitlementRequest.ValidateAll() if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForEntitlementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceListGrantsForEntitlementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceListGrantsForEntitlementRequestMultiError) AllErrors() []error { return m }

// GrantsReaderServiceListGrantsForEntitlementRequestValidationError is the
// validation error returned by
// GrantsReaderServiceListGrantsForEntitlementRequest.Validate if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForEntitlementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) ErrorName() string {
	return "GrantsReaderServiceListGrantsForEntitlementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceListGrantsForEntitlementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceListGrantsForEntitlementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceListGrantsForEntitlementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceListGrantsForEntitlementRequestValidationError{}

// Validate checks the field values on
// GrantsReaderServiceListGrantsForEntitlementResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrantsReaderServiceListGrantsForEntitlementResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrantsReaderServiceListGrantsForEntitlementResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// GrantsReaderServiceListGrantsForEntitlementResponseMultiError, or nil if
// none found.
func (m *GrantsReaderServiceListGrantsForEntitlementResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceListGrantsForEntitlementResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForEntitlementResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForEntitlementResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrantsReaderServiceListGrantsForEntitlementResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetNextPageToken() != "" {

		if l := len(m.GetNextPageToken()); l < 1 || l > 2048 {
			err := GrantsReaderServiceListGrantsForEntitlementResponseValidationError{
				field:  "NextPageToken",
				reason: "value length must be between 1 and 2048 bytes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GrantsReaderServiceListGrantsForEntitlementResponseMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceListGrantsForEntitlementResponseMultiError is an error
// wrapping multiple validation errors returned by
// GrantsReaderServiceListGrantsForEntitlementResponse.ValidateAll() if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForEntitlementResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceListGrantsForEntitlementResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceListGrantsForEntitlementResponseMultiError) AllErrors() []error { return m }

// GrantsReaderServiceListGrantsForEntitlementResponseValidationError is the
// validation error returned by
// GrantsReaderServiceListGrantsForEntitlementResponse.Validate if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForEntitlementResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) ErrorName() string {
	return "GrantsReaderServiceListGrantsForEntitlementResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceListGrantsForEntitlementResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceListGrantsForEntitlementResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceListGrantsForEntitlementResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceListGrantsForEntitlementResponseValidationError{}

// Validate checks the field values on
// GrantsReaderServiceListGrantsForResourceTypeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrantsReaderServiceListGrantsForResourceTypeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrantsReaderServiceListGrantsForResourceTypeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// GrantsReaderServiceListGrantsForResourceTypeRequestMultiError, or nil if
// none found.
func (m *GrantsReaderServiceListGrantsForResourceTypeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceListGrantsForResourceTypeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetResourceTypeId() != "" {

		if l := len(m.GetResourceTypeId()); l < 1 || l > 2048 {
			err := GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
				field:  "ResourceTypeId",
				reason: "value length must be between 1 and 2048 bytes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPageSize() != 0 {

		if m.GetPageSize() > 250 {
			err := GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
				field:  "PageSize",
				reason: "value must be less than or equal to 250",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPageToken() != "" {

		if l := len(m.GetPageToken()); l < 1 || l > 2048 {
			err := GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
				field:  "PageToken",
				reason: "value length must be between 1 and 2048 bytes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	for idx, item := range m.GetAnnotations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{
					field:  fmt.Sprintf("Annotations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GrantsReaderServiceListGrantsForResourceTypeRequestMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceListGrantsForResourceTypeRequestMultiError is an error
// wrapping multiple validation errors returned by
// GrantsReaderServiceListGrantsForResourceTypeRequest.ValidateAll() if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForResourceTypeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceListGrantsForResourceTypeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceListGrantsForResourceTypeRequestMultiError) AllErrors() []error { return m }

// GrantsReaderServiceListGrantsForResourceTypeRequestValidationError is the
// validation error returned by
// GrantsReaderServiceListGrantsForResourceTypeRequest.Validate if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForResourceTypeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) ErrorName() string {
	return "GrantsReaderServiceListGrantsForResourceTypeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceListGrantsForResourceTypeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceListGrantsForResourceTypeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceListGrantsForResourceTypeRequestValidationError{}

// Validate checks the field values on
// GrantsReaderServiceListGrantsForResourceTypeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrantsReaderServiceListGrantsForResourceTypeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrantsReaderServiceListGrantsForResourceTypeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// GrantsReaderServiceListGrantsForResourceTypeResponseMultiError, or nil if
// none found.
func (m *GrantsReaderServiceListGrantsForResourceTypeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantsReaderServiceListGrantsForResourceTypeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetNextPageToken() != "" {

		if l := len(m.GetNextPageToken()); l < 1 || l > 2048 {
			err := GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{
				field:  "NextPageToken",
				reason: "value length must be between 1 and 2048 bytes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GrantsReaderServiceListGrantsForResourceTypeResponseMultiError(errors)
	}

	return nil
}

// GrantsReaderServiceListGrantsForResourceTypeResponseMultiError is an error
// wrapping multiple validation errors returned by
// GrantsReaderServiceListGrantsForResourceTypeResponse.ValidateAll() if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForResourceTypeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantsReaderServiceListGrantsForResourceTypeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantsReaderServiceListGrantsForResourceTypeResponseMultiError) AllErrors() []error { return m }

// GrantsReaderServiceListGrantsForResourceTypeResponseValidationError is the
// validation error returned by
// GrantsReaderServiceListGrantsForResourceTypeResponse.Validate if the
// designated constraints aren't met.
type GrantsReaderServiceListGrantsForResourceTypeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) ErrorName() string {
	return "GrantsReaderServiceListGrantsForResourceTypeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GrantsReaderServiceListGrantsForResourceTypeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantsReaderServiceListGrantsForResourceTypeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantsReaderServiceListGrantsForResourceTypeResponseValidationError{}