// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamicvalidator_test

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/dynamicvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ExampleAlsoRequires() {
	// Used within a Schema method of a DataSource, Provider, or Resource
	_ = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"example_attr": schema.DynamicAttribute{
				Optional: true,
				Validators: []validator.Dynamic{
					// Validate this attribute must be configured with other_attr.
					dynamicvalidator.AlsoRequires(path.Expressions{
						path.MatchRoot("other_attr"),
					}...),
				},
			},
			"other_attr": schema.DynamicAttribute{
				Optional: true,
			},
		},
	}
}
