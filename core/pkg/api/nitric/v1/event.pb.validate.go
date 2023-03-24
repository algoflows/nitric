// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/event/v1/event.proto

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

// Validate checks the field values on EventPublishRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EventPublishRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventPublishRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventPublishRequestMultiError, or nil if none found.
func (m *EventPublishRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EventPublishRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTopic()) > 256 {
		err := EventPublishRequestValidationError{
			field:  "Topic",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_EventPublishRequest_Topic_Pattern.MatchString(m.GetTopic()) {
		err := EventPublishRequestValidationError{
			field:  "Topic",
			reason: "value does not match regex pattern \"^\\\\w+([.\\\\-]\\\\w+)*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEvent() == nil {
		err := EventPublishRequestValidationError{
			field:  "Event",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEvent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventPublishRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventPublishRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventPublishRequestValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetDelay() != 0 {

		if val := m.GetDelay(); val < 10 || val > 2592000 {
			err := EventPublishRequestValidationError{
				field:  "Delay",
				reason: "value must be inside range [10, 2592000]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return EventPublishRequestMultiError(errors)
	}

	return nil
}

// EventPublishRequestMultiError is an error wrapping multiple validation
// errors returned by EventPublishRequest.ValidateAll() if the designated
// constraints aren't met.
type EventPublishRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventPublishRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventPublishRequestMultiError) AllErrors() []error { return m }

// EventPublishRequestValidationError is the validation error returned by
// EventPublishRequest.Validate if the designated constraints aren't met.
type EventPublishRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventPublishRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventPublishRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventPublishRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventPublishRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventPublishRequestValidationError) ErrorName() string {
	return "EventPublishRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EventPublishRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventPublishRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventPublishRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventPublishRequestValidationError{}

var _EventPublishRequest_Topic_Pattern = regexp.MustCompile("^\\w+([.\\-]\\w+)*$")

// Validate checks the field values on EventPublishResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EventPublishResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventPublishResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventPublishResponseMultiError, or nil if none found.
func (m *EventPublishResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EventPublishResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return EventPublishResponseMultiError(errors)
	}

	return nil
}

// EventPublishResponseMultiError is an error wrapping multiple validation
// errors returned by EventPublishResponse.ValidateAll() if the designated
// constraints aren't met.
type EventPublishResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventPublishResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventPublishResponseMultiError) AllErrors() []error { return m }

// EventPublishResponseValidationError is the validation error returned by
// EventPublishResponse.Validate if the designated constraints aren't met.
type EventPublishResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventPublishResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventPublishResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventPublishResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventPublishResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventPublishResponseValidationError) ErrorName() string {
	return "EventPublishResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EventPublishResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventPublishResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventPublishResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventPublishResponseValidationError{}

// Validate checks the field values on TopicListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TopicListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopicListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TopicListRequestMultiError, or nil if none found.
func (m *TopicListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TopicListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TopicListRequestMultiError(errors)
	}

	return nil
}

// TopicListRequestMultiError is an error wrapping multiple validation errors
// returned by TopicListRequest.ValidateAll() if the designated constraints
// aren't met.
type TopicListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopicListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopicListRequestMultiError) AllErrors() []error { return m }

// TopicListRequestValidationError is the validation error returned by
// TopicListRequest.Validate if the designated constraints aren't met.
type TopicListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopicListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopicListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopicListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopicListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopicListRequestValidationError) ErrorName() string { return "TopicListRequestValidationError" }

// Error satisfies the builtin error interface
func (e TopicListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopicListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopicListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopicListRequestValidationError{}

// Validate checks the field values on TopicListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TopicListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopicListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TopicListResponseMultiError, or nil if none found.
func (m *TopicListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TopicListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTopics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TopicListResponseValidationError{
						field:  fmt.Sprintf("Topics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TopicListResponseValidationError{
						field:  fmt.Sprintf("Topics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TopicListResponseValidationError{
					field:  fmt.Sprintf("Topics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TopicListResponseMultiError(errors)
	}

	return nil
}

// TopicListResponseMultiError is an error wrapping multiple validation errors
// returned by TopicListResponse.ValidateAll() if the designated constraints
// aren't met.
type TopicListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopicListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopicListResponseMultiError) AllErrors() []error { return m }

// TopicListResponseValidationError is the validation error returned by
// TopicListResponse.Validate if the designated constraints aren't met.
type TopicListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopicListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopicListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopicListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopicListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopicListResponseValidationError) ErrorName() string {
	return "TopicListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TopicListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopicListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopicListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopicListResponseValidationError{}

// Validate checks the field values on NitricTopic with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NitricTopic) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NitricTopic with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NitricTopicMultiError, or
// nil if none found.
func (m *NitricTopic) ValidateAll() error {
	return m.validate(true)
}

func (m *NitricTopic) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return NitricTopicMultiError(errors)
	}

	return nil
}

// NitricTopicMultiError is an error wrapping multiple validation errors
// returned by NitricTopic.ValidateAll() if the designated constraints aren't met.
type NitricTopicMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NitricTopicMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NitricTopicMultiError) AllErrors() []error { return m }

// NitricTopicValidationError is the validation error returned by
// NitricTopic.Validate if the designated constraints aren't met.
type NitricTopicValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NitricTopicValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NitricTopicValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NitricTopicValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NitricTopicValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NitricTopicValidationError) ErrorName() string { return "NitricTopicValidationError" }

// Error satisfies the builtin error interface
func (e NitricTopicValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNitricTopic.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NitricTopicValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NitricTopicValidationError{}

// Validate checks the field values on NitricEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NitricEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NitricEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NitricEventMultiError, or
// nil if none found.
func (m *NitricEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *NitricEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for PayloadType

	if all {
		switch v := interface{}(m.GetPayload()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NitricEventValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NitricEventValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NitricEventValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NitricEventMultiError(errors)
	}

	return nil
}

// NitricEventMultiError is an error wrapping multiple validation errors
// returned by NitricEvent.ValidateAll() if the designated constraints aren't met.
type NitricEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NitricEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NitricEventMultiError) AllErrors() []error { return m }

// NitricEventValidationError is the validation error returned by
// NitricEvent.Validate if the designated constraints aren't met.
type NitricEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NitricEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NitricEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NitricEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NitricEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NitricEventValidationError) ErrorName() string { return "NitricEventValidationError" }

// Error satisfies the builtin error interface
func (e NitricEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNitricEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NitricEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NitricEventValidationError{}
