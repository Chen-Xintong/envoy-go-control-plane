// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/endpoint/v3/endpoint_components.proto

package endpointv3

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

	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

	_ = v3.HealthStatus(0)
)

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Endpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Endpoint with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EndpointMultiError, or nil
// if none found.
func (m *Endpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *Endpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EndpointValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EndpointValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHealthCheckConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EndpointValidationError{
					field:  "HealthCheckConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EndpointValidationError{
					field:  "HealthCheckConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHealthCheckConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointValidationError{
				field:  "HealthCheckConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Hostname

	if len(errors) > 0 {
		return EndpointMultiError(errors)
	}

	return nil
}

// EndpointMultiError is an error wrapping multiple validation errors returned
// by Endpoint.ValidateAll() if the designated constraints aren't met.
type EndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointMultiError) AllErrors() []error { return m }

// EndpointValidationError is the validation error returned by
// Endpoint.Validate if the designated constraints aren't met.
type EndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointValidationError) ErrorName() string { return "EndpointValidationError" }

// Error satisfies the builtin error interface
func (e EndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointValidationError{}

// Validate checks the field values on LbEndpoint with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LbEndpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LbEndpoint with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LbEndpointMultiError, or
// nil if none found.
func (m *LbEndpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *LbEndpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HealthStatus

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LbEndpointValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LbEndpointValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LbEndpointValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetLoadBalancingWeight(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := LbEndpointValidationError{
				field:  "LoadBalancingWeight",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	switch m.HostIdentifier.(type) {

	case *LbEndpoint_Endpoint:

		if all {
			switch v := interface{}(m.GetEndpoint()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LbEndpointValidationError{
						field:  "Endpoint",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LbEndpointValidationError{
						field:  "Endpoint",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LbEndpointValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LbEndpoint_EndpointName:
		// no validation rules for EndpointName

	}

	if len(errors) > 0 {
		return LbEndpointMultiError(errors)
	}

	return nil
}

// LbEndpointMultiError is an error wrapping multiple validation errors
// returned by LbEndpoint.ValidateAll() if the designated constraints aren't met.
type LbEndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LbEndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LbEndpointMultiError) AllErrors() []error { return m }

// LbEndpointValidationError is the validation error returned by
// LbEndpoint.Validate if the designated constraints aren't met.
type LbEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LbEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LbEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LbEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LbEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LbEndpointValidationError) ErrorName() string { return "LbEndpointValidationError" }

// Error satisfies the builtin error interface
func (e LbEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLbEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LbEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LbEndpointValidationError{}

// Validate checks the field values on LedsClusterLocalityConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LedsClusterLocalityConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LedsClusterLocalityConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LedsClusterLocalityConfigMultiError, or nil if none found.
func (m *LedsClusterLocalityConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *LedsClusterLocalityConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLedsConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LedsClusterLocalityConfigValidationError{
					field:  "LedsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LedsClusterLocalityConfigValidationError{
					field:  "LedsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLedsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LedsClusterLocalityConfigValidationError{
				field:  "LedsConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LedsCollectionName

	if len(errors) > 0 {
		return LedsClusterLocalityConfigMultiError(errors)
	}

	return nil
}

// LedsClusterLocalityConfigMultiError is an error wrapping multiple validation
// errors returned by LedsClusterLocalityConfig.ValidateAll() if the
// designated constraints aren't met.
type LedsClusterLocalityConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LedsClusterLocalityConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LedsClusterLocalityConfigMultiError) AllErrors() []error { return m }

// LedsClusterLocalityConfigValidationError is the validation error returned by
// LedsClusterLocalityConfig.Validate if the designated constraints aren't met.
type LedsClusterLocalityConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LedsClusterLocalityConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LedsClusterLocalityConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LedsClusterLocalityConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LedsClusterLocalityConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LedsClusterLocalityConfigValidationError) ErrorName() string {
	return "LedsClusterLocalityConfigValidationError"
}

// Error satisfies the builtin error interface
func (e LedsClusterLocalityConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLedsClusterLocalityConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LedsClusterLocalityConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LedsClusterLocalityConfigValidationError{}

// Validate checks the field values on LocalityLbEndpoints with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LocalityLbEndpoints) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalityLbEndpoints with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LocalityLbEndpointsMultiError, or nil if none found.
func (m *LocalityLbEndpoints) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityLbEndpoints) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocality()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalityLbEndpointsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalityLbEndpointsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityLbEndpointsValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetLbEndpoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  fmt.Sprintf("LbEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  fmt.Sprintf("LbEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbEndpointsValidationError{
					field:  fmt.Sprintf("LbEndpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if wrapper := m.GetLoadBalancingWeight(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := LocalityLbEndpointsValidationError{
				field:  "LoadBalancingWeight",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPriority() > 128 {
		err := LocalityLbEndpointsValidationError{
			field:  "Priority",
			reason: "value must be less than or equal to 128",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetProximity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalityLbEndpointsValidationError{
					field:  "Proximity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalityLbEndpointsValidationError{
					field:  "Proximity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProximity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityLbEndpointsValidationError{
				field:  "Proximity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.LbConfig.(type) {

	case *LocalityLbEndpoints_LoadBalancerEndpoints:

		if all {
			switch v := interface{}(m.GetLoadBalancerEndpoints()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  "LoadBalancerEndpoints",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  "LoadBalancerEndpoints",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLoadBalancerEndpoints()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbEndpointsValidationError{
					field:  "LoadBalancerEndpoints",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LocalityLbEndpoints_LedsClusterLocalityConfig:

		if all {
			switch v := interface{}(m.GetLedsClusterLocalityConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  "LedsClusterLocalityConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbEndpointsValidationError{
						field:  "LedsClusterLocalityConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLedsClusterLocalityConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbEndpointsValidationError{
					field:  "LedsClusterLocalityConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LocalityLbEndpointsMultiError(errors)
	}

	return nil
}

// LocalityLbEndpointsMultiError is an error wrapping multiple validation
// errors returned by LocalityLbEndpoints.ValidateAll() if the designated
// constraints aren't met.
type LocalityLbEndpointsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityLbEndpointsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityLbEndpointsMultiError) AllErrors() []error { return m }

// LocalityLbEndpointsValidationError is the validation error returned by
// LocalityLbEndpoints.Validate if the designated constraints aren't met.
type LocalityLbEndpointsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbEndpointsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbEndpointsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbEndpointsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbEndpointsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbEndpointsValidationError) ErrorName() string {
	return "LocalityLbEndpointsValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityLbEndpointsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbEndpoints.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbEndpointsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbEndpointsValidationError{}

// Validate checks the field values on Endpoint_HealthCheckConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Endpoint_HealthCheckConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Endpoint_HealthCheckConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Endpoint_HealthCheckConfigMultiError, or nil if none found.
func (m *Endpoint_HealthCheckConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Endpoint_HealthCheckConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPortValue() > 65535 {
		err := Endpoint_HealthCheckConfigValidationError{
			field:  "PortValue",
			reason: "value must be less than or equal to 65535",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Hostname

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Endpoint_HealthCheckConfigValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Endpoint_HealthCheckConfigValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Endpoint_HealthCheckConfigValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Endpoint_HealthCheckConfigMultiError(errors)
	}

	return nil
}

// Endpoint_HealthCheckConfigMultiError is an error wrapping multiple
// validation errors returned by Endpoint_HealthCheckConfig.ValidateAll() if
// the designated constraints aren't met.
type Endpoint_HealthCheckConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Endpoint_HealthCheckConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Endpoint_HealthCheckConfigMultiError) AllErrors() []error { return m }

// Endpoint_HealthCheckConfigValidationError is the validation error returned
// by Endpoint_HealthCheckConfig.Validate if the designated constraints aren't met.
type Endpoint_HealthCheckConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Endpoint_HealthCheckConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Endpoint_HealthCheckConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Endpoint_HealthCheckConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Endpoint_HealthCheckConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Endpoint_HealthCheckConfigValidationError) ErrorName() string {
	return "Endpoint_HealthCheckConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Endpoint_HealthCheckConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint_HealthCheckConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Endpoint_HealthCheckConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Endpoint_HealthCheckConfigValidationError{}

// Validate checks the field values on LocalityLbEndpoints_LbEndpointList with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *LocalityLbEndpoints_LbEndpointList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalityLbEndpoints_LbEndpointList
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// LocalityLbEndpoints_LbEndpointListMultiError, or nil if none found.
func (m *LocalityLbEndpoints_LbEndpointList) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityLbEndpoints_LbEndpointList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLbEndpoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbEndpoints_LbEndpointListValidationError{
						field:  fmt.Sprintf("LbEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbEndpoints_LbEndpointListValidationError{
						field:  fmt.Sprintf("LbEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbEndpoints_LbEndpointListValidationError{
					field:  fmt.Sprintf("LbEndpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LocalityLbEndpoints_LbEndpointListMultiError(errors)
	}

	return nil
}

// LocalityLbEndpoints_LbEndpointListMultiError is an error wrapping multiple
// validation errors returned by
// LocalityLbEndpoints_LbEndpointList.ValidateAll() if the designated
// constraints aren't met.
type LocalityLbEndpoints_LbEndpointListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityLbEndpoints_LbEndpointListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityLbEndpoints_LbEndpointListMultiError) AllErrors() []error { return m }

// LocalityLbEndpoints_LbEndpointListValidationError is the validation error
// returned by LocalityLbEndpoints_LbEndpointList.Validate if the designated
// constraints aren't met.
type LocalityLbEndpoints_LbEndpointListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbEndpoints_LbEndpointListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbEndpoints_LbEndpointListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbEndpoints_LbEndpointListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbEndpoints_LbEndpointListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbEndpoints_LbEndpointListValidationError) ErrorName() string {
	return "LocalityLbEndpoints_LbEndpointListValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityLbEndpoints_LbEndpointListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbEndpoints_LbEndpointList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbEndpoints_LbEndpointListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbEndpoints_LbEndpointListValidationError{}
