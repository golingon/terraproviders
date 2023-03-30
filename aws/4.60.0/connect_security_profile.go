// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewConnectSecurityProfile(name string, args ConnectSecurityProfileArgs) *ConnectSecurityProfile {
	return &ConnectSecurityProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectSecurityProfile)(nil)

type ConnectSecurityProfile struct {
	Name  string
	Args  ConnectSecurityProfileArgs
	state *connectSecurityProfileState
}

func (csp *ConnectSecurityProfile) Type() string {
	return "aws_connect_security_profile"
}

func (csp *ConnectSecurityProfile) LocalName() string {
	return csp.Name
}

func (csp *ConnectSecurityProfile) Configuration() interface{} {
	return csp.Args
}

func (csp *ConnectSecurityProfile) Attributes() connectSecurityProfileAttributes {
	return connectSecurityProfileAttributes{ref: terra.ReferenceResource(csp)}
}

func (csp *ConnectSecurityProfile) ImportState(av io.Reader) error {
	csp.state = &connectSecurityProfileState{}
	if err := json.NewDecoder(av).Decode(csp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csp.Type(), csp.LocalName(), err)
	}
	return nil
}

func (csp *ConnectSecurityProfile) State() (*connectSecurityProfileState, bool) {
	return csp.state, csp.state != nil
}

func (csp *ConnectSecurityProfile) StateMust() *connectSecurityProfileState {
	if csp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csp.Type(), csp.LocalName()))
	}
	return csp.state
}

func (csp *ConnectSecurityProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(csp)
}

type ConnectSecurityProfileArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Permissions: set of string, optional
	Permissions terra.SetValue[terra.StringValue] `hcl:"permissions,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that ConnectSecurityProfile depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type connectSecurityProfileAttributes struct {
	ref terra.Reference
}

func (csp connectSecurityProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("arn"))
}

func (csp connectSecurityProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("description"))
}

func (csp connectSecurityProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("id"))
}

func (csp connectSecurityProfileAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("instance_id"))
}

func (csp connectSecurityProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("name"))
}

func (csp connectSecurityProfileAttributes) OrganizationResourceId() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("organization_resource_id"))
}

func (csp connectSecurityProfileAttributes) Permissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](csp.ref.Append("permissions"))
}

func (csp connectSecurityProfileAttributes) SecurityProfileId() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("security_profile_id"))
}

func (csp connectSecurityProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](csp.ref.Append("tags"))
}

func (csp connectSecurityProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](csp.ref.Append("tags_all"))
}

type connectSecurityProfileState struct {
	Arn                    string            `json:"arn"`
	Description            string            `json:"description"`
	Id                     string            `json:"id"`
	InstanceId             string            `json:"instance_id"`
	Name                   string            `json:"name"`
	OrganizationResourceId string            `json:"organization_resource_id"`
	Permissions            []string          `json:"permissions"`
	SecurityProfileId      string            `json:"security_profile_id"`
	Tags                   map[string]string `json:"tags"`
	TagsAll                map[string]string `json:"tags_all"`
}
