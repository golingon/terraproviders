// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_wafv2_ip_set

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_wafv2_ip_set.
type Resource struct {
	Name      string
	Args      Args
	state     *awsWafv2IpSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awis *Resource) Type() string {
	return "aws_wafv2_ip_set"
}

// LocalName returns the local name for [Resource].
func (awis *Resource) LocalName() string {
	return awis.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awis *Resource) Configuration() interface{} {
	return awis.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awis *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awis)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awis *Resource) Dependencies() terra.Dependencies {
	return awis.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awis *Resource) LifecycleManagement() *terra.Lifecycle {
	return awis.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awis *Resource) Attributes() awsWafv2IpSetAttributes {
	return awsWafv2IpSetAttributes{ref: terra.ReferenceResource(awis)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awis *Resource) ImportState(state io.Reader) error {
	awis.state = &awsWafv2IpSetState{}
	if err := json.NewDecoder(state).Decode(awis.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awis.Type(), awis.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awis *Resource) State() (*awsWafv2IpSetState, bool) {
	return awis.state, awis.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awis *Resource) StateMust() *awsWafv2IpSetState {
	if awis.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awis.Type(), awis.LocalName()))
	}
	return awis.state
}

// Args contains the configurations for aws_wafv2_ip_set.
type Args struct {
	// Addresses: set of string, optional
	Addresses terra.SetValue[terra.StringValue] `hcl:"addresses,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddressVersion: string, required
	IpAddressVersion terra.StringValue `hcl:"ip_address_version,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsWafv2IpSetAttributes struct {
	ref terra.Reference
}

// Addresses returns a reference to field addresses of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Addresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](awis.ref.Append("addresses"))
}

// Arn returns a reference to field arn of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("arn"))
}

// Description returns a reference to field description of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("id"))
}

// IpAddressVersion returns a reference to field ip_address_version of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) IpAddressVersion() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("ip_address_version"))
}

// LockToken returns a reference to field lock_token of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) LockToken() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("lock_token"))
}

// Name returns a reference to field name of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(awis.ref.Append("scope"))
}

// Tags returns a reference to field tags of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awis.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafv2_ip_set.
func (awis awsWafv2IpSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awis.ref.Append("tags_all"))
}

type awsWafv2IpSetState struct {
	Addresses        []string          `json:"addresses"`
	Arn              string            `json:"arn"`
	Description      string            `json:"description"`
	Id               string            `json:"id"`
	IpAddressVersion string            `json:"ip_address_version"`
	LockToken        string            `json:"lock_token"`
	Name             string            `json:"name"`
	Scope            string            `json:"scope"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
}
