// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	eksidentityproviderconfig "github.com/golingon/terraproviders/aws/5.12.0/eksidentityproviderconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEksIdentityProviderConfig creates a new instance of [EksIdentityProviderConfig].
func NewEksIdentityProviderConfig(name string, args EksIdentityProviderConfigArgs) *EksIdentityProviderConfig {
	return &EksIdentityProviderConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EksIdentityProviderConfig)(nil)

// EksIdentityProviderConfig represents the Terraform resource aws_eks_identity_provider_config.
type EksIdentityProviderConfig struct {
	Name      string
	Args      EksIdentityProviderConfigArgs
	state     *eksIdentityProviderConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) Type() string {
	return "aws_eks_identity_provider_config"
}

// LocalName returns the local name for [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) LocalName() string {
	return eipc.Name
}

// Configuration returns the configuration (args) for [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) Configuration() interface{} {
	return eipc.Args
}

// DependOn is used for other resources to depend on [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(eipc)
}

// Dependencies returns the list of resources [EksIdentityProviderConfig] depends_on.
func (eipc *EksIdentityProviderConfig) Dependencies() terra.Dependencies {
	return eipc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) LifecycleManagement() *terra.Lifecycle {
	return eipc.Lifecycle
}

// Attributes returns the attributes for [EksIdentityProviderConfig].
func (eipc *EksIdentityProviderConfig) Attributes() eksIdentityProviderConfigAttributes {
	return eksIdentityProviderConfigAttributes{ref: terra.ReferenceResource(eipc)}
}

// ImportState imports the given attribute values into [EksIdentityProviderConfig]'s state.
func (eipc *EksIdentityProviderConfig) ImportState(av io.Reader) error {
	eipc.state = &eksIdentityProviderConfigState{}
	if err := json.NewDecoder(av).Decode(eipc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eipc.Type(), eipc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EksIdentityProviderConfig] has state.
func (eipc *EksIdentityProviderConfig) State() (*eksIdentityProviderConfigState, bool) {
	return eipc.state, eipc.state != nil
}

// StateMust returns the state for [EksIdentityProviderConfig]. Panics if the state is nil.
func (eipc *EksIdentityProviderConfig) StateMust() *eksIdentityProviderConfigState {
	if eipc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eipc.Type(), eipc.LocalName()))
	}
	return eipc.state
}

// EksIdentityProviderConfigArgs contains the configurations for aws_eks_identity_provider_config.
type EksIdentityProviderConfigArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Oidc: required
	Oidc *eksidentityproviderconfig.Oidc `hcl:"oidc,block" validate:"required"`
	// Timeouts: optional
	Timeouts *eksidentityproviderconfig.Timeouts `hcl:"timeouts,block"`
}
type eksIdentityProviderConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eipc.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(eipc.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eipc.ref.Append("id"))
}

// Status returns a reference to field status of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(eipc.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eipc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eks_identity_provider_config.
func (eipc eksIdentityProviderConfigAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eipc.ref.Append("tags_all"))
}

func (eipc eksIdentityProviderConfigAttributes) Oidc() terra.ListValue[eksidentityproviderconfig.OidcAttributes] {
	return terra.ReferenceAsList[eksidentityproviderconfig.OidcAttributes](eipc.ref.Append("oidc"))
}

func (eipc eksIdentityProviderConfigAttributes) Timeouts() eksidentityproviderconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eksidentityproviderconfig.TimeoutsAttributes](eipc.ref.Append("timeouts"))
}

type eksIdentityProviderConfigState struct {
	Arn         string                                   `json:"arn"`
	ClusterName string                                   `json:"cluster_name"`
	Id          string                                   `json:"id"`
	Status      string                                   `json:"status"`
	Tags        map[string]string                        `json:"tags"`
	TagsAll     map[string]string                        `json:"tags_all"`
	Oidc        []eksidentityproviderconfig.OidcState    `json:"oidc"`
	Timeouts    *eksidentityproviderconfig.TimeoutsState `json:"timeouts"`
}
