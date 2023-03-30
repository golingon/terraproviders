// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLambdaLayerVersionPermission(name string, args LambdaLayerVersionPermissionArgs) *LambdaLayerVersionPermission {
	return &LambdaLayerVersionPermission{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaLayerVersionPermission)(nil)

type LambdaLayerVersionPermission struct {
	Name  string
	Args  LambdaLayerVersionPermissionArgs
	state *lambdaLayerVersionPermissionState
}

func (llvp *LambdaLayerVersionPermission) Type() string {
	return "aws_lambda_layer_version_permission"
}

func (llvp *LambdaLayerVersionPermission) LocalName() string {
	return llvp.Name
}

func (llvp *LambdaLayerVersionPermission) Configuration() interface{} {
	return llvp.Args
}

func (llvp *LambdaLayerVersionPermission) Attributes() lambdaLayerVersionPermissionAttributes {
	return lambdaLayerVersionPermissionAttributes{ref: terra.ReferenceResource(llvp)}
}

func (llvp *LambdaLayerVersionPermission) ImportState(av io.Reader) error {
	llvp.state = &lambdaLayerVersionPermissionState{}
	if err := json.NewDecoder(av).Decode(llvp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llvp.Type(), llvp.LocalName(), err)
	}
	return nil
}

func (llvp *LambdaLayerVersionPermission) State() (*lambdaLayerVersionPermissionState, bool) {
	return llvp.state, llvp.state != nil
}

func (llvp *LambdaLayerVersionPermission) StateMust() *lambdaLayerVersionPermissionState {
	if llvp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llvp.Type(), llvp.LocalName()))
	}
	return llvp.state
}

func (llvp *LambdaLayerVersionPermission) DependOn() terra.Reference {
	return terra.ReferenceResource(llvp)
}

type LambdaLayerVersionPermissionArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LayerName: string, required
	LayerName terra.StringValue `hcl:"layer_name,attr" validate:"required"`
	// OrganizationId: string, optional
	OrganizationId terra.StringValue `hcl:"organization_id,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// StatementId: string, required
	StatementId terra.StringValue `hcl:"statement_id,attr" validate:"required"`
	// VersionNumber: number, required
	VersionNumber terra.NumberValue `hcl:"version_number,attr" validate:"required"`
	// DependsOn contains resources that LambdaLayerVersionPermission depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lambdaLayerVersionPermissionAttributes struct {
	ref terra.Reference
}

func (llvp lambdaLayerVersionPermissionAttributes) Action() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("action"))
}

func (llvp lambdaLayerVersionPermissionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("id"))
}

func (llvp lambdaLayerVersionPermissionAttributes) LayerName() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("layer_name"))
}

func (llvp lambdaLayerVersionPermissionAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("organization_id"))
}

func (llvp lambdaLayerVersionPermissionAttributes) Policy() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("policy"))
}

func (llvp lambdaLayerVersionPermissionAttributes) Principal() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("principal"))
}

func (llvp lambdaLayerVersionPermissionAttributes) RevisionId() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("revision_id"))
}

func (llvp lambdaLayerVersionPermissionAttributes) StatementId() terra.StringValue {
	return terra.ReferenceString(llvp.ref.Append("statement_id"))
}

func (llvp lambdaLayerVersionPermissionAttributes) VersionNumber() terra.NumberValue {
	return terra.ReferenceNumber(llvp.ref.Append("version_number"))
}

type lambdaLayerVersionPermissionState struct {
	Action         string  `json:"action"`
	Id             string  `json:"id"`
	LayerName      string  `json:"layer_name"`
	OrganizationId string  `json:"organization_id"`
	Policy         string  `json:"policy"`
	Principal      string  `json:"principal"`
	RevisionId     string  `json:"revision_id"`
	StatementId    string  `json:"statement_id"`
	VersionNumber  float64 `json:"version_number"`
}
