// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLambdaPermission(name string, args LambdaPermissionArgs) *LambdaPermission {
	return &LambdaPermission{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaPermission)(nil)

type LambdaPermission struct {
	Name  string
	Args  LambdaPermissionArgs
	state *lambdaPermissionState
}

func (lp *LambdaPermission) Type() string {
	return "aws_lambda_permission"
}

func (lp *LambdaPermission) LocalName() string {
	return lp.Name
}

func (lp *LambdaPermission) Configuration() interface{} {
	return lp.Args
}

func (lp *LambdaPermission) Attributes() lambdaPermissionAttributes {
	return lambdaPermissionAttributes{ref: terra.ReferenceResource(lp)}
}

func (lp *LambdaPermission) ImportState(av io.Reader) error {
	lp.state = &lambdaPermissionState{}
	if err := json.NewDecoder(av).Decode(lp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lp.Type(), lp.LocalName(), err)
	}
	return nil
}

func (lp *LambdaPermission) State() (*lambdaPermissionState, bool) {
	return lp.state, lp.state != nil
}

func (lp *LambdaPermission) StateMust() *lambdaPermissionState {
	if lp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lp.Type(), lp.LocalName()))
	}
	return lp.state
}

func (lp *LambdaPermission) DependOn() terra.Reference {
	return terra.ReferenceResource(lp)
}

type LambdaPermissionArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// EventSourceToken: string, optional
	EventSourceToken terra.StringValue `hcl:"event_source_token,attr"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// FunctionUrlAuthType: string, optional
	FunctionUrlAuthType terra.StringValue `hcl:"function_url_auth_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// PrincipalOrgId: string, optional
	PrincipalOrgId terra.StringValue `hcl:"principal_org_id,attr"`
	// Qualifier: string, optional
	Qualifier terra.StringValue `hcl:"qualifier,attr"`
	// SourceAccount: string, optional
	SourceAccount terra.StringValue `hcl:"source_account,attr"`
	// SourceArn: string, optional
	SourceArn terra.StringValue `hcl:"source_arn,attr"`
	// StatementId: string, optional
	StatementId terra.StringValue `hcl:"statement_id,attr"`
	// StatementIdPrefix: string, optional
	StatementIdPrefix terra.StringValue `hcl:"statement_id_prefix,attr"`
	// DependsOn contains resources that LambdaPermission depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lambdaPermissionAttributes struct {
	ref terra.Reference
}

func (lp lambdaPermissionAttributes) Action() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("action"))
}

func (lp lambdaPermissionAttributes) EventSourceToken() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("event_source_token"))
}

func (lp lambdaPermissionAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("function_name"))
}

func (lp lambdaPermissionAttributes) FunctionUrlAuthType() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("function_url_auth_type"))
}

func (lp lambdaPermissionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("id"))
}

func (lp lambdaPermissionAttributes) Principal() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("principal"))
}

func (lp lambdaPermissionAttributes) PrincipalOrgId() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("principal_org_id"))
}

func (lp lambdaPermissionAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("qualifier"))
}

func (lp lambdaPermissionAttributes) SourceAccount() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("source_account"))
}

func (lp lambdaPermissionAttributes) SourceArn() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("source_arn"))
}

func (lp lambdaPermissionAttributes) StatementId() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("statement_id"))
}

func (lp lambdaPermissionAttributes) StatementIdPrefix() terra.StringValue {
	return terra.ReferenceString(lp.ref.Append("statement_id_prefix"))
}

type lambdaPermissionState struct {
	Action              string `json:"action"`
	EventSourceToken    string `json:"event_source_token"`
	FunctionName        string `json:"function_name"`
	FunctionUrlAuthType string `json:"function_url_auth_type"`
	Id                  string `json:"id"`
	Principal           string `json:"principal"`
	PrincipalOrgId      string `json:"principal_org_id"`
	Qualifier           string `json:"qualifier"`
	SourceAccount       string `json:"source_account"`
	SourceArn           string `json:"source_arn"`
	StatementId         string `json:"statement_id"`
	StatementIdPrefix   string `json:"statement_id_prefix"`
}
