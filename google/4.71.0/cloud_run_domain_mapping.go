// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrundomainmapping "github.com/golingon/terraproviders/google/4.71.0/cloudrundomainmapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunDomainMapping creates a new instance of [CloudRunDomainMapping].
func NewCloudRunDomainMapping(name string, args CloudRunDomainMappingArgs) *CloudRunDomainMapping {
	return &CloudRunDomainMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunDomainMapping)(nil)

// CloudRunDomainMapping represents the Terraform resource google_cloud_run_domain_mapping.
type CloudRunDomainMapping struct {
	Name      string
	Args      CloudRunDomainMappingArgs
	state     *cloudRunDomainMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) Type() string {
	return "google_cloud_run_domain_mapping"
}

// LocalName returns the local name for [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) LocalName() string {
	return crdm.Name
}

// Configuration returns the configuration (args) for [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) Configuration() interface{} {
	return crdm.Args
}

// DependOn is used for other resources to depend on [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(crdm)
}

// Dependencies returns the list of resources [CloudRunDomainMapping] depends_on.
func (crdm *CloudRunDomainMapping) Dependencies() terra.Dependencies {
	return crdm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) LifecycleManagement() *terra.Lifecycle {
	return crdm.Lifecycle
}

// Attributes returns the attributes for [CloudRunDomainMapping].
func (crdm *CloudRunDomainMapping) Attributes() cloudRunDomainMappingAttributes {
	return cloudRunDomainMappingAttributes{ref: terra.ReferenceResource(crdm)}
}

// ImportState imports the given attribute values into [CloudRunDomainMapping]'s state.
func (crdm *CloudRunDomainMapping) ImportState(av io.Reader) error {
	crdm.state = &cloudRunDomainMappingState{}
	if err := json.NewDecoder(av).Decode(crdm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdm.Type(), crdm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunDomainMapping] has state.
func (crdm *CloudRunDomainMapping) State() (*cloudRunDomainMappingState, bool) {
	return crdm.state, crdm.state != nil
}

// StateMust returns the state for [CloudRunDomainMapping]. Panics if the state is nil.
func (crdm *CloudRunDomainMapping) StateMust() *cloudRunDomainMappingState {
	if crdm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdm.Type(), crdm.LocalName()))
	}
	return crdm.state
}

// CloudRunDomainMappingArgs contains the configurations for google_cloud_run_domain_mapping.
type CloudRunDomainMappingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Status: min=0
	Status []cloudrundomainmapping.Status `hcl:"status,block" validate:"min=0"`
	// Metadata: required
	Metadata *cloudrundomainmapping.Metadata `hcl:"metadata,block" validate:"required"`
	// Spec: required
	Spec *cloudrundomainmapping.Spec `hcl:"spec,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudrundomainmapping.Timeouts `hcl:"timeouts,block"`
}
type cloudRunDomainMappingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_cloud_run_domain_mapping.
func (crdm cloudRunDomainMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crdm.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_domain_mapping.
func (crdm cloudRunDomainMappingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crdm.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_domain_mapping.
func (crdm cloudRunDomainMappingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crdm.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_domain_mapping.
func (crdm cloudRunDomainMappingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crdm.ref.Append("project"))
}

func (crdm cloudRunDomainMappingAttributes) Status() terra.ListValue[cloudrundomainmapping.StatusAttributes] {
	return terra.ReferenceAsList[cloudrundomainmapping.StatusAttributes](crdm.ref.Append("status"))
}

func (crdm cloudRunDomainMappingAttributes) Metadata() terra.ListValue[cloudrundomainmapping.MetadataAttributes] {
	return terra.ReferenceAsList[cloudrundomainmapping.MetadataAttributes](crdm.ref.Append("metadata"))
}

func (crdm cloudRunDomainMappingAttributes) Spec() terra.ListValue[cloudrundomainmapping.SpecAttributes] {
	return terra.ReferenceAsList[cloudrundomainmapping.SpecAttributes](crdm.ref.Append("spec"))
}

func (crdm cloudRunDomainMappingAttributes) Timeouts() cloudrundomainmapping.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudrundomainmapping.TimeoutsAttributes](crdm.ref.Append("timeouts"))
}

type cloudRunDomainMappingState struct {
	Id       string                                `json:"id"`
	Location string                                `json:"location"`
	Name     string                                `json:"name"`
	Project  string                                `json:"project"`
	Status   []cloudrundomainmapping.StatusState   `json:"status"`
	Metadata []cloudrundomainmapping.MetadataState `json:"metadata"`
	Spec     []cloudrundomainmapping.SpecState     `json:"spec"`
	Timeouts *cloudrundomainmapping.TimeoutsState  `json:"timeouts"`
}
