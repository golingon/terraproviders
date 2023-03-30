// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiGatewayBasePathMapping(name string, args ApiGatewayBasePathMappingArgs) *ApiGatewayBasePathMapping {
	return &ApiGatewayBasePathMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayBasePathMapping)(nil)

type ApiGatewayBasePathMapping struct {
	Name  string
	Args  ApiGatewayBasePathMappingArgs
	state *apiGatewayBasePathMappingState
}

func (agbpm *ApiGatewayBasePathMapping) Type() string {
	return "aws_api_gateway_base_path_mapping"
}

func (agbpm *ApiGatewayBasePathMapping) LocalName() string {
	return agbpm.Name
}

func (agbpm *ApiGatewayBasePathMapping) Configuration() interface{} {
	return agbpm.Args
}

func (agbpm *ApiGatewayBasePathMapping) Attributes() apiGatewayBasePathMappingAttributes {
	return apiGatewayBasePathMappingAttributes{ref: terra.ReferenceResource(agbpm)}
}

func (agbpm *ApiGatewayBasePathMapping) ImportState(av io.Reader) error {
	agbpm.state = &apiGatewayBasePathMappingState{}
	if err := json.NewDecoder(av).Decode(agbpm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agbpm.Type(), agbpm.LocalName(), err)
	}
	return nil
}

func (agbpm *ApiGatewayBasePathMapping) State() (*apiGatewayBasePathMappingState, bool) {
	return agbpm.state, agbpm.state != nil
}

func (agbpm *ApiGatewayBasePathMapping) StateMust() *apiGatewayBasePathMappingState {
	if agbpm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agbpm.Type(), agbpm.LocalName()))
	}
	return agbpm.state
}

func (agbpm *ApiGatewayBasePathMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(agbpm)
}

type ApiGatewayBasePathMappingArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// BasePath: string, optional
	BasePath terra.StringValue `hcl:"base_path,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StageName: string, optional
	StageName terra.StringValue `hcl:"stage_name,attr"`
	// DependsOn contains resources that ApiGatewayBasePathMapping depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiGatewayBasePathMappingAttributes struct {
	ref terra.Reference
}

func (agbpm apiGatewayBasePathMappingAttributes) ApiId() terra.StringValue {
	return terra.ReferenceString(agbpm.ref.Append("api_id"))
}

func (agbpm apiGatewayBasePathMappingAttributes) BasePath() terra.StringValue {
	return terra.ReferenceString(agbpm.ref.Append("base_path"))
}

func (agbpm apiGatewayBasePathMappingAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(agbpm.ref.Append("domain_name"))
}

func (agbpm apiGatewayBasePathMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(agbpm.ref.Append("id"))
}

func (agbpm apiGatewayBasePathMappingAttributes) StageName() terra.StringValue {
	return terra.ReferenceString(agbpm.ref.Append("stage_name"))
}

type apiGatewayBasePathMappingState struct {
	ApiId      string `json:"api_id"`
	BasePath   string `json:"base_path"`
	DomainName string `json:"domain_name"`
	Id         string `json:"id"`
	StageName  string `json:"stage_name"`
}
