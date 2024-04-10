// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cloudsearchdomain "github.com/golingon/terraproviders/aws/5.44.0/cloudsearchdomain"
	"io"
)

// NewCloudsearchDomain creates a new instance of [CloudsearchDomain].
func NewCloudsearchDomain(name string, args CloudsearchDomainArgs) *CloudsearchDomain {
	return &CloudsearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudsearchDomain)(nil)

// CloudsearchDomain represents the Terraform resource aws_cloudsearch_domain.
type CloudsearchDomain struct {
	Name      string
	Args      CloudsearchDomainArgs
	state     *cloudsearchDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudsearchDomain].
func (cd *CloudsearchDomain) Type() string {
	return "aws_cloudsearch_domain"
}

// LocalName returns the local name for [CloudsearchDomain].
func (cd *CloudsearchDomain) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [CloudsearchDomain].
func (cd *CloudsearchDomain) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [CloudsearchDomain].
func (cd *CloudsearchDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [CloudsearchDomain] depends_on.
func (cd *CloudsearchDomain) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudsearchDomain].
func (cd *CloudsearchDomain) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [CloudsearchDomain].
func (cd *CloudsearchDomain) Attributes() cloudsearchDomainAttributes {
	return cloudsearchDomainAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [CloudsearchDomain]'s state.
func (cd *CloudsearchDomain) ImportState(av io.Reader) error {
	cd.state = &cloudsearchDomainState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudsearchDomain] has state.
func (cd *CloudsearchDomain) State() (*cloudsearchDomainState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [CloudsearchDomain]. Panics if the state is nil.
func (cd *CloudsearchDomain) StateMust() *cloudsearchDomainState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// CloudsearchDomainArgs contains the configurations for aws_cloudsearch_domain.
type CloudsearchDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MultiAz: bool, optional
	MultiAz terra.BoolValue `hcl:"multi_az,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// EndpointOptions: optional
	EndpointOptions *cloudsearchdomain.EndpointOptions `hcl:"endpoint_options,block"`
	// IndexField: min=0
	IndexField []cloudsearchdomain.IndexField `hcl:"index_field,block" validate:"min=0"`
	// ScalingParameters: optional
	ScalingParameters *cloudsearchdomain.ScalingParameters `hcl:"scaling_parameters,block"`
	// Timeouts: optional
	Timeouts *cloudsearchdomain.Timeouts `hcl:"timeouts,block"`
}
type cloudsearchDomainAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("arn"))
}

// DocumentServiceEndpoint returns a reference to field document_service_endpoint of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) DocumentServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("document_service_endpoint"))
}

// DomainId returns a reference to field domain_id of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("domain_id"))
}

// Id returns a reference to field id of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// MultiAz returns a reference to field multi_az of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) MultiAz() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("multi_az"))
}

// Name returns a reference to field name of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

// SearchServiceEndpoint returns a reference to field search_service_endpoint of aws_cloudsearch_domain.
func (cd cloudsearchDomainAttributes) SearchServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("search_service_endpoint"))
}

func (cd cloudsearchDomainAttributes) EndpointOptions() terra.ListValue[cloudsearchdomain.EndpointOptionsAttributes] {
	return terra.ReferenceAsList[cloudsearchdomain.EndpointOptionsAttributes](cd.ref.Append("endpoint_options"))
}

func (cd cloudsearchDomainAttributes) IndexField() terra.SetValue[cloudsearchdomain.IndexFieldAttributes] {
	return terra.ReferenceAsSet[cloudsearchdomain.IndexFieldAttributes](cd.ref.Append("index_field"))
}

func (cd cloudsearchDomainAttributes) ScalingParameters() terra.ListValue[cloudsearchdomain.ScalingParametersAttributes] {
	return terra.ReferenceAsList[cloudsearchdomain.ScalingParametersAttributes](cd.ref.Append("scaling_parameters"))
}

func (cd cloudsearchDomainAttributes) Timeouts() cloudsearchdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudsearchdomain.TimeoutsAttributes](cd.ref.Append("timeouts"))
}

type cloudsearchDomainState struct {
	Arn                     string                                     `json:"arn"`
	DocumentServiceEndpoint string                                     `json:"document_service_endpoint"`
	DomainId                string                                     `json:"domain_id"`
	Id                      string                                     `json:"id"`
	MultiAz                 bool                                       `json:"multi_az"`
	Name                    string                                     `json:"name"`
	SearchServiceEndpoint   string                                     `json:"search_service_endpoint"`
	EndpointOptions         []cloudsearchdomain.EndpointOptionsState   `json:"endpoint_options"`
	IndexField              []cloudsearchdomain.IndexFieldState        `json:"index_field"`
	ScalingParameters       []cloudsearchdomain.ScalingParametersState `json:"scaling_parameters"`
	Timeouts                *cloudsearchdomain.TimeoutsState           `json:"timeouts"`
}
