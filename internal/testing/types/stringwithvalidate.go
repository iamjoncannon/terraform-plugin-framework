package types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ attr.TypeWithValidate = StringTypeWithValidateError{}
	_ attr.TypeWithValidate = StringTypeWithValidateWarning{}
)

type StringTypeWithValidateError struct {
	StringType
}

type StringTypeWithValidateWarning struct {
	StringType
}

func (t StringTypeWithValidateError) Validate(ctx context.Context, in tftypes.Value) []*tfprotov6.Diagnostic {
	return []*tfprotov6.Diagnostic{TestErrorDiagnostic}
}

func (t StringTypeWithValidateWarning) Validate(ctx context.Context, in tftypes.Value) []*tfprotov6.Diagnostic {
	return []*tfprotov6.Diagnostic{TestWarningDiagnostic}
}
