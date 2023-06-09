// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogprovisionedproduct "github.com/golingon/terraproviders/aws/4.60.0/servicecatalogprovisionedproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogProvisionedProduct creates a new instance of [ServicecatalogProvisionedProduct].
func NewServicecatalogProvisionedProduct(name string, args ServicecatalogProvisionedProductArgs) *ServicecatalogProvisionedProduct {
	return &ServicecatalogProvisionedProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogProvisionedProduct)(nil)

// ServicecatalogProvisionedProduct represents the Terraform resource aws_servicecatalog_provisioned_product.
type ServicecatalogProvisionedProduct struct {
	Name      string
	Args      ServicecatalogProvisionedProductArgs
	state     *servicecatalogProvisionedProductState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) Type() string {
	return "aws_servicecatalog_provisioned_product"
}

// LocalName returns the local name for [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) LocalName() string {
	return spp.Name
}

// Configuration returns the configuration (args) for [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) Configuration() interface{} {
	return spp.Args
}

// DependOn is used for other resources to depend on [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) DependOn() terra.Reference {
	return terra.ReferenceResource(spp)
}

// Dependencies returns the list of resources [ServicecatalogProvisionedProduct] depends_on.
func (spp *ServicecatalogProvisionedProduct) Dependencies() terra.Dependencies {
	return spp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) LifecycleManagement() *terra.Lifecycle {
	return spp.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogProvisionedProduct].
func (spp *ServicecatalogProvisionedProduct) Attributes() servicecatalogProvisionedProductAttributes {
	return servicecatalogProvisionedProductAttributes{ref: terra.ReferenceResource(spp)}
}

// ImportState imports the given attribute values into [ServicecatalogProvisionedProduct]'s state.
func (spp *ServicecatalogProvisionedProduct) ImportState(av io.Reader) error {
	spp.state = &servicecatalogProvisionedProductState{}
	if err := json.NewDecoder(av).Decode(spp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spp.Type(), spp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogProvisionedProduct] has state.
func (spp *ServicecatalogProvisionedProduct) State() (*servicecatalogProvisionedProductState, bool) {
	return spp.state, spp.state != nil
}

// StateMust returns the state for [ServicecatalogProvisionedProduct]. Panics if the state is nil.
func (spp *ServicecatalogProvisionedProduct) StateMust() *servicecatalogProvisionedProductState {
	if spp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spp.Type(), spp.LocalName()))
	}
	return spp.state
}

// ServicecatalogProvisionedProductArgs contains the configurations for aws_servicecatalog_provisioned_product.
type ServicecatalogProvisionedProductArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreErrors: bool, optional
	IgnoreErrors terra.BoolValue `hcl:"ignore_errors,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationArns: list of string, optional
	NotificationArns terra.ListValue[terra.StringValue] `hcl:"notification_arns,attr"`
	// PathId: string, optional
	PathId terra.StringValue `hcl:"path_id,attr"`
	// PathName: string, optional
	PathName terra.StringValue `hcl:"path_name,attr"`
	// ProductId: string, optional
	ProductId terra.StringValue `hcl:"product_id,attr"`
	// ProductName: string, optional
	ProductName terra.StringValue `hcl:"product_name,attr"`
	// ProvisioningArtifactId: string, optional
	ProvisioningArtifactId terra.StringValue `hcl:"provisioning_artifact_id,attr"`
	// ProvisioningArtifactName: string, optional
	ProvisioningArtifactName terra.StringValue `hcl:"provisioning_artifact_name,attr"`
	// RetainPhysicalResources: bool, optional
	RetainPhysicalResources terra.BoolValue `hcl:"retain_physical_resources,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Outputs: min=0
	Outputs []servicecatalogprovisionedproduct.Outputs `hcl:"outputs,block" validate:"min=0"`
	// ProvisioningParameters: min=0
	ProvisioningParameters []servicecatalogprovisionedproduct.ProvisioningParameters `hcl:"provisioning_parameters,block" validate:"min=0"`
	// StackSetProvisioningPreferences: optional
	StackSetProvisioningPreferences *servicecatalogprovisionedproduct.StackSetProvisioningPreferences `hcl:"stack_set_provisioning_preferences,block"`
	// Timeouts: optional
	Timeouts *servicecatalogprovisionedproduct.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogProvisionedProductAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("accept_language"))
}

// Arn returns a reference to field arn of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("arn"))
}

// CloudwatchDashboardNames returns a reference to field cloudwatch_dashboard_names of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) CloudwatchDashboardNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](spp.ref.Append("cloudwatch_dashboard_names"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("created_time"))
}

// Id returns a reference to field id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("id"))
}

// IgnoreErrors returns a reference to field ignore_errors of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) IgnoreErrors() terra.BoolValue {
	return terra.ReferenceAsBool(spp.ref.Append("ignore_errors"))
}

// LastProvisioningRecordId returns a reference to field last_provisioning_record_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) LastProvisioningRecordId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("last_provisioning_record_id"))
}

