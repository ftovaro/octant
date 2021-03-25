/*
Copyright (c) 2021 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package component

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware-tanzu/octant/internal/util/json"
)

func TestSignpost_NewSignpost(t *testing.T) {
	so := NewSignpost(NewText("test"), NewIcon("user"), "Message")
	slice, _ := so.MarshalJSON()
	test := &Signpost{
		Base: newBase(TypeSignPost, nil),
		Config: SignpostConfig{
			Component: NewText("test"),
			Trigger:   NewIcon("user"),
			Message:   "Message",
			Reverse:   false,
			Column:    false,
		},
	}
	fmt.Println("Test")
	fmt.Printf("+%v", test.UnmarshalJSON(slice))
	fmt.Println("End Test")
}

func Test_Signpost_Marshal(t *testing.T) {
	test := []struct {
		name         string
		input        Component
		expectedPath string
		isErr        bool
	}{
		{
			name: "in general",
			input: &Signpost{
				Base: newBase(TypeSignPost, nil),
				Config: SignpostConfig{
					Component: NewText("test"),
					Trigger:   NewIcon("user"),
					Message:   "Message",
					Reverse:   false,
					Column:    false,
				},
			},
			expectedPath: "signpost.json",
			isErr:        false,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := json.Marshal(tc.input)
			isErr := err != nil
			if isErr != tc.isErr {
				t.Fatalf("UnExpected error: %v", err)
			}

			expected, err := ioutil.ReadFile(path.Join("testdata", tc.expectedPath))
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
