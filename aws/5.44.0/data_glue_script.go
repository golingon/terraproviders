// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datagluescript "github.com/golingon/terraproviders/aws/5.44.0/datagluescript"
)

// NewDataGlueScript creates a new instance of [DataGlueScript].
func NewDataGlueScript(name string, args DataGlueScriptArgs) *DataGlueScript {
	return &DataGlueScript{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGlueScript)(nil)

// DataGlueScript represents the Terraform data resource aws_glue_script.
type DataGlueScript struct {
	Name string
	Args DataGlueScriptArgs
}

// DataSource returns the Terraform object type for [DataGlueScript].
func (gs *DataGlueScript) DataSource() string {
	return "aws_glue_script"
}

// LocalName returns the local name for [DataGlueScript].
func (gs *DataGlueScript) LocalName() string {
	return gs.Name
}

// Configuration returns the configuration (args) for [DataGlueScript].
func (gs *DataGlueScript) Configuration() interface{} {
	return gs.Args
}

// Attributes returns the attributes for [DataGlueScript].
func (gs *DataGlueScript) Attributes() dataGlueScriptAttributes {
	return dataGlueScriptAttributes{ref: terra.ReferenceDataResource(gs)}
}

// DataGlueScriptArgs contains the configurations for aws_glue_script.
type DataGlueScriptArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Language: string, optional
	Language terra.StringValue `hcl:"language,attr"`
	// DagEdge: min=1
	DagEdge []datagluescript.DagEdge `hcl:"dag_edge,block" validate:"min=1"`
	// DagNode: min=1
	DagNode []datagluescript.DagNode `hcl:"dag_node,block" validate:"min=1"`
}
type dataGlueScriptAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_glue_script.
func (gs dataGlueScriptAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("id"))
}

// Language returns a reference to field language of aws_glue_script.
func (gs dataGlueScriptAttributes) Language() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("language"))
}

// PythonScript returns a reference to field python_script of aws_glue_script.
func (gs dataGlueScriptAttributes) PythonScript() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("python_script"))
}

// ScalaCode returns a reference to field scala_code of aws_glue_script.
func (gs dataGlueScriptAttributes) ScalaCode() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("scala_code"))
}

func (gs dataGlueScriptAttributes) DagEdge() terra.ListValue[datagluescript.DagEdgeAttributes] {
	return terra.ReferenceAsList[datagluescript.DagEdgeAttributes](gs.ref.Append("dag_edge"))
}

func (gs dataGlueScriptAttributes) DagNode() terra.ListValue[datagluescript.DagNodeAttributes] {
	return terra.ReferenceAsList[datagluescript.DagNodeAttributes](gs.ref.Append("dag_node"))
}
