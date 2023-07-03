// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataTransformDecode creates a new instance of [DataTransformDecode].
func NewDataTransformDecode(name string, args DataTransformDecodeArgs) *DataTransformDecode {
	return &DataTransformDecode{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTransformDecode)(nil)

// DataTransformDecode represents the Terraform data resource vault_transform_decode.
type DataTransformDecode struct {
	Name string
	Args DataTransformDecodeArgs
}

// DataSource returns the Terraform object type for [DataTransformDecode].
func (td *DataTransformDecode) DataSource() string {
	return "vault_transform_decode"
}

// LocalName returns the local name for [DataTransformDecode].
func (td *DataTransformDecode) LocalName() string {
	return td.Name
}

// Configuration returns the configuration (args) for [DataTransformDecode].
func (td *DataTransformDecode) Configuration() interface{} {
	return td.Args
}

// Attributes returns the attributes for [DataTransformDecode].
func (td *DataTransformDecode) Attributes() dataTransformDecodeAttributes {
	return dataTransformDecodeAttributes{ref: terra.ReferenceDataResource(td)}
}

// DataTransformDecodeArgs contains the configurations for vault_transform_decode.
type DataTransformDecodeArgs struct {
	// BatchInput: list of map of string, optional
	BatchInput terra.ListValue[terra.MapValue[terra.StringValue]] `hcl:"batch_input,attr"`
	// BatchResults: list of map of string, optional
	BatchResults terra.ListValue[terra.MapValue[terra.StringValue]] `hcl:"batch_results,attr"`
	// DecodedValue: string, optional
	DecodedValue terra.StringValue `hcl:"decoded_value,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// RoleName: string, required
	RoleName terra.StringValue `hcl:"role_name,attr" validate:"required"`
	// Transformation: string, optional
	Transformation terra.StringValue `hcl:"transformation,attr"`
	// Tweak: string, optional
	Tweak terra.StringValue `hcl:"tweak,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}
type dataTransformDecodeAttributes struct {
	ref terra.Reference
}

// BatchInput returns a reference to field batch_input of vault_transform_decode.
func (td dataTransformDecodeAttributes) BatchInput() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](td.ref.Append("batch_input"))
}

// BatchResults returns a reference to field batch_results of vault_transform_decode.
func (td dataTransformDecodeAttributes) BatchResults() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](td.ref.Append("batch_results"))
}

// DecodedValue returns a reference to field decoded_value of vault_transform_decode.
func (td dataTransformDecodeAttributes) DecodedValue() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("decoded_value"))
}

// Id returns a reference to field id of vault_transform_decode.
func (td dataTransformDecodeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_transform_decode.
func (td dataTransformDecodeAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_transform_decode.
func (td dataTransformDecodeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("path"))
}

// RoleName returns a reference to field role_name of vault_transform_decode.
func (td dataTransformDecodeAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("role_name"))
}

// Transformation returns a reference to field transformation of vault_transform_decode.
func (td dataTransformDecodeAttributes) Transformation() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("transformation"))
}

// Tweak returns a reference to field tweak of vault_transform_decode.
func (td dataTransformDecodeAttributes) Tweak() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("tweak"))
}

// Value returns a reference to field value of vault_transform_decode.
func (td dataTransformDecodeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("value"))
}
