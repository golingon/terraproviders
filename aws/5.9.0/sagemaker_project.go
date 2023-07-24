// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerproject "github.com/golingon/terraproviders/aws/5.9.0/sagemakerproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerProject creates a new instance of [SagemakerProject].
func NewSagemakerProject(name string, args SagemakerProjectArgs) *SagemakerProject {
	return &SagemakerProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerProject)(nil)

// SagemakerProject represents the Terraform resource aws_sagemaker_project.
type SagemakerProject struct {
	Name      string
	Args      SagemakerProjectArgs
	state     *sagemakerProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerProject].
func (sp *SagemakerProject) Type() string {
	return "aws_sagemaker_project"
}

// LocalName returns the local name for [SagemakerProject].
func (sp *SagemakerProject) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [SagemakerProject].
func (sp *SagemakerProject) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [SagemakerProject].
func (sp *SagemakerProject) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [SagemakerProject] depends_on.
func (sp *SagemakerProject) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerProject].
func (sp *SagemakerProject) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [SagemakerProject].
func (sp *SagemakerProject) Attributes() sagemakerProjectAttributes {
	return sagemakerProjectAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [SagemakerProject]'s state.
func (sp *SagemakerProject) ImportState(av io.Reader) error {
	sp.state = &sagemakerProjectState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerProject] has state.
func (sp *SagemakerProject) State() (*sagemakerProjectState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [SagemakerProject]. Panics if the state is nil.
func (sp *SagemakerProject) StateMust() *sagemakerProjectState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// SagemakerProjectArgs contains the configurations for aws_sagemaker_project.
type SagemakerProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProjectDescription: string, optional
	ProjectDescription terra.StringValue `hcl:"project_description,attr"`
	// ProjectName: string, required
	ProjectName terra.StringValue `hcl:"project_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ServiceCatalogProvisioningDetails: required
	ServiceCatalogProvisioningDetails *sagemakerproject.ServiceCatalogProvisioningDetails `hcl:"service_catalog_provisioning_details,block" validate:"required"`
}
type sagemakerProjectAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// ProjectDescription returns a reference to field project_description of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) ProjectDescription() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("project_description"))
}

// ProjectId returns a reference to field project_id of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("project_id"))
}

// ProjectName returns a reference to field project_name of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) ProjectName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("project_name"))
}

// Tags returns a reference to field tags of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_project.
func (sp sagemakerProjectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags_all"))
}

func (sp sagemakerProjectAttributes) ServiceCatalogProvisioningDetails() terra.ListValue[sagemakerproject.ServiceCatalogProvisioningDetailsAttributes] {
	return terra.ReferenceAsList[sagemakerproject.ServiceCatalogProvisioningDetailsAttributes](sp.ref.Append("service_catalog_provisioning_details"))
}

type sagemakerProjectState struct {
	Arn                               string                                                    `json:"arn"`
	Id                                string                                                    `json:"id"`
	ProjectDescription                string                                                    `json:"project_description"`
	ProjectId                         string                                                    `json:"project_id"`
	ProjectName                       string                                                    `json:"project_name"`
	Tags                              map[string]string                                         `json:"tags"`
	TagsAll                           map[string]string                                         `json:"tags_all"`
	ServiceCatalogProvisioningDetails []sagemakerproject.ServiceCatalogProvisioningDetailsState `json:"service_catalog_provisioning_details"`
}
