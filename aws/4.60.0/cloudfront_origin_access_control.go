// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontOriginAccessControl creates a new instance of [CloudfrontOriginAccessControl].
func NewCloudfrontOriginAccessControl(name string, args CloudfrontOriginAccessControlArgs) *CloudfrontOriginAccessControl {
	return &CloudfrontOriginAccessControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontOriginAccessControl)(nil)

// CloudfrontOriginAccessControl represents the Terraform resource aws_cloudfront_origin_access_control.
type CloudfrontOriginAccessControl struct {
	Name      string
	Args      CloudfrontOriginAccessControlArgs
	state     *cloudfrontOriginAccessControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) Type() string {
	return "aws_cloudfront_origin_access_control"
}

// LocalName returns the local name for [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) LocalName() string {
	return coac.Name
}

// Configuration returns the configuration (args) for [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) Configuration() interface{} {
	return coac.Args
}

// DependOn is used for other resources to depend on [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) DependOn() terra.Reference {
	return terra.ReferenceResource(coac)
}

// Dependencies returns the list of resources [CloudfrontOriginAccessControl] depends_on.
func (coac *CloudfrontOriginAccessControl) Dependencies() terra.Dependencies {
	return coac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) LifecycleManagement() *terra.Lifecycle {
	return coac.Lifecycle
}

// Attributes returns the attributes for [CloudfrontOriginAccessControl].
func (coac *CloudfrontOriginAccessControl) Attributes() cloudfrontOriginAccessControlAttributes {
	return cloudfrontOriginAccessControlAttributes{ref: terra.ReferenceResource(coac)}
}

// ImportState imports the given attribute values into [CloudfrontOriginAccessControl]'s state.
func (coac *CloudfrontOriginAccessControl) ImportState(av io.Reader) error {
	coac.state = &cloudfrontOriginAccessControlState{}
	if err := json.NewDecoder(av).Decode(coac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", coac.Type(), coac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontOriginAccessControl] has state.
func (coac *CloudfrontOriginAccessControl) State() (*cloudfrontOriginAccessControlState, bool) {
	return coac.state, coac.state != nil
}

// StateMust returns the state for [CloudfrontOriginAccessControl]. Panics if the state is nil.
func (coac *CloudfrontOriginAccessControl) StateMust() *cloudfrontOriginAccessControlState {
	if coac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", coac.Type(), coac.LocalName()))
	}
	return coac.state
}

// CloudfrontOriginAccessControlArgs contains the configurations for aws_cloudfront_origin_access_control.
type CloudfrontOriginAccessControlArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OriginAccessControlOriginType: string, required
	OriginAccessControlOriginType terra.StringValue `hcl:"origin_access_control_origin_type,attr" validate:"required"`
	// SigningBehavior: string, required
	SigningBehavior terra.StringValue `hcl:"signing_behavior,attr" validate:"required"`
	// SigningProtocol: string, required
	SigningProtocol terra.StringValue `hcl:"signing_protocol,attr" validate:"required"`
}
type cloudfrontOriginAccessControlAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("description"))
}

// Etag returns a reference to field etag of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("name"))
}

// OriginAccessControlOriginType returns a reference to field origin_access_control_origin_type of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) OriginAccessControlOriginType() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("origin_access_control_origin_type"))
}

// SigningBehavior returns a reference to field signing_behavior of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) SigningBehavior() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("signing_behavior"))
}

// SigningProtocol returns a reference to field signing_protocol of aws_cloudfront_origin_access_control.
func (coac cloudfrontOriginAccessControlAttributes) SigningProtocol() terra.StringValue {
	return terra.ReferenceAsString(coac.ref.Append("signing_protocol"))
}

type cloudfrontOriginAccessControlState struct {
	Description                   string `json:"description"`
	Etag                          string `json:"etag"`
	Id                            string `json:"id"`
	Name                          string `json:"name"`
	OriginAccessControlOriginType string `json:"origin_access_control_origin_type"`
	SigningBehavior               string `json:"signing_behavior"`
	SigningProtocol               string `json:"signing_protocol"`
}
