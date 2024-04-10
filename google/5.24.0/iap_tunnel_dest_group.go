// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	iaptunneldestgroup "github.com/golingon/terraproviders/google/5.24.0/iaptunneldestgroup"
	"io"
)

// NewIapTunnelDestGroup creates a new instance of [IapTunnelDestGroup].
func NewIapTunnelDestGroup(name string, args IapTunnelDestGroupArgs) *IapTunnelDestGroup {
	return &IapTunnelDestGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelDestGroup)(nil)

// IapTunnelDestGroup represents the Terraform resource google_iap_tunnel_dest_group.
type IapTunnelDestGroup struct {
	Name      string
	Args      IapTunnelDestGroupArgs
	state     *iapTunnelDestGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) Type() string {
	return "google_iap_tunnel_dest_group"
}

// LocalName returns the local name for [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) LocalName() string {
	return itdg.Name
}

// Configuration returns the configuration (args) for [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) Configuration() interface{} {
	return itdg.Args
}

// DependOn is used for other resources to depend on [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(itdg)
}

// Dependencies returns the list of resources [IapTunnelDestGroup] depends_on.
func (itdg *IapTunnelDestGroup) Dependencies() terra.Dependencies {
	return itdg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) LifecycleManagement() *terra.Lifecycle {
	return itdg.Lifecycle
}

// Attributes returns the attributes for [IapTunnelDestGroup].
func (itdg *IapTunnelDestGroup) Attributes() iapTunnelDestGroupAttributes {
	return iapTunnelDestGroupAttributes{ref: terra.ReferenceResource(itdg)}
}

// ImportState imports the given attribute values into [IapTunnelDestGroup]'s state.
func (itdg *IapTunnelDestGroup) ImportState(av io.Reader) error {
	itdg.state = &iapTunnelDestGroupState{}
	if err := json.NewDecoder(av).Decode(itdg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itdg.Type(), itdg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelDestGroup] has state.
func (itdg *IapTunnelDestGroup) State() (*iapTunnelDestGroupState, bool) {
	return itdg.state, itdg.state != nil
}

// StateMust returns the state for [IapTunnelDestGroup]. Panics if the state is nil.
func (itdg *IapTunnelDestGroup) StateMust() *iapTunnelDestGroupState {
	if itdg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itdg.Type(), itdg.LocalName()))
	}
	return itdg.state
}

// IapTunnelDestGroupArgs contains the configurations for google_iap_tunnel_dest_group.
type IapTunnelDestGroupArgs struct {
	// Cidrs: list of string, optional
	Cidrs terra.ListValue[terra.StringValue] `hcl:"cidrs,attr"`
	// Fqdns: list of string, optional
	Fqdns terra.ListValue[terra.StringValue] `hcl:"fqdns,attr"`
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *iaptunneldestgroup.Timeouts `hcl:"timeouts,block"`
}
type iapTunnelDestGroupAttributes struct {
	ref terra.Reference
}

// Cidrs returns a reference to field cidrs of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Cidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](itdg.ref.Append("cidrs"))
}

// Fqdns returns a reference to field fqdns of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](itdg.ref.Append("fqdns"))
}

// GroupName returns a reference to field group_name of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(itdg.ref.Append("group_name"))
}

// Id returns a reference to field id of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itdg.ref.Append("id"))
}

// Name returns a reference to field name of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itdg.ref.Append("name"))
}

// Project returns a reference to field project of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itdg.ref.Append("project"))
}

// Region returns a reference to field region of google_iap_tunnel_dest_group.
func (itdg iapTunnelDestGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(itdg.ref.Append("region"))
}

func (itdg iapTunnelDestGroupAttributes) Timeouts() iaptunneldestgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iaptunneldestgroup.TimeoutsAttributes](itdg.ref.Append("timeouts"))
}

type iapTunnelDestGroupState struct {
	Cidrs     []string                          `json:"cidrs"`
	Fqdns     []string                          `json:"fqdns"`
	GroupName string                            `json:"group_name"`
	Id        string                            `json:"id"`
	Name      string                            `json:"name"`
	Project   string                            `json:"project"`
	Region    string                            `json:"region"`
	Timeouts  *iaptunneldestgroup.TimeoutsState `json:"timeouts"`
}
