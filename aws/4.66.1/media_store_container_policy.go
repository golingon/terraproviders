// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaStoreContainerPolicy creates a new instance of [MediaStoreContainerPolicy].
func NewMediaStoreContainerPolicy(name string, args MediaStoreContainerPolicyArgs) *MediaStoreContainerPolicy {
	return &MediaStoreContainerPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaStoreContainerPolicy)(nil)

// MediaStoreContainerPolicy represents the Terraform resource aws_media_store_container_policy.
type MediaStoreContainerPolicy struct {
	Name      string
	Args      MediaStoreContainerPolicyArgs
	state     *mediaStoreContainerPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) Type() string {
	return "aws_media_store_container_policy"
}

// LocalName returns the local name for [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) LocalName() string {
	return mscp.Name
}

// Configuration returns the configuration (args) for [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) Configuration() interface{} {
	return mscp.Args
}

// DependOn is used for other resources to depend on [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(mscp)
}

// Dependencies returns the list of resources [MediaStoreContainerPolicy] depends_on.
func (mscp *MediaStoreContainerPolicy) Dependencies() terra.Dependencies {
	return mscp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) LifecycleManagement() *terra.Lifecycle {
	return mscp.Lifecycle
}

// Attributes returns the attributes for [MediaStoreContainerPolicy].
func (mscp *MediaStoreContainerPolicy) Attributes() mediaStoreContainerPolicyAttributes {
	return mediaStoreContainerPolicyAttributes{ref: terra.ReferenceResource(mscp)}
}

// ImportState imports the given attribute values into [MediaStoreContainerPolicy]'s state.
func (mscp *MediaStoreContainerPolicy) ImportState(av io.Reader) error {
	mscp.state = &mediaStoreContainerPolicyState{}
	if err := json.NewDecoder(av).Decode(mscp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mscp.Type(), mscp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaStoreContainerPolicy] has state.
func (mscp *MediaStoreContainerPolicy) State() (*mediaStoreContainerPolicyState, bool) {
	return mscp.state, mscp.state != nil
}

// StateMust returns the state for [MediaStoreContainerPolicy]. Panics if the state is nil.
func (mscp *MediaStoreContainerPolicy) StateMust() *mediaStoreContainerPolicyState {
	if mscp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mscp.Type(), mscp.LocalName()))
	}
	return mscp.state
}

// MediaStoreContainerPolicyArgs contains the configurations for aws_media_store_container_policy.
type MediaStoreContainerPolicyArgs struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type mediaStoreContainerPolicyAttributes struct {
	ref terra.Reference
}

// ContainerName returns a reference to field container_name of aws_media_store_container_policy.
func (mscp mediaStoreContainerPolicyAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(mscp.ref.Append("container_name"))
}

// Id returns a reference to field id of aws_media_store_container_policy.
func (mscp mediaStoreContainerPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mscp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_media_store_container_policy.
func (mscp mediaStoreContainerPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(mscp.ref.Append("policy"))
}

type mediaStoreContainerPolicyState struct {
	ContainerName string `json:"container_name"`
	Id            string `json:"id"`
	Policy        string `json:"policy"`
}
