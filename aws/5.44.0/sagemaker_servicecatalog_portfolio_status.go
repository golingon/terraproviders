// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSagemakerServicecatalogPortfolioStatus creates a new instance of [SagemakerServicecatalogPortfolioStatus].
func NewSagemakerServicecatalogPortfolioStatus(name string, args SagemakerServicecatalogPortfolioStatusArgs) *SagemakerServicecatalogPortfolioStatus {
	return &SagemakerServicecatalogPortfolioStatus{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerServicecatalogPortfolioStatus)(nil)

// SagemakerServicecatalogPortfolioStatus represents the Terraform resource aws_sagemaker_servicecatalog_portfolio_status.
type SagemakerServicecatalogPortfolioStatus struct {
	Name      string
	Args      SagemakerServicecatalogPortfolioStatusArgs
	state     *sagemakerServicecatalogPortfolioStatusState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) Type() string {
	return "aws_sagemaker_servicecatalog_portfolio_status"
}

// LocalName returns the local name for [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) LocalName() string {
	return ssps.Name
}

// Configuration returns the configuration (args) for [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) Configuration() interface{} {
	return ssps.Args
}

// DependOn is used for other resources to depend on [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) DependOn() terra.Reference {
	return terra.ReferenceResource(ssps)
}

// Dependencies returns the list of resources [SagemakerServicecatalogPortfolioStatus] depends_on.
func (ssps *SagemakerServicecatalogPortfolioStatus) Dependencies() terra.Dependencies {
	return ssps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) LifecycleManagement() *terra.Lifecycle {
	return ssps.Lifecycle
}

// Attributes returns the attributes for [SagemakerServicecatalogPortfolioStatus].
func (ssps *SagemakerServicecatalogPortfolioStatus) Attributes() sagemakerServicecatalogPortfolioStatusAttributes {
	return sagemakerServicecatalogPortfolioStatusAttributes{ref: terra.ReferenceResource(ssps)}
}

// ImportState imports the given attribute values into [SagemakerServicecatalogPortfolioStatus]'s state.
func (ssps *SagemakerServicecatalogPortfolioStatus) ImportState(av io.Reader) error {
	ssps.state = &sagemakerServicecatalogPortfolioStatusState{}
	if err := json.NewDecoder(av).Decode(ssps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssps.Type(), ssps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerServicecatalogPortfolioStatus] has state.
func (ssps *SagemakerServicecatalogPortfolioStatus) State() (*sagemakerServicecatalogPortfolioStatusState, bool) {
	return ssps.state, ssps.state != nil
}

// StateMust returns the state for [SagemakerServicecatalogPortfolioStatus]. Panics if the state is nil.
func (ssps *SagemakerServicecatalogPortfolioStatus) StateMust() *sagemakerServicecatalogPortfolioStatusState {
	if ssps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssps.Type(), ssps.LocalName()))
	}
	return ssps.state
}

// SagemakerServicecatalogPortfolioStatusArgs contains the configurations for aws_sagemaker_servicecatalog_portfolio_status.
type SagemakerServicecatalogPortfolioStatusArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
}
type sagemakerServicecatalogPortfolioStatusAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_sagemaker_servicecatalog_portfolio_status.
func (ssps sagemakerServicecatalogPortfolioStatusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssps.ref.Append("id"))
}

// Status returns a reference to field status of aws_sagemaker_servicecatalog_portfolio_status.
func (ssps sagemakerServicecatalogPortfolioStatusAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssps.ref.Append("status"))
}

type sagemakerServicecatalogPortfolioStatusState struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}