// LastRecordId returns a reference to field last_record_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) LastRecordId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("last_record_id"))
}

// LastSuccessfulProvisioningRecordId returns a reference to field last_successful_provisioning_record_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) LastSuccessfulProvisioningRecordId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("last_successful_provisioning_record_id"))
}

// LaunchRoleArn returns a reference to field launch_role_arn of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) LaunchRoleArn() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("launch_role_arn"))
}

// Name returns a reference to field name of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("name"))
}

// NotificationArns returns a reference to field notification_arns of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) NotificationArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spp.ref.Append("notification_arns"))
}

// PathId returns a reference to field path_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) PathId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("path_id"))
}

// PathName returns a reference to field path_name of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) PathName() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("path_name"))
}

// ProductId returns a reference to field product_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("product_id"))
}

// ProductName returns a reference to field product_name of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) ProductName() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("product_name"))
}

// ProvisioningArtifactId returns a reference to field provisioning_artifact_id of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) ProvisioningArtifactId() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("provisioning_artifact_id"))
}

// ProvisioningArtifactName returns a reference to field provisioning_artifact_name of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) ProvisioningArtifactName() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("provisioning_artifact_name"))
}

// RetainPhysicalResources returns a reference to field retain_physical_resources of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) RetainPhysicalResources() terra.BoolValue {
	return terra.ReferenceAsBool(spp.ref.Append("retain_physical_resources"))
}

// Status returns a reference to field status of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("status_message"))
}

// Tags returns a reference to field tags of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spp.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_servicecatalog_provisioned_product.
func (spp servicecatalogProvisionedProductAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(spp.ref.Append("type"))
}

func (spp servicecatalogProvisionedProductAttributes) Outputs() terra.SetValue[servicecatalogprovisionedproduct.OutputsAttributes] {
	return terra.ReferenceAsSet[servicecatalogprovisionedproduct.OutputsAttributes](spp.ref.Append("outputs"))
}

func (spp servicecatalogProvisionedProductAttributes) ProvisioningParameters() terra.ListValue[servicecatalogprovisionedproduct.ProvisioningParametersAttributes] {
	return terra.ReferenceAsList[servicecatalogprovisionedproduct.ProvisioningParametersAttributes](spp.ref.Append("provisioning_parameters"))
}

func (spp servicecatalogProvisionedProductAttributes) StackSetProvisioningPreferences() terra.ListValue[servicecatalogprovisionedproduct.StackSetProvisioningPreferencesAttributes] {
	return terra.ReferenceAsList[servicecatalogprovisionedproduct.StackSetProvisioningPreferencesAttributes](spp.ref.Append("stack_set_provisioning_preferences"))
}

func (spp servicecatalogProvisionedProductAttributes) Timeouts() servicecatalogprovisionedproduct.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogprovisionedproduct.TimeoutsAttributes](spp.ref.Append("timeouts"))
}

type servicecatalogProvisionedProductState struct {
	AcceptLanguage                     string                                                                  `json:"accept_language"`
	Arn                                string                                                                  `json:"arn"`
	CloudwatchDashboardNames           []string                                                                `json:"cloudwatch_dashboard_names"`
	CreatedTime                        string                                                                  `json:"created_time"`
	Id                                 string                                                                  `json:"id"`
	IgnoreErrors                       bool                                                                    `json:"ignore_errors"`
	LastProvisioningRecordId           string                                                                  `json:"last_provisioning_record_id"`
	LastRecordId                       string                                                                  `json:"last_record_id"`
	LastSuccessfulProvisioningRecordId string                                                                  `json:"last_successful_provisioning_record_id"`
	LaunchRoleArn                      string                                                                  `json:"launch_role_arn"`
	Name                               string                                                                  `json:"name"`
	NotificationArns                   []string                                                                `json:"notification_arns"`
	PathId                             string                                                                  `json:"path_id"`
	PathName                           string                                                                  `json:"path_name"`
	ProductId                          string                                                                  `json:"product_id"`
	ProductName                        string                                                                  `json:"product_name"`
	ProvisioningArtifactId             string                                                                  `json:"provisioning_artifact_id"`
	ProvisioningArtifactName           string                                                                  `json:"provisioning_artifact_name"`
	RetainPhysicalResources            bool                                                                    `json:"retain_physical_resources"`
	Status                             string                                                                  `json:"status"`
	StatusMessage                      string                                                                  `json:"status_message"`
	Tags                               map[string]string                                                       `json:"tags"`
	TagsAll                            map[string]string                                                       `json:"tags_all"`
	Type                               string                                                                  `json:"type"`
	Outputs                            []servicecatalogprovisionedproduct.OutputsState                         `json:"outputs"`
	ProvisioningParameters             []servicecatalogprovisionedproduct.ProvisioningParametersState          `json:"provisioning_parameters"`
	StackSetProvisioningPreferences    []servicecatalogprovisionedproduct.StackSetProvisioningPreferencesState `json:"stack_set_provisioning_preferences"`
	Timeouts                           *servicecatalogprovisionedproduct.TimeoutsState                         `json:"timeouts"`
}
