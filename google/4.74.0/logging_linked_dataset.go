// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	logginglinkeddataset "github.com/golingon/terraproviders/google/4.74.0/logginglinkeddataset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingLinkedDataset creates a new instance of [LoggingLinkedDataset].
func NewLoggingLinkedDataset(name string, args LoggingLinkedDatasetArgs) *LoggingLinkedDataset {
	return &LoggingLinkedDataset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingLinkedDataset)(nil)

// LoggingLinkedDataset represents the Terraform resource google_logging_linked_dataset.
type LoggingLinkedDataset struct {
	Name      string
	Args      LoggingLinkedDatasetArgs
	state     *loggingLinkedDatasetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) Type() string {
	return "google_logging_linked_dataset"
}

// LocalName returns the local name for [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) LocalName() string {
	return lld.Name
}

// Configuration returns the configuration (args) for [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) Configuration() interface{} {
	return lld.Args
}

// DependOn is used for other resources to depend on [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) DependOn() terra.Reference {
	return terra.ReferenceResource(lld)
}

// Dependencies returns the list of resources [LoggingLinkedDataset] depends_on.
func (lld *LoggingLinkedDataset) Dependencies() terra.Dependencies {
	return lld.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) LifecycleManagement() *terra.Lifecycle {
	return lld.Lifecycle
}

// Attributes returns the attributes for [LoggingLinkedDataset].
func (lld *LoggingLinkedDataset) Attributes() loggingLinkedDatasetAttributes {
	return loggingLinkedDatasetAttributes{ref: terra.ReferenceResource(lld)}
}

// ImportState imports the given attribute values into [LoggingLinkedDataset]'s state.
func (lld *LoggingLinkedDataset) ImportState(av io.Reader) error {
	lld.state = &loggingLinkedDatasetState{}
	if err := json.NewDecoder(av).Decode(lld.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lld.Type(), lld.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingLinkedDataset] has state.
func (lld *LoggingLinkedDataset) State() (*loggingLinkedDatasetState, bool) {
	return lld.state, lld.state != nil
}

// StateMust returns the state for [LoggingLinkedDataset]. Panics if the state is nil.
func (lld *LoggingLinkedDataset) StateMust() *loggingLinkedDatasetState {
	if lld.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lld.Type(), lld.LocalName()))
	}
	return lld.state
}

// LoggingLinkedDatasetArgs contains the configurations for google_logging_linked_dataset.
type LoggingLinkedDatasetArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, required
	LinkId terra.StringValue `hcl:"link_id,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// BigqueryDataset: min=0
	BigqueryDataset []logginglinkeddataset.BigqueryDataset `hcl:"bigquery_dataset,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *logginglinkeddataset.Timeouts `hcl:"timeouts,block"`
}
type loggingLinkedDatasetAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("bucket"))
}

// CreateTime returns a reference to field create_time of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("create_time"))
}

// Description returns a reference to field description of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("description"))
}

// Id returns a reference to field id of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("lifecycle_state"))
}

// LinkId returns a reference to field link_id of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("link_id"))
}

// Location returns a reference to field location of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("location"))
}

// Name returns a reference to field name of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("name"))
}

// Parent returns a reference to field parent of google_logging_linked_dataset.
func (lld loggingLinkedDatasetAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(lld.ref.Append("parent"))
}

func (lld loggingLinkedDatasetAttributes) BigqueryDataset() terra.ListValue[logginglinkeddataset.BigqueryDatasetAttributes] {
	return terra.ReferenceAsList[logginglinkeddataset.BigqueryDatasetAttributes](lld.ref.Append("bigquery_dataset"))
}

func (lld loggingLinkedDatasetAttributes) Timeouts() logginglinkeddataset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logginglinkeddataset.TimeoutsAttributes](lld.ref.Append("timeouts"))
}

type loggingLinkedDatasetState struct {
	Bucket          string                                      `json:"bucket"`
	CreateTime      string                                      `json:"create_time"`
	Description     string                                      `json:"description"`
	Id              string                                      `json:"id"`
	LifecycleState  string                                      `json:"lifecycle_state"`
	LinkId          string                                      `json:"link_id"`
	Location        string                                      `json:"location"`
	Name            string                                      `json:"name"`
	Parent          string                                      `json:"parent"`
	BigqueryDataset []logginglinkeddataset.BigqueryDatasetState `json:"bigquery_dataset"`
	Timeouts        *logginglinkeddataset.TimeoutsState         `json:"timeouts"`
}
