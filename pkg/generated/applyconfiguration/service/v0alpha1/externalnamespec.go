// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

// ExternalNameSpecApplyConfiguration represents a declarative configuration of the ExternalNameSpec type for use
// with apply.
type ExternalNameSpecApplyConfiguration struct {
	Host *string `json:"host,omitempty"`
}

// ExternalNameSpecApplyConfiguration constructs a declarative configuration of the ExternalNameSpec type for use with
// apply.
func ExternalNameSpec() *ExternalNameSpecApplyConfiguration {
	return &ExternalNameSpecApplyConfiguration{}
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *ExternalNameSpecApplyConfiguration) WithHost(value string) *ExternalNameSpecApplyConfiguration {
	b.Host = &value
	return b
}
