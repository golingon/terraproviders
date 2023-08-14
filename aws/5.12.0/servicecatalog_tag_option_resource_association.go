// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogtagoptionresourceassociation "github.com/golingon/terraproviders/aws/5.12.0/servicecatalogtagoptionresourceassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogTagOptionResourceAssociation creates a new instance of [ServicecatalogTagOptionResourceAssociation].
func NewServicecatalogTagOptionResourceAssociation(name string, args ServicecatalogTagOptionResourceAssociationArgs) *ServicecatalogTagOptionResourceAssociation {
	return &ServicecatalogTagOptionResourceAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogTagOptionResourceAssociation)(nil)

// ServicecatalogTagOptionResourceAssociation represents the Terraform resource aws_servicecatalog_tag_option_resource_association.
type ServicecatalogTagOptionResourceAssociation struct {
	Name      string
	Args      ServicecatalogTagOptionResourceAssociationArgs
	state     *servicecatalogTagOptionResourceAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) Type() string {
	return "aws_servicecatalog_tag_option_resource_association"
}

// LocalName returns the local name for [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) LocalName() string {
	return stora.Name
}

// Configuration returns the configuration (args) for [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) Configuration() interface{} {
	return stora.Args
}

// DependOn is used for other resources to depend on [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(stora)
}

// Dependencies returns the list of resources [ServicecatalogTagOptionResourceAssociation] depends_on.
func (stora *ServicecatalogTagOptionResourceAssociation) Dependencies() terra.Dependencies {
	return stora.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) LifecycleManagement() *terra.Lifecycle {
	return stora.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogTagOptionResourceAssociation].
func (stora *ServicecatalogTagOptionResourceAssociation) Attributes() servicecatalogTagOptionResourceAssociationAttributes {
	return servicecatalogTagOptionResourceAssociationAttributes{ref: terra.ReferenceResource(stora)}
}

// ImportState imports the given attribute values into [ServicecatalogTagOptionResourceAssociation]'s state.
func (stora *ServicecatalogTagOptionResourceAssociation) ImportState(av io.Reader) error {
	stora.state = &servicecatalogTagOptionResourceAssociationState{}
	if err := json.NewDecoder(av).Decode(stora.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", stora.Type(), stora.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogTagOptionResourceAssociation] has state.
func (stora *ServicecatalogTagOptionResourceAssociation) State() (*servicecatalogTagOptionResourceAssociationState, bool) {
	return stora.state, stora.state != nil
}

// StateMust returns the state for [ServicecatalogTagOptionResourceAssociation]. Panics if the state is nil.
func (stora *ServicecatalogTagOptionResourceAssociation) StateMust() *servicecatalogTagOptionResourceAssociationState {
	if stora.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", stora.Type(), stora.LocalName()))
	}
	return stora.state
}

// ServicecatalogTagOptionResourceAssociationArgs contains the configurations for aws_servicecatalog_tag_option_resource_association.
type ServicecatalogTagOptionResourceAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// TagOptionId: string, required
	TagOptionId terra.StringValue `hcl:"tag_option_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogtagoptionresourceassociation.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogTagOptionResourceAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("id"))
}

// ResourceArn returns a reference to field resource_arn of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("resource_arn"))
}

// ResourceCreatedTime returns a reference to field resource_created_time of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) ResourceCreatedTime() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("resource_created_time"))
}

// ResourceDescription returns a reference to field resource_description of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) ResourceDescription() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("resource_description"))
}

// ResourceId returns a reference to field resource_id of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("resource_id"))
}

// ResourceName returns a reference to field resource_name of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("resource_name"))
}

// TagOptionId returns a reference to field tag_option_id of aws_servicecatalog_tag_option_resource_association.
func (stora servicecatalogTagOptionResourceAssociationAttributes) TagOptionId() terra.StringValue {
	return terra.ReferenceAsString(stora.ref.Append("tag_option_id"))
}

func (stora servicecatalogTagOptionResourceAssociationAttributes) Timeouts() servicecatalogtagoptionresourceassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogtagoptionresourceassociation.TimeoutsAttributes](stora.ref.Append("timeouts"))
}

type servicecatalogTagOptionResourceAssociationState struct {
	Id                  string                                                    `json:"id"`
	ResourceArn         string                                                    `json:"resource_arn"`
	ResourceCreatedTime string                                                    `json:"resource_created_time"`
	ResourceDescription string                                                    `json:"resource_description"`
	ResourceId          string                                                    `json:"resource_id"`
	ResourceName        string                                                    `json:"resource_name"`
	TagOptionId         string                                                    `json:"tag_option_id"`
	Timeouts            *servicecatalogtagoptionresourceassociation.TimeoutsState `json:"timeouts"`
}
