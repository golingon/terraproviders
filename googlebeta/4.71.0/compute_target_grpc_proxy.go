// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computetargetgrpcproxy "github.com/golingon/terraproviders/googlebeta/4.71.0/computetargetgrpcproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeTargetGrpcProxy creates a new instance of [ComputeTargetGrpcProxy].
func NewComputeTargetGrpcProxy(name string, args ComputeTargetGrpcProxyArgs) *ComputeTargetGrpcProxy {
	return &ComputeTargetGrpcProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetGrpcProxy)(nil)

// ComputeTargetGrpcProxy represents the Terraform resource google_compute_target_grpc_proxy.
type ComputeTargetGrpcProxy struct {
	Name      string
	Args      ComputeTargetGrpcProxyArgs
	state     *computeTargetGrpcProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) Type() string {
	return "google_compute_target_grpc_proxy"
}

// LocalName returns the local name for [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) LocalName() string {
	return ctgp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) Configuration() interface{} {
	return ctgp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(ctgp)
}

// Dependencies returns the list of resources [ComputeTargetGrpcProxy] depends_on.
func (ctgp *ComputeTargetGrpcProxy) Dependencies() terra.Dependencies {
	return ctgp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) LifecycleManagement() *terra.Lifecycle {
	return ctgp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetGrpcProxy].
func (ctgp *ComputeTargetGrpcProxy) Attributes() computeTargetGrpcProxyAttributes {
	return computeTargetGrpcProxyAttributes{ref: terra.ReferenceResource(ctgp)}
}

// ImportState imports the given attribute values into [ComputeTargetGrpcProxy]'s state.
func (ctgp *ComputeTargetGrpcProxy) ImportState(av io.Reader) error {
	ctgp.state = &computeTargetGrpcProxyState{}
	if err := json.NewDecoder(av).Decode(ctgp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctgp.Type(), ctgp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetGrpcProxy] has state.
func (ctgp *ComputeTargetGrpcProxy) State() (*computeTargetGrpcProxyState, bool) {
	return ctgp.state, ctgp.state != nil
}

// StateMust returns the state for [ComputeTargetGrpcProxy]. Panics if the state is nil.
func (ctgp *ComputeTargetGrpcProxy) StateMust() *computeTargetGrpcProxyState {
	if ctgp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctgp.Type(), ctgp.LocalName()))
	}
	return ctgp.state
}

// ComputeTargetGrpcProxyArgs contains the configurations for google_compute_target_grpc_proxy.
type ComputeTargetGrpcProxyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// UrlMap: string, optional
	UrlMap terra.StringValue `hcl:"url_map,attr"`
	// ValidateForProxyless: bool, optional
	ValidateForProxyless terra.BoolValue `hcl:"validate_for_proxyless,attr"`
	// Timeouts: optional
	Timeouts *computetargetgrpcproxy.Timeouts `hcl:"timeouts,block"`
}
type computeTargetGrpcProxyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("self_link"))
}

// SelfLinkWithId returns a reference to field self_link_with_id of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) SelfLinkWithId() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("self_link_with_id"))
}

// UrlMap returns a reference to field url_map of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(ctgp.ref.Append("url_map"))
}

// ValidateForProxyless returns a reference to field validate_for_proxyless of google_compute_target_grpc_proxy.
func (ctgp computeTargetGrpcProxyAttributes) ValidateForProxyless() terra.BoolValue {
	return terra.ReferenceAsBool(ctgp.ref.Append("validate_for_proxyless"))
}

func (ctgp computeTargetGrpcProxyAttributes) Timeouts() computetargetgrpcproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargetgrpcproxy.TimeoutsAttributes](ctgp.ref.Append("timeouts"))
}

type computeTargetGrpcProxyState struct {
	CreationTimestamp    string                                `json:"creation_timestamp"`
	Description          string                                `json:"description"`
	Fingerprint          string                                `json:"fingerprint"`
	Id                   string                                `json:"id"`
	Name                 string                                `json:"name"`
	Project              string                                `json:"project"`
	SelfLink             string                                `json:"self_link"`
	SelfLinkWithId       string                                `json:"self_link_with_id"`
	UrlMap               string                                `json:"url_map"`
	ValidateForProxyless bool                                  `json:"validate_for_proxyless"`
	Timeouts             *computetargetgrpcproxy.TimeoutsState `json:"timeouts"`
}
