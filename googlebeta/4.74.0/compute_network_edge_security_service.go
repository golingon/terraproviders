// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computenetworkedgesecurityservice "github.com/golingon/terraproviders/googlebeta/4.74.0/computenetworkedgesecurityservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkEdgeSecurityService creates a new instance of [ComputeNetworkEdgeSecurityService].
func NewComputeNetworkEdgeSecurityService(name string, args ComputeNetworkEdgeSecurityServiceArgs) *ComputeNetworkEdgeSecurityService {
	return &ComputeNetworkEdgeSecurityService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkEdgeSecurityService)(nil)

// ComputeNetworkEdgeSecurityService represents the Terraform resource google_compute_network_edge_security_service.
type ComputeNetworkEdgeSecurityService struct {
	Name      string
	Args      ComputeNetworkEdgeSecurityServiceArgs
	state     *computeNetworkEdgeSecurityServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) Type() string {
	return "google_compute_network_edge_security_service"
}

// LocalName returns the local name for [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) LocalName() string {
	return cness.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) Configuration() interface{} {
	return cness.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) DependOn() terra.Reference {
	return terra.ReferenceResource(cness)
}

// Dependencies returns the list of resources [ComputeNetworkEdgeSecurityService] depends_on.
func (cness *ComputeNetworkEdgeSecurityService) Dependencies() terra.Dependencies {
	return cness.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) LifecycleManagement() *terra.Lifecycle {
	return cness.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkEdgeSecurityService].
func (cness *ComputeNetworkEdgeSecurityService) Attributes() computeNetworkEdgeSecurityServiceAttributes {
	return computeNetworkEdgeSecurityServiceAttributes{ref: terra.ReferenceResource(cness)}
}

// ImportState imports the given attribute values into [ComputeNetworkEdgeSecurityService]'s state.
func (cness *ComputeNetworkEdgeSecurityService) ImportState(av io.Reader) error {
	cness.state = &computeNetworkEdgeSecurityServiceState{}
	if err := json.NewDecoder(av).Decode(cness.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cness.Type(), cness.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkEdgeSecurityService] has state.
func (cness *ComputeNetworkEdgeSecurityService) State() (*computeNetworkEdgeSecurityServiceState, bool) {
	return cness.state, cness.state != nil
}

// StateMust returns the state for [ComputeNetworkEdgeSecurityService]. Panics if the state is nil.
func (cness *ComputeNetworkEdgeSecurityService) StateMust() *computeNetworkEdgeSecurityServiceState {
	if cness.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cness.Type(), cness.LocalName()))
	}
	return cness.state
}

// ComputeNetworkEdgeSecurityServiceArgs contains the configurations for google_compute_network_edge_security_service.
type ComputeNetworkEdgeSecurityServiceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
	// Timeouts: optional
	Timeouts *computenetworkedgesecurityservice.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkEdgeSecurityServiceAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("region"))
}

// SecurityPolicy returns a reference to field security_policy of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("security_policy"))
}

// SelfLink returns a reference to field self_link of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("self_link"))
}

// SelfLinkWithServiceId returns a reference to field self_link_with_service_id of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) SelfLinkWithServiceId() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("self_link_with_service_id"))
}

// ServiceId returns a reference to field service_id of google_compute_network_edge_security_service.
func (cness computeNetworkEdgeSecurityServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(cness.ref.Append("service_id"))
}

func (cness computeNetworkEdgeSecurityServiceAttributes) Timeouts() computenetworkedgesecurityservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkedgesecurityservice.TimeoutsAttributes](cness.ref.Append("timeouts"))
}

type computeNetworkEdgeSecurityServiceState struct {
	CreationTimestamp     string                                           `json:"creation_timestamp"`
	Description           string                                           `json:"description"`
	Fingerprint           string                                           `json:"fingerprint"`
	Id                    string                                           `json:"id"`
	Name                  string                                           `json:"name"`
	Project               string                                           `json:"project"`
	Region                string                                           `json:"region"`
	SecurityPolicy        string                                           `json:"security_policy"`
	SelfLink              string                                           `json:"self_link"`
	SelfLinkWithServiceId string                                           `json:"self_link_with_service_id"`
	ServiceId             string                                           `json:"service_id"`
	Timeouts              *computenetworkedgesecurityservice.TimeoutsState `json:"timeouts"`
}
