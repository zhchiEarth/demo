// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/account.proto

package v1

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

// Validate checks the field values on AccountListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AccountListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListRequestMultiError, or nil if none found.
func (m *AccountListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAddresses()) > 0 {

		_AccountListRequest_Addresses_Unique := make(map[string]struct{}, len(m.GetAddresses()))

		for idx, item := range m.GetAddresses() {
			_, _ = idx, item

			if _, exists := _AccountListRequest_Addresses_Unique[item]; exists {
				err := AccountListRequestValidationError{
					field:  fmt.Sprintf("Addresses[%v]", idx),
					reason: "repeated value must contain unique items",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			} else {
				_AccountListRequest_Addresses_Unique[item] = struct{}{}
			}

			if len(item) > 42 {
				err := AccountListRequestValidationError{
					field:  fmt.Sprintf("Addresses[%v]", idx),
					reason: "value length must be at most 42 bytes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_AccountListRequest_Addresses_Pattern.MatchString(item) {
				err := AccountListRequestValidationError{
					field:  fmt.Sprintf("Addresses[%v]", idx),
					reason: "value does not match regex pattern \"^0x[0-9a-fA-F]{40}$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}

	}

	if m.GetMaxHealth() != 0 {

		if m.GetMaxHealth() <= 0 {
			err := AccountListRequestValidationError{
				field:  "MaxHealth",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetMinBorrowValueInEth() != 0 {

		if m.GetMinBorrowValueInEth() <= 0 {
			err := AccountListRequestValidationError{
				field:  "MinBorrowValueInEth",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPageNumber() != 0 {

		if m.GetPageNumber() <= 0 {
			err := AccountListRequestValidationError{
				field:  "PageNumber",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPageSize() != 0 {

		if val := m.GetPageSize(); val <= 0 || val > 100 {
			err := AccountListRequestValidationError{
				field:  "PageSize",
				reason: "value must be inside range (0, 100]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return AccountListRequestMultiError(errors)
	}

	return nil
}

// AccountListRequestMultiError is an error wrapping multiple validation errors
// returned by AccountListRequest.ValidateAll() if the designated constraints
// aren't met.
type AccountListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListRequestMultiError) AllErrors() []error { return m }

// AccountListRequestValidationError is the validation error returned by
// AccountListRequest.Validate if the designated constraints aren't met.
type AccountListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListRequestValidationError) ErrorName() string {
	return "AccountListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AccountListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListRequestValidationError{}

var _AccountListRequest_Addresses_Pattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// Validate checks the field values on AccountListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AccountListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListReplyMultiError, or nil if none found.
func (m *AccountListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AccountListReplyValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AccountListReplyValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccountListReplyValidationError{
					field:  fmt.Sprintf("Accounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPaginationSummary()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountListReplyValidationError{
					field:  "PaginationSummary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountListReplyValidationError{
					field:  "PaginationSummary",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginationSummary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountListReplyValidationError{
				field:  "PaginationSummary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountListReplyValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountListReplyValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountListReplyValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccountListReplyMultiError(errors)
	}

	return nil
}

// AccountListReplyMultiError is an error wrapping multiple validation errors
// returned by AccountListReply.ValidateAll() if the designated constraints
// aren't met.
type AccountListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListReplyMultiError) AllErrors() []error { return m }

// AccountListReplyValidationError is the validation error returned by
// AccountListReply.Validate if the designated constraints aren't met.
type AccountListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListReplyValidationError) ErrorName() string { return "AccountListReplyValidationError" }

// Error satisfies the builtin error interface
func (e AccountListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListReplyValidationError{}

// Validate checks the field values on PaginationSummary with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PaginationSummary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginationSummary with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaginationSummaryMultiError, or nil if none found.
func (m *PaginationSummary) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginationSummary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageNumber

	// no validation rules for PageSize

	// no validation rules for TotalEntries

	// no validation rules for TotalPages

	if len(errors) > 0 {
		return PaginationSummaryMultiError(errors)
	}

	return nil
}

// PaginationSummaryMultiError is an error wrapping multiple validation errors
// returned by PaginationSummary.ValidateAll() if the designated constraints
// aren't met.
type PaginationSummaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationSummaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationSummaryMultiError) AllErrors() []error { return m }

// PaginationSummaryValidationError is the validation error returned by
// PaginationSummary.Validate if the designated constraints aren't met.
type PaginationSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationSummaryValidationError) ErrorName() string {
	return "PaginationSummaryValidationError"
}

// Error satisfies the builtin error interface
func (e PaginationSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationSummaryValidationError{}

// Validate checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Token) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokenMultiError, or nil if none found.
func (m *Token) ValidateAll() error {
	return m.validate(true)
}

func (m *Token) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Symbol

	// no validation rules for BorrowBalanceUnderlying

	// no validation rules for SupplyBalanceUnderlying

	// no validation rules for LifetimeBorrowInterestAccrued

	// no validation rules for LifetimeSupplyInterestAccrued

	// no validation rules for SafeWithdrawAmountUnderlying

	if len(errors) > 0 {
		return TokenMultiError(errors)
	}

	return nil
}

// TokenMultiError is an error wrapping multiple validation errors returned by
// Token.ValidateAll() if the designated constraints aren't met.
type TokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenMultiError) AllErrors() []error { return m }

// TokenValidationError is the validation error returned by Token.Validate if
// the designated constraints aren't met.
type TokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenValidationError) ErrorName() string { return "TokenValidationError" }

// Error satisfies the builtin error interface
func (e TokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenValidationError{}

// Validate checks the field values on AccountListReply_Account with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AccountListReply_Account) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListReply_Account with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListReply_AccountMultiError, or nil if none found.
func (m *AccountListReply_Account) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListReply_Account) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Health

	for idx, item := range m.GetTokens() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AccountListReply_AccountValidationError{
						field:  fmt.Sprintf("Tokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AccountListReply_AccountValidationError{
						field:  fmt.Sprintf("Tokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccountListReply_AccountValidationError{
					field:  fmt.Sprintf("Tokens[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalBorrowValueInEth

	// no validation rules for TotalCollateralValueInEth

	if len(errors) > 0 {
		return AccountListReply_AccountMultiError(errors)
	}

	return nil
}

// AccountListReply_AccountMultiError is an error wrapping multiple validation
// errors returned by AccountListReply_Account.ValidateAll() if the designated
// constraints aren't met.
type AccountListReply_AccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListReply_AccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListReply_AccountMultiError) AllErrors() []error { return m }

// AccountListReply_AccountValidationError is the validation error returned by
// AccountListReply_Account.Validate if the designated constraints aren't met.
type AccountListReply_AccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListReply_AccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListReply_AccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListReply_AccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListReply_AccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListReply_AccountValidationError) ErrorName() string {
	return "AccountListReply_AccountValidationError"
}

// Error satisfies the builtin error interface
func (e AccountListReply_AccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListReply_Account.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListReply_AccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListReply_AccountValidationError{}