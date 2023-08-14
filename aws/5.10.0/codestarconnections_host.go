// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codestarconnectionshost "github.com/golingon/terraproviders/aws/5.10.0/codestarconnectionshost"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodestarconnectionsHost creates a new instance of [CodestarconnectionsHost].
func NewCodestarconnectionsHost(name string, args CodestarconnectionsHostArgs) *CodestarconnectionsHost {
	return &CodestarconnectionsHost{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodestarconnectionsHost)(nil)

// CodestarconnectionsHost represents the Terraform resource aws_codestarconnections_host.
type CodestarconnectionsHost struct {
	Name      string
	Args      CodestarconnectionsHostArgs
	state     *codestarconnectionsHostState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) Type() string {
	return "aws_codestarconnections_host"
}

// LocalName returns the local name for [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) LocalName() string {
	return ch.Name
}

// Configuration returns the configuration (args) for [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) Configuration() interface{} {
	return ch.Args
}

// DependOn is used for other resources to depend on [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) DependOn() terra.Reference {
	return terra.ReferenceResource(ch)
}

// Dependencies returns the list of resources [CodestarconnectionsHost] depends_on.
func (ch *CodestarconnectionsHost) Dependencies() terra.Dependencies {
	return ch.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) LifecycleManagement() *terra.Lifecycle {
	return ch.Lifecycle
}

// Attributes returns the attributes for [CodestarconnectionsHost].
func (ch *CodestarconnectionsHost) Attributes() codestarconnectionsHostAttributes {
	return codestarconnectionsHostAttributes{ref: terra.ReferenceResource(ch)}
}

// ImportState imports the given attribute values into [CodestarconnectionsHost]'s state.
func (ch *CodestarconnectionsHost) ImportState(av io.Reader) error {
	ch.state = &codestarconnectionsHostState{}
	if err := json.NewDecoder(av).Decode(ch.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ch.Type(), ch.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodestarconnectionsHost] has state.
func (ch *CodestarconnectionsHost) State() (*codestarconnectionsHostState, bool) {
	return ch.state, ch.state != nil
}

// StateMust returns the state for [CodestarconnectionsHost]. Panics if the state is nil.
func (ch *CodestarconnectionsHost) StateMust() *codestarconnectionsHostState {
	if ch.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ch.Type(), ch.LocalName()))
	}
	return ch.state
}

// CodestarconnectionsHostArgs contains the configurations for aws_codestarconnections_host.
type CodestarconnectionsHostArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProviderEndpoint: string, required
	ProviderEndpoint terra.StringValue `hcl:"provider_endpoint,attr" validate:"required"`
	// ProviderType: string, required
	ProviderType terra.StringValue `hcl:"provider_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *codestarconnectionshost.Timeouts `hcl:"timeouts,block"`
	// VpcConfiguration: optional
	VpcConfiguration *codestarconnectionshost.VpcConfiguration `hcl:"vpc_configuration,block"`
}
type codestarconnectionsHostAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("arn"))
}

// Id returns a reference to field id of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("id"))
}

// Name returns a reference to field name of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("name"))
}

// ProviderEndpoint returns a reference to field provider_endpoint of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) ProviderEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("provider_endpoint"))
}

// ProviderType returns a reference to field provider_type of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) ProviderType() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("provider_type"))
}

// Status returns a reference to field status of aws_codestarconnections_host.
func (ch codestarconnectionsHostAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("status"))
}

func (ch codestarconnectionsHostAttributes) Timeouts() codestarconnectionshost.TimeoutsAttributes {
	return terra.ReferenceAsSingle[codestarconnectionshost.TimeoutsAttributes](ch.ref.Append("timeouts"))
}

func (ch codestarconnectionsHostAttributes) VpcConfiguration() terra.ListValue[codestarconnectionshost.VpcConfigurationAttributes] {
	return terra.ReferenceAsList[codestarconnectionshost.VpcConfigurationAttributes](ch.ref.Append("vpc_configuration"))
}

type codestarconnectionsHostState struct {
	Arn              string                                          `json:"arn"`
	Id               string                                          `json:"id"`
	Name             string                                          `json:"name"`
	ProviderEndpoint string                                          `json:"provider_endpoint"`
	ProviderType     string                                          `json:"provider_type"`
	Status           string                                          `json:"status"`
	Timeouts         *codestarconnectionshost.TimeoutsState          `json:"timeouts"`
	VpcConfiguration []codestarconnectionshost.VpcConfigurationState `json:"vpc_configuration"`
}