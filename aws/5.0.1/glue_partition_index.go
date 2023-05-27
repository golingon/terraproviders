// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluepartitionindex "github.com/golingon/terraproviders/aws/5.0.1/gluepartitionindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGluePartitionIndex creates a new instance of [GluePartitionIndex].
func NewGluePartitionIndex(name string, args GluePartitionIndexArgs) *GluePartitionIndex {
	return &GluePartitionIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GluePartitionIndex)(nil)

// GluePartitionIndex represents the Terraform resource aws_glue_partition_index.
type GluePartitionIndex struct {
	Name      string
	Args      GluePartitionIndexArgs
	state     *gluePartitionIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GluePartitionIndex].
func (gpi *GluePartitionIndex) Type() string {
	return "aws_glue_partition_index"
}

// LocalName returns the local name for [GluePartitionIndex].
func (gpi *GluePartitionIndex) LocalName() string {
	return gpi.Name
}

// Configuration returns the configuration (args) for [GluePartitionIndex].
func (gpi *GluePartitionIndex) Configuration() interface{} {
	return gpi.Args
}

// DependOn is used for other resources to depend on [GluePartitionIndex].
func (gpi *GluePartitionIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(gpi)
}

// Dependencies returns the list of resources [GluePartitionIndex] depends_on.
func (gpi *GluePartitionIndex) Dependencies() terra.Dependencies {
	return gpi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GluePartitionIndex].
func (gpi *GluePartitionIndex) LifecycleManagement() *terra.Lifecycle {
	return gpi.Lifecycle
}

// Attributes returns the attributes for [GluePartitionIndex].
func (gpi *GluePartitionIndex) Attributes() gluePartitionIndexAttributes {
	return gluePartitionIndexAttributes{ref: terra.ReferenceResource(gpi)}
}

// ImportState imports the given attribute values into [GluePartitionIndex]'s state.
func (gpi *GluePartitionIndex) ImportState(av io.Reader) error {
	gpi.state = &gluePartitionIndexState{}
	if err := json.NewDecoder(av).Decode(gpi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gpi.Type(), gpi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GluePartitionIndex] has state.
func (gpi *GluePartitionIndex) State() (*gluePartitionIndexState, bool) {
	return gpi.state, gpi.state != nil
}

// StateMust returns the state for [GluePartitionIndex]. Panics if the state is nil.
func (gpi *GluePartitionIndex) StateMust() *gluePartitionIndexState {
	if gpi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gpi.Type(), gpi.LocalName()))
	}
	return gpi.state
}

// GluePartitionIndexArgs contains the configurations for aws_glue_partition_index.
type GluePartitionIndexArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// PartitionIndex: required
	PartitionIndex *gluepartitionindex.PartitionIndex `hcl:"partition_index,block" validate:"required"`
	// Timeouts: optional
	Timeouts *gluepartitionindex.Timeouts `hcl:"timeouts,block"`
}
type gluePartitionIndexAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_glue_partition_index.
func (gpi gluePartitionIndexAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gpi.ref.Append("catalog_id"))
}

// DatabaseName returns a reference to field database_name of aws_glue_partition_index.
func (gpi gluePartitionIndexAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(gpi.ref.Append("database_name"))
}

// Id returns a reference to field id of aws_glue_partition_index.
func (gpi gluePartitionIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpi.ref.Append("id"))
}

// TableName returns a reference to field table_name of aws_glue_partition_index.
func (gpi gluePartitionIndexAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(gpi.ref.Append("table_name"))
}

func (gpi gluePartitionIndexAttributes) PartitionIndex() terra.ListValue[gluepartitionindex.PartitionIndexAttributes] {
	return terra.ReferenceAsList[gluepartitionindex.PartitionIndexAttributes](gpi.ref.Append("partition_index"))
}

func (gpi gluePartitionIndexAttributes) Timeouts() gluepartitionindex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gluepartitionindex.TimeoutsAttributes](gpi.ref.Append("timeouts"))
}

type gluePartitionIndexState struct {
	CatalogId      string                                   `json:"catalog_id"`
	DatabaseName   string                                   `json:"database_name"`
	Id             string                                   `json:"id"`
	TableName      string                                   `json:"table_name"`
	PartitionIndex []gluepartitionindex.PartitionIndexState `json:"partition_index"`
	Timeouts       *gluepartitionindex.TimeoutsState        `json:"timeouts"`
}
