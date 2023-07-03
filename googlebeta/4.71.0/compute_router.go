// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computerouter "github.com/golingon/terraproviders/googlebeta/4.71.0/computerouter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRouter creates a new instance of [ComputeRouter].
func NewComputeRouter(name string, args ComputeRouterArgs) *ComputeRouter {
	return &ComputeRouter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRouter)(nil)

// ComputeRouter represents the Terraform resource google_compute_router.
type ComputeRouter struct {
	Name      string
	Args      ComputeRouterArgs
	state     *computeRouterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRouter].
func (cr *ComputeRouter) Type() string {
	return "google_compute_router"
}

// LocalName returns the local name for [ComputeRouter].
func (cr *ComputeRouter) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [ComputeRouter].
func (cr *ComputeRouter) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [ComputeRouter].
func (cr *ComputeRouter) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [ComputeRouter] depends_on.
func (cr *ComputeRouter) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRouter].
func (cr *ComputeRouter) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [ComputeRouter].
func (cr *ComputeRouter) Attributes() computeRouterAttributes {
	return computeRouterAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [ComputeRouter]'s state.
func (cr *ComputeRouter) ImportState(av io.Reader) error {
	cr.state = &computeRouterState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRouter] has state.
func (cr *ComputeRouter) State() (*computeRouterState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [ComputeRouter]. Panics if the state is nil.
func (cr *ComputeRouter) StateMust() *computeRouterState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// ComputeRouterArgs contains the configurations for google_compute_router.
type ComputeRouterArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EncryptedInterconnectRouter: bool, optional
	EncryptedInterconnectRouter terra.BoolValue `hcl:"encrypted_interconnect_router,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Bgp: optional
	Bgp *computerouter.Bgp `hcl:"bgp,block"`
	// Timeouts: optional
	Timeouts *computerouter.Timeouts `hcl:"timeouts,block"`
}
type computeRouterAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_router.
func (cr computeRouterAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_router.
func (cr computeRouterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// EncryptedInterconnectRouter returns a reference to field encrypted_interconnect_router of google_compute_router.
func (cr computeRouterAttributes) EncryptedInterconnectRouter() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("encrypted_interconnect_router"))
}

// Id returns a reference to field id of google_compute_router.
func (cr computeRouterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_router.
func (cr computeRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_router.
func (cr computeRouterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_router.
func (cr computeRouterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router.
func (cr computeRouterAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_router.
func (cr computeRouterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("self_link"))
}

func (cr computeRouterAttributes) Bgp() terra.ListValue[computerouter.BgpAttributes] {
	return terra.ReferenceAsList[computerouter.BgpAttributes](cr.ref.Append("bgp"))
}

func (cr computeRouterAttributes) Timeouts() computerouter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computerouter.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type computeRouterState struct {
	CreationTimestamp           string                       `json:"creation_timestamp"`
	Description                 string                       `json:"description"`
	EncryptedInterconnectRouter bool                         `json:"encrypted_interconnect_router"`
	Id                          string                       `json:"id"`
	Name                        string                       `json:"name"`
	Network                     string                       `json:"network"`
	Project                     string                       `json:"project"`
	Region                      string                       `json:"region"`
	SelfLink                    string                       `json:"self_link"`
	Bgp                         []computerouter.BgpState     `json:"bgp"`
	Timeouts                    *computerouter.TimeoutsState `json:"timeouts"`
}
