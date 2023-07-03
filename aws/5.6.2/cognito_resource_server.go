// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitoresourceserver "github.com/golingon/terraproviders/aws/5.6.2/cognitoresourceserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoResourceServer creates a new instance of [CognitoResourceServer].
func NewCognitoResourceServer(name string, args CognitoResourceServerArgs) *CognitoResourceServer {
	return &CognitoResourceServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoResourceServer)(nil)

// CognitoResourceServer represents the Terraform resource aws_cognito_resource_server.
type CognitoResourceServer struct {
	Name      string
	Args      CognitoResourceServerArgs
	state     *cognitoResourceServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoResourceServer].
func (crs *CognitoResourceServer) Type() string {
	return "aws_cognito_resource_server"
}

// LocalName returns the local name for [CognitoResourceServer].
func (crs *CognitoResourceServer) LocalName() string {
	return crs.Name
}

// Configuration returns the configuration (args) for [CognitoResourceServer].
func (crs *CognitoResourceServer) Configuration() interface{} {
	return crs.Args
}

// DependOn is used for other resources to depend on [CognitoResourceServer].
func (crs *CognitoResourceServer) DependOn() terra.Reference {
	return terra.ReferenceResource(crs)
}

// Dependencies returns the list of resources [CognitoResourceServer] depends_on.
func (crs *CognitoResourceServer) Dependencies() terra.Dependencies {
	return crs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoResourceServer].
func (crs *CognitoResourceServer) LifecycleManagement() *terra.Lifecycle {
	return crs.Lifecycle
}

// Attributes returns the attributes for [CognitoResourceServer].
func (crs *CognitoResourceServer) Attributes() cognitoResourceServerAttributes {
	return cognitoResourceServerAttributes{ref: terra.ReferenceResource(crs)}
}

// ImportState imports the given attribute values into [CognitoResourceServer]'s state.
func (crs *CognitoResourceServer) ImportState(av io.Reader) error {
	crs.state = &cognitoResourceServerState{}
	if err := json.NewDecoder(av).Decode(crs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crs.Type(), crs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoResourceServer] has state.
func (crs *CognitoResourceServer) State() (*cognitoResourceServerState, bool) {
	return crs.state, crs.state != nil
}

// StateMust returns the state for [CognitoResourceServer]. Panics if the state is nil.
func (crs *CognitoResourceServer) StateMust() *cognitoResourceServerState {
	if crs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crs.Type(), crs.LocalName()))
	}
	return crs.state
}

// CognitoResourceServerArgs contains the configurations for aws_cognito_resource_server.
type CognitoResourceServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identifier: string, required
	Identifier terra.StringValue `hcl:"identifier,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
	// Scope: min=0,max=100
	Scope []cognitoresourceserver.Scope `hcl:"scope,block" validate:"min=0,max=100"`
}
type cognitoResourceServerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cognito_resource_server.
func (crs cognitoResourceServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_cognito_resource_server.
func (crs cognitoResourceServerAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("identifier"))
}

// Name returns a reference to field name of aws_cognito_resource_server.
func (crs cognitoResourceServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("name"))
}

// ScopeIdentifiers returns a reference to field scope_identifiers of aws_cognito_resource_server.
func (crs cognitoResourceServerAttributes) ScopeIdentifiers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crs.ref.Append("scope_identifiers"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_resource_server.
func (crs cognitoResourceServerAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("user_pool_id"))
}

func (crs cognitoResourceServerAttributes) Scope() terra.SetValue[cognitoresourceserver.ScopeAttributes] {
	return terra.ReferenceAsSet[cognitoresourceserver.ScopeAttributes](crs.ref.Append("scope"))
}

type cognitoResourceServerState struct {
	Id               string                             `json:"id"`
	Identifier       string                             `json:"identifier"`
	Name             string                             `json:"name"`
	ScopeIdentifiers []string                           `json:"scope_identifiers"`
	UserPoolId       string                             `json:"user_pool_id"`
	Scope            []cognitoresourceserver.ScopeState `json:"scope"`
}
