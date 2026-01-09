/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package imds

import (
	"testing"
)

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"instance-id", "meta-data/instance-id"},
		{"meta-data/instance-id", "meta-data/instance-id"},
		{"dynamic/instance-identity/document", "dynamic/instance-identity/document"},
		{"user-data", "user-data"},
		{"placement/region", "meta-data/placement/region"},
		{"/instance-id/", "meta-data/instance-id"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := NormalizePath(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePath(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsDirectory(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{"single value", []byte("i-1234567890abcdef0"), false},
		{"directory listing", []byte("instance-id\ninstance-type\nami-id"), true},
		{"single with newline", []byte("value\n"), false},
		{"json object", []byte(`{"key": "value"}`), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDirectory(tt.input)
			if result != tt.expected {
				t.Errorf("IsDirectory(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindSimilar(t *testing.T) {
	keys := []string{
		"instance-id",
		"instance-type",
		"availability-zone",
		"region",
		"ami-id",
	}

	tests := []struct {
		query       string
		expectFirst string
		expectEmpty bool
	}{
		{"instance-id", "instance-id", false},
		{"instanc-id", "instance-id", false},
		{"instance", "instance-id", false},
		{"region", "region", false},
		{"regin", "region", false},
		{"xyznotreal", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			result := FindSimilar(tt.query, keys, 5)
			if tt.expectEmpty {
				if len(result) != 0 {
					t.Errorf("FindSimilar(%q) = %v, want empty", tt.query, result)
				}
			} else {
				if len(result) == 0 {
					t.Errorf("FindSimilar(%q) returned empty, want %q first", tt.query, tt.expectFirst)
				} else if result[0] != tt.expectFirst {
					t.Errorf("FindSimilar(%q)[0] = %q, want %q", tt.query, result[0], tt.expectFirst)
				}
			}
		})
	}
}

func TestParseJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		wantOK bool
	}{
		{"object", []byte(`{"key": "value"}`), true},
		{"array", []byte(`[1, 2, 3]`), true},
		{"string", []byte(`"hello"`), false},
		{"number", []byte(`123`), false},
		{"invalid", []byte(`not json`), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := parseJSON(tt.input)
			if ok != tt.wantOK {
				t.Errorf("parseJSON(%q) ok = %v, want %v", tt.input, ok, tt.wantOK)
			}
		})
	}
}
