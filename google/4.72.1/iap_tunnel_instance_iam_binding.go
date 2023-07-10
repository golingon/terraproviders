// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iaptunnelinstanceiambinding "github.com/golingon/terraproviders/google/4.72.1/iaptunnelinstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapTunnelInstanceIamBinding creates a new instance of [IapTunnelInstanceIamBinding].
func NewIapTunnelInstanceIamBinding(name string, args IapTunnelInstanceIamBindingArgs) *IapTunnelInstanceIamBinding {
	return &IapTunnelInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelInstanceIamBinding)(nil)

// IapTunnelInstanceIamBinding represents the Terraform resource google_iap_tunnel_instance_iam_binding.
type IapTunnelInstanceIamBinding struct {
	Name      string
	Args      IapTunnelInstanceIamBindingArgs
	state     *iapTunnelInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) Type() string {
	return "google_iap_tunnel_instance_iam_binding"
}

// LocalName returns the local name for [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) LocalName() string {
	return itiib.Name
}

// Configuration returns the configuration (args) for [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) Configuration() interface{} {
	return itiib.Args
}

// DependOn is used for other resources to depend on [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(itiib)
}

// Dependencies returns the list of resources [IapTunnelInstanceIamBinding] depends_on.
func (itiib *IapTunnelInstanceIamBinding) Dependencies() terra.Dependencies {
	return itiib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return itiib.Lifecycle
}

// Attributes returns the attributes for [IapTunnelInstanceIamBinding].
func (itiib *IapTunnelInstanceIamBinding) Attributes() iapTunnelInstanceIamBindingAttributes {
	return iapTunnelInstanceIamBindingAttributes{ref: terra.ReferenceResource(itiib)}
}

// ImportState imports the given attribute values into [IapTunnelInstanceIamBinding]'s state.
func (itiib *IapTunnelInstanceIamBinding) ImportState(av io.Reader) error {
	itiib.state = &iapTunnelInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(itiib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itiib.Type(), itiib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelInstanceIamBinding] has state.
func (itiib *IapTunnelInstanceIamBinding) State() (*iapTunnelInstanceIamBindingState, bool) {
	return itiib.state, itiib.state != nil
}

// StateMust returns the state for [IapTunnelInstanceIamBinding]. Panics if the state is nil.
func (itiib *IapTunnelInstanceIamBinding) StateMust() *iapTunnelInstanceIamBindingState {
	if itiib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itiib.Type(), itiib.LocalName()))
	}
	return itiib.state
}

// IapTunnelInstanceIamBindingArgs contains the configurations for google_iap_tunnel_instance_iam_binding.
type IapTunnelInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Condition: optional
	Condition *iaptunnelinstanceiambinding.Condition `hcl:"condition,block"`
}
type iapTunnelInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("instance"))
}

// Members returns a reference to field members of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](itiib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("role"))
}

// Zone returns a reference to field zone of google_iap_tunnel_instance_iam_binding.
func (itiib iapTunnelInstanceIamBindingAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(itiib.ref.Append("zone"))
}

func (itiib iapTunnelInstanceIamBindingAttributes) Condition() terra.ListValue[iaptunnelinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iaptunnelinstanceiambinding.ConditionAttributes](itiib.ref.Append("condition"))
}

type iapTunnelInstanceIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Instance  string                                       `json:"instance"`
	Members   []string                                     `json:"members"`
	Project   string                                       `json:"project"`
	Role      string                                       `json:"role"`
	Zone      string                                       `json:"zone"`
	Condition []iaptunnelinstanceiambinding.ConditionState `json:"condition"`
}
