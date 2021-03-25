/*
Copyright (c) 2021 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package component

import (
	"github.com/vmware-tanzu/octant/internal/util/json"
)

// Signpost is a component for a signpost
// +octant:component
type Signpost struct {
	Base
	Config SignpostConfig `json:"config"`
}

// SignpostConfig is the contents of a signpost
type SignpostConfig struct {
	// Component is the component that will be displayed.
	Component Component `json:"component,omitempty"`
	// Trigger is the component that will trigger the signpost.
	Trigger Component `json:"trigger"`
	// Message is the text that will be displayed.
	Message string `json:"message"`
	// Reverse is the order of the components.
	Reverse bool `json:"reverse,omitempty"`
	// Column is the layout.
	Column bool `json:"column,omitempty"`
}

type signpostMarshal Signpost

// MarshalJSON implements json.Marshaler
func (t *Signpost) MarshalJSON() ([]byte, error) {
	m := signpostMarshal(*t)
	m.Metadata.Type = TypeSignPost
	return json.Marshal(&m)
}

func NewSignpost(c Component, t Component, m string) *Signpost {
	so := &Signpost{
		Base: newBase(TypeSignPost, nil),
		Config: SignpostConfig{
			Component: c,
			Trigger:   t,
			Message:   m,
			Reverse:   false,
			Column:    false,
		},
	}

	return so
}
