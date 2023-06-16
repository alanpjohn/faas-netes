/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// ProfileSpecApplyConfiguration represents an declarative configuration of the ProfileSpec type for use
// with apply.
type ProfileSpecApplyConfiguration struct {
	Tolerations               []v1.Toleration               `json:"tolerations,omitempty"`
	RuntimeClassName          *string                       `json:"runtimeClassName,omitempty"`
	PodSecurityContext        *v1.PodSecurityContext        `json:"podSecurityContext,omitempty"`
	Affinity                  *v1.Affinity                  `json:"affinity,omitempty"`
	TopologySpreadConstraints []v1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}

// ProfileSpecApplyConfiguration constructs an declarative configuration of the ProfileSpec type for use with
// apply.
func ProfileSpec() *ProfileSpecApplyConfiguration {
	return &ProfileSpecApplyConfiguration{}
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *ProfileSpecApplyConfiguration) WithTolerations(values ...v1.Toleration) *ProfileSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithRuntimeClassName sets the RuntimeClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuntimeClassName field is set to the value of the last call.
func (b *ProfileSpecApplyConfiguration) WithRuntimeClassName(value string) *ProfileSpecApplyConfiguration {
	b.RuntimeClassName = &value
	return b
}

// WithPodSecurityContext sets the PodSecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodSecurityContext field is set to the value of the last call.
func (b *ProfileSpecApplyConfiguration) WithPodSecurityContext(value v1.PodSecurityContext) *ProfileSpecApplyConfiguration {
	b.PodSecurityContext = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *ProfileSpecApplyConfiguration) WithAffinity(value v1.Affinity) *ProfileSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *ProfileSpecApplyConfiguration) WithTopologySpreadConstraints(values ...v1.TopologySpreadConstraint) *ProfileSpecApplyConfiguration {
	for i := range values {
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, values[i])
	}
	return b
}