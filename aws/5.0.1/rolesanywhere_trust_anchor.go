// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rolesanywheretrustanchor "github.com/golingon/terraproviders/aws/5.0.1/rolesanywheretrustanchor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRolesanywhereTrustAnchor creates a new instance of [RolesanywhereTrustAnchor].
func NewRolesanywhereTrustAnchor(name string, args RolesanywhereTrustAnchorArgs) *RolesanywhereTrustAnchor {
	return &RolesanywhereTrustAnchor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RolesanywhereTrustAnchor)(nil)

// RolesanywhereTrustAnchor represents the Terraform resource aws_rolesanywhere_trust_anchor.
type RolesanywhereTrustAnchor struct {
	Name      string
	Args      RolesanywhereTrustAnchorArgs
	state     *rolesanywhereTrustAnchorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) Type() string {
	return "aws_rolesanywhere_trust_anchor"
}

// LocalName returns the local name for [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) LocalName() string {
	return rta.Name
}

// Configuration returns the configuration (args) for [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) Configuration() interface{} {
	return rta.Args
}

// DependOn is used for other resources to depend on [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) DependOn() terra.Reference {
	return terra.ReferenceResource(rta)
}

// Dependencies returns the list of resources [RolesanywhereTrustAnchor] depends_on.
func (rta *RolesanywhereTrustAnchor) Dependencies() terra.Dependencies {
	return rta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) LifecycleManagement() *terra.Lifecycle {
	return rta.Lifecycle
}

// Attributes returns the attributes for [RolesanywhereTrustAnchor].
func (rta *RolesanywhereTrustAnchor) Attributes() rolesanywhereTrustAnchorAttributes {
	return rolesanywhereTrustAnchorAttributes{ref: terra.ReferenceResource(rta)}
}

// ImportState imports the given attribute values into [RolesanywhereTrustAnchor]'s state.
func (rta *RolesanywhereTrustAnchor) ImportState(av io.Reader) error {
	rta.state = &rolesanywhereTrustAnchorState{}
	if err := json.NewDecoder(av).Decode(rta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rta.Type(), rta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RolesanywhereTrustAnchor] has state.
func (rta *RolesanywhereTrustAnchor) State() (*rolesanywhereTrustAnchorState, bool) {
	return rta.state, rta.state != nil
}

// StateMust returns the state for [RolesanywhereTrustAnchor]. Panics if the state is nil.
func (rta *RolesanywhereTrustAnchor) StateMust() *rolesanywhereTrustAnchorState {
	if rta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rta.Type(), rta.LocalName()))
	}
	return rta.state
}

// RolesanywhereTrustAnchorArgs contains the configurations for aws_rolesanywhere_trust_anchor.
type RolesanywhereTrustAnchorArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Source: required
	Source *rolesanywheretrustanchor.Source `hcl:"source,block" validate:"required"`
}
type rolesanywhereTrustAnchorAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rta.ref.Append("arn"))
}

// Enabled returns a reference to field enabled of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rta.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rta.ref.Append("id"))
}

// Name returns a reference to field name of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rta.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rta.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rolesanywhere_trust_anchor.
func (rta rolesanywhereTrustAnchorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rta.ref.Append("tags_all"))
}

func (rta rolesanywhereTrustAnchorAttributes) Source() terra.ListValue[rolesanywheretrustanchor.SourceAttributes] {
	return terra.ReferenceAsList[rolesanywheretrustanchor.SourceAttributes](rta.ref.Append("source"))
}

type rolesanywhereTrustAnchorState struct {
	Arn     string                                 `json:"arn"`
	Enabled bool                                   `json:"enabled"`
	Id      string                                 `json:"id"`
	Name    string                                 `json:"name"`
	Tags    map[string]string                      `json:"tags"`
	TagsAll map[string]string                      `json:"tags_all"`
	Source  []rolesanywheretrustanchor.SourceState `json:"source"`
}
