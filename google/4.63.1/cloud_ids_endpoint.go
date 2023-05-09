// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudidsendpoint "github.com/golingon/terraproviders/google/4.63.1/cloudidsendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudIdsEndpoint creates a new instance of [CloudIdsEndpoint].
func NewCloudIdsEndpoint(name string, args CloudIdsEndpointArgs) *CloudIdsEndpoint {
	return &CloudIdsEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudIdsEndpoint)(nil)

// CloudIdsEndpoint represents the Terraform resource google_cloud_ids_endpoint.
type CloudIdsEndpoint struct {
	Name      string
	Args      CloudIdsEndpointArgs
	state     *cloudIdsEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) Type() string {
	return "google_cloud_ids_endpoint"
}

// LocalName returns the local name for [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) LocalName() string {
	return cie.Name
}

// Configuration returns the configuration (args) for [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) Configuration() interface{} {
	return cie.Args
}

// DependOn is used for other resources to depend on [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(cie)
}

// Dependencies returns the list of resources [CloudIdsEndpoint] depends_on.
func (cie *CloudIdsEndpoint) Dependencies() terra.Dependencies {
	return cie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) LifecycleManagement() *terra.Lifecycle {
	return cie.Lifecycle
}

// Attributes returns the attributes for [CloudIdsEndpoint].
func (cie *CloudIdsEndpoint) Attributes() cloudIdsEndpointAttributes {
	return cloudIdsEndpointAttributes{ref: terra.ReferenceResource(cie)}
}

// ImportState imports the given attribute values into [CloudIdsEndpoint]'s state.
func (cie *CloudIdsEndpoint) ImportState(av io.Reader) error {
	cie.state = &cloudIdsEndpointState{}
	if err := json.NewDecoder(av).Decode(cie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cie.Type(), cie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudIdsEndpoint] has state.
func (cie *CloudIdsEndpoint) State() (*cloudIdsEndpointState, bool) {
	return cie.state, cie.state != nil
}

// StateMust returns the state for [CloudIdsEndpoint]. Panics if the state is nil.
func (cie *CloudIdsEndpoint) StateMust() *cloudIdsEndpointState {
	if cie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cie.Type(), cie.LocalName()))
	}
	return cie.state
}

// CloudIdsEndpointArgs contains the configurations for google_cloud_ids_endpoint.
type CloudIdsEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Severity: string, required
	Severity terra.StringValue `hcl:"severity,attr" validate:"required"`
	// ThreatExceptions: list of string, optional
	ThreatExceptions terra.ListValue[terra.StringValue] `hcl:"threat_exceptions,attr"`
	// Timeouts: optional
	Timeouts *cloudidsendpoint.Timeouts `hcl:"timeouts,block"`
}
type cloudIdsEndpointAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("create_time"))
}

// Description returns a reference to field description of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("description"))
}

// EndpointForwardingRule returns a reference to field endpoint_forwarding_rule of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) EndpointForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("endpoint_forwarding_rule"))
}

// EndpointIp returns a reference to field endpoint_ip of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) EndpointIp() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("endpoint_ip"))
}

// Id returns a reference to field id of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("name"))
}

// Network returns a reference to field network of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("network"))
}

// Project returns a reference to field project of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("project"))
}

// Severity returns a reference to field severity of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("severity"))
}

// ThreatExceptions returns a reference to field threat_exceptions of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) ThreatExceptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cie.ref.Append("threat_exceptions"))
}

// UpdateTime returns a reference to field update_time of google_cloud_ids_endpoint.
func (cie cloudIdsEndpointAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cie.ref.Append("update_time"))
}

func (cie cloudIdsEndpointAttributes) Timeouts() cloudidsendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudidsendpoint.TimeoutsAttributes](cie.ref.Append("timeouts"))
}

type cloudIdsEndpointState struct {
	CreateTime             string                          `json:"create_time"`
	Description            string                          `json:"description"`
	EndpointForwardingRule string                          `json:"endpoint_forwarding_rule"`
	EndpointIp             string                          `json:"endpoint_ip"`
	Id                     string                          `json:"id"`
	Location               string                          `json:"location"`
	Name                   string                          `json:"name"`
	Network                string                          `json:"network"`
	Project                string                          `json:"project"`
	Severity               string                          `json:"severity"`
	ThreatExceptions       []string                        `json:"threat_exceptions"`
	UpdateTime             string                          `json:"update_time"`
	Timeouts               *cloudidsendpoint.TimeoutsState `json:"timeouts"`
}
