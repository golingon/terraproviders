// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogproduct "github.com/golingon/terraproviders/aws/5.0.1/servicecatalogproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogProduct creates a new instance of [ServicecatalogProduct].
func NewServicecatalogProduct(name string, args ServicecatalogProductArgs) *ServicecatalogProduct {
	return &ServicecatalogProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogProduct)(nil)

// ServicecatalogProduct represents the Terraform resource aws_servicecatalog_product.
type ServicecatalogProduct struct {
	Name      string
	Args      ServicecatalogProductArgs
	state     *servicecatalogProductState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogProduct].
func (sp *ServicecatalogProduct) Type() string {
	return "aws_servicecatalog_product"
}

// LocalName returns the local name for [ServicecatalogProduct].
func (sp *ServicecatalogProduct) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [ServicecatalogProduct].
func (sp *ServicecatalogProduct) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [ServicecatalogProduct].
func (sp *ServicecatalogProduct) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [ServicecatalogProduct] depends_on.
func (sp *ServicecatalogProduct) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogProduct].
func (sp *ServicecatalogProduct) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogProduct].
func (sp *ServicecatalogProduct) Attributes() servicecatalogProductAttributes {
	return servicecatalogProductAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [ServicecatalogProduct]'s state.
func (sp *ServicecatalogProduct) ImportState(av io.Reader) error {
	sp.state = &servicecatalogProductState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogProduct] has state.
func (sp *ServicecatalogProduct) State() (*servicecatalogProductState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [ServicecatalogProduct]. Panics if the state is nil.
func (sp *ServicecatalogProduct) StateMust() *servicecatalogProductState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// ServicecatalogProductArgs contains the configurations for aws_servicecatalog_product.
type ServicecatalogProductArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Distributor: string, optional
	Distributor terra.StringValue `hcl:"distributor,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
	// SupportDescription: string, optional
	SupportDescription terra.StringValue `hcl:"support_description,attr"`
	// SupportEmail: string, optional
	SupportEmail terra.StringValue `hcl:"support_email,attr"`
	// SupportUrl: string, optional
	SupportUrl terra.StringValue `hcl:"support_url,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ProvisioningArtifactParameters: required
	ProvisioningArtifactParameters *servicecatalogproduct.ProvisioningArtifactParameters `hcl:"provisioning_artifact_parameters,block" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogproduct.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogProductAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("accept_language"))
}

// Arn returns a reference to field arn of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// Distributor returns a reference to field distributor of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Distributor() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("distributor"))
}

// HasDefaultPath returns a reference to field has_default_path of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) HasDefaultPath() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("has_default_path"))
}

// Id returns a reference to field id of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("owner"))
}

// Status returns a reference to field status of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("status"))
}

// SupportDescription returns a reference to field support_description of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) SupportDescription() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_description"))
}

// SupportEmail returns a reference to field support_email of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) SupportEmail() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_email"))
}

// SupportUrl returns a reference to field support_url of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) SupportUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("support_url"))
}

// Tags returns a reference to field tags of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_servicecatalog_product.
func (sp servicecatalogProductAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

func (sp servicecatalogProductAttributes) ProvisioningArtifactParameters() terra.ListValue[servicecatalogproduct.ProvisioningArtifactParametersAttributes] {
	return terra.ReferenceAsList[servicecatalogproduct.ProvisioningArtifactParametersAttributes](sp.ref.Append("provisioning_artifact_parameters"))
}

func (sp servicecatalogProductAttributes) Timeouts() servicecatalogproduct.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogproduct.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

type servicecatalogProductState struct {
	AcceptLanguage                 string                                                      `json:"accept_language"`
	Arn                            string                                                      `json:"arn"`
	CreatedTime                    string                                                      `json:"created_time"`
	Description                    string                                                      `json:"description"`
	Distributor                    string                                                      `json:"distributor"`
	HasDefaultPath                 bool                                                        `json:"has_default_path"`
	Id                             string                                                      `json:"id"`
	Name                           string                                                      `json:"name"`
	Owner                          string                                                      `json:"owner"`
	Status                         string                                                      `json:"status"`
	SupportDescription             string                                                      `json:"support_description"`
	SupportEmail                   string                                                      `json:"support_email"`
	SupportUrl                     string                                                      `json:"support_url"`
	Tags                           map[string]string                                           `json:"tags"`
	TagsAll                        map[string]string                                           `json:"tags_all"`
	Type                           string                                                      `json:"type"`
	ProvisioningArtifactParameters []servicecatalogproduct.ProvisioningArtifactParametersState `json:"provisioning_artifact_parameters"`
	Timeouts                       *servicecatalogproduct.TimeoutsState                        `json:"timeouts"`
}
