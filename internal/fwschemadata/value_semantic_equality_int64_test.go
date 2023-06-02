// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fwschemadata_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschemadata"
	testtypes "github.com/hashicorp/terraform-plugin-framework/internal/testing/types"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestValueSemanticEqualityInt64(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request  fwschemadata.ValueSemanticEqualityRequest
		expected *fwschemadata.ValueSemanticEqualityResponse
	}{
		"Int64Value": {
			request: fwschemadata.ValueSemanticEqualityRequest{
				Path:             path.Root("test"),
				PriorValue:       types.Int64Value(12),
				ProposedNewValue: types.Int64Value(24),
			},
			expected: &fwschemadata.ValueSemanticEqualityResponse{
				NewValue: types.Int64Value(24),
			},
		},
		"Int64ValuableWithSemanticEquals-true": {
			request: fwschemadata.ValueSemanticEqualityRequest{
				Path: path.Root("test"),
				PriorValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(12),
					SemanticEquals: true,
				},
				ProposedNewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(24),
					SemanticEquals: true,
				},
			},
			expected: &fwschemadata.ValueSemanticEqualityResponse{
				NewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(12),
					SemanticEquals: true,
				},
			},
		},
		"Int64ValuableWithSemanticEquals-false": {
			request: fwschemadata.ValueSemanticEqualityRequest{
				Path: path.Root("test"),
				PriorValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(12),
					SemanticEquals: false,
				},
				ProposedNewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(24),
					SemanticEquals: false,
				},
			},
			expected: &fwschemadata.ValueSemanticEqualityResponse{
				NewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(24),
					SemanticEquals: false,
				},
			},
		},
		"Int64ValuableWithSemanticEquals-diagnostics": {
			request: fwschemadata.ValueSemanticEqualityRequest{
				Path: path.Root("test"),
				PriorValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(12),
					SemanticEquals: false,
					SemanticEqualsDiagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic("test summary 1", "test detail 1"),
						diag.NewErrorDiagnostic("test summary 2", "test detail 2"),
					},
				},
				ProposedNewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(24),
					SemanticEquals: false,
					SemanticEqualsDiagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic("test summary 1", "test detail 1"),
						diag.NewErrorDiagnostic("test summary 2", "test detail 2"),
					},
				},
			},
			expected: &fwschemadata.ValueSemanticEqualityResponse{
				NewValue: testtypes.Int64ValueWithSemanticEquals{
					Int64Value:     types.Int64Value(24),
					SemanticEquals: false,
					SemanticEqualsDiagnostics: diag.Diagnostics{
						diag.NewErrorDiagnostic("test summary 1", "test detail 1"),
						diag.NewErrorDiagnostic("test summary 2", "test detail 2"),
					},
				},
				Diagnostics: diag.Diagnostics{
					diag.NewErrorDiagnostic("test summary 1", "test detail 1"),
					diag.NewErrorDiagnostic("test summary 2", "test detail 2"),
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := &fwschemadata.ValueSemanticEqualityResponse{
				NewValue: testCase.request.ProposedNewValue,
			}

			fwschemadata.ValueSemanticEqualityInt64(context.Background(), testCase.request, got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
