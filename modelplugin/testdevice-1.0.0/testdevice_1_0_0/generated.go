// Code generated by GENERATOR. DO NOT EDIT.
/*
Package testdevice_1_0_0 is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /home/scondon/go/pkg/mod/github.com/openconfig/ygot@v0.6.1-0.20200103195725-e3c44fa43926/genutil/names.go
using the following YANG input files:
	- test1@2018-02-20.yang
Imported modules were sourced from:
	- yang/...
*/
package testdevice_1_0_0

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

var (
	SchemaTree map[string]*yang.Entry
)

func init() {
	var err error
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " +  err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root: &Device{},
		SchemaTree: uzp,
		Unmarshal: Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn )
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	Cont1A	*Test1_Cont1A	`path:"cont1a" module:"test1"`
	Cont1BState	*Test1_Cont1BState	`path:"cont1b-state" module:"test1"`
	LeafAtTopLevel	*string	`path:"leafAtTopLevel" module:"test1"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A represents the /test1/cont1a YANG schema element.
type Test1_Cont1A struct {
	Cont2A	*Test1_Cont1A_Cont2A	`path:"cont2a" module:"test1"`
	Leaf1A	*string	`path:"leaf1a" module:"test1"`
	List2A	map[string]*Test1_Cont1A_List2A	`path:"list2a" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A) IsYANGGoStruct() {}

// NewList2A creates a new entry in the List2A list of the
// Test1_Cont1A struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1A) NewList2A(Name string) (*Test1_Cont1A_List2A, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2A == nil {
		t.List2A = make(map[string]*Test1_Cont1A_List2A)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2A[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2A", key)
	}

	t.List2A[key] = &Test1_Cont1A_List2A{
		Name: &Name,
	}

	return t.List2A[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_Cont2A represents the /test1/cont1a/cont2a YANG schema element.
type Test1_Cont1A_Cont2A struct {
	Leaf2A	*uint8	`path:"leaf2a" module:"test1"`
	Leaf2B	*float64	`path:"leaf2b" module:"test1"`
	Leaf2C	*string	`path:"leaf2c" module:"test1"`
	Leaf2D	*float64	`path:"leaf2d" module:"test1"`
	Leaf2E	[]int16	`path:"leaf2e" module:"test1"`
	Leaf2F	Binary	`path:"leaf2f" module:"test1"`
	Leaf2G	*bool	`path:"leaf2g" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_Cont2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_Cont2A) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_Cont2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_Cont2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_Cont2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_List2A represents the /test1/cont1a/list2a YANG schema element.
type Test1_Cont1A_List2A struct {
	Name	*string	`path:"name" module:"test1"`
	TxPower	*uint16	`path:"tx-power" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_List2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_List2A) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1A_List2A struct, which is a YANG list entry.
func (t *Test1_Cont1A_List2A) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_List2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_List2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_List2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState represents the /test1/cont1b-state YANG schema element.
type Test1_Cont1BState struct {
	Leaf2D	*uint16	`path:"leaf2d" module:"test1"`
	List2B	map[uint8]*Test1_Cont1BState_List2B	`path:"list2b" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState) IsYANGGoStruct() {}

// NewList2B creates a new entry in the List2B list of the
// Test1_Cont1BState struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1BState) NewList2B(Index uint8) (*Test1_Cont1BState_List2B, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2B == nil {
		t.List2B = make(map[uint8]*Test1_Cont1BState_List2B)
	}

	key := Index

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2B[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2B", key)
	}

	t.List2B[key] = &Test1_Cont1BState_List2B{
		Index: &Index,
	}

	return t.List2B[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState_List2B represents the /test1/cont1b-state/list2b YANG schema element.
type Test1_Cont1BState_List2B struct {
	Index	*uint8	`path:"index" module:"test1"`
	Leaf3C	*string	`path:"leaf3c" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState_List2B implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState_List2B) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1BState_List2B struct, which is a YANG list entry.
func (t *Test1_Cont1BState_List2B) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Index == nil {
		return nil, fmt.Errorf("nil value for key Index")
	}

	return map[string]interface{}{
		"index": *t.Index,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState_List2B) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState_List2B"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState_List2B) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }



var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5c, 0x4b, 0x6f, 0xe2, 0x48,
		0x10, 0xbe, 0xf3, 0x2b, 0x2c, 0x1f, 0x57, 0x41, 0xb1, 0x4d, 0x12, 0x65, 0xb8, 0x91, 0xc9, 0x8e,
		0x56, 0x9a, 0x64, 0x77, 0x34, 0x13, 0xed, 0x61, 0x23, 0xb4, 0x6a, 0x4c, 0xc3, 0xb4, 0xd6, 0xb4,
		0x91, 0xdd, 0x9e, 0x01, 0x8d, 0xf8, 0xef, 0x2b, 0x3f, 0x20, 0xf8, 0xd9, 0x55, 0x4e, 0x1c, 0x08,
		0xa9, 0x53, 0x14, 0xbb, 0xda, 0xdd, 0xf5, 0xe8, 0xaf, 0xca, 0x9f, 0xab, 0xf9, 0xd5, 0x33, 0x0c,
		0xc3, 0x30, 0xff, 0x64, 0x0b, 0x6e, 0x0e, 0x0d, 0x73, 0xca, 0x7f, 0x08, 0x97, 0x9b, 0x67, 0xe9,
		0xd5, 0xcf, 0x42, 0x4e, 0xcd, 0xa1, 0x61, 0x67, 0xff, 0x7e, 0xf4, 0xe5, 0x4c, 0xcc, 0xcd, 0xa1,
		0x61, 0x65, 0x17, 0x6e, 0x45, 0x60, 0x0e, 0x8d, 0xf4, 0x11, 0xc9, 0x05, 0xd7, 0x97, 0xca, 0x66,
		0xb9, 0x6b, 0xb9, 0xc7, 0x67, 0xf7, 0xcf, 0xf2, 0x77, 0xf3, 0xd3, 0xec, 0x2e, 0x17, 0xa7, 0xdb,
		0xdd, 0xf8, 0x12, 0xf0, 0x99, 0x58, 0x95, 0x66, 0xc9, 0xcd, 0xa4, 0xec, 0xc2, 0x2c, 0xc9, 0xdd,
		0x6f, 0x7e, 0x14, 0xb8, 0xbc, 0x72, 0x64, 0xba, 0x12, 0xbe, 0xfe, 0xe9, 0x07, 0xf1, 0x62, 0xcc,
		0x65, 0x3a, 0xc9, 0x59, 0xb5, 0xe0, 0x1f, 0x2c, 0x1c, 0x05, 0xf3, 0x68, 0xc1, 0xa5, 0x32, 0x87,
		0x86, 0x0a, 0x22, 0x5e, 0x23, 0xb8, 0x27, 0x15, 0xaf, 0xa9, 0x24, 0xb4, 0xc9, 0x5d, 0xd9, 0x14,
		0x34, 0x2d, 0x1a, 0x38, 0x67, 0x68, 0x87, 0xd5, 0x2b, 0xb2, 0x6f, 0x70, 0x87, 0xd5, 0x69, 0x51,
		0x6d, 0x78, 0xad, 0x03, 0x20, 0x8e, 0x80, 0x39, 0x04, 0xea, 0x18, 0xb4, 0x83, 0xd0, 0x8e, 0x02,
		0x3b, 0xac, 0xda, 0x71, 0x35, 0x0e, 0xd4, 0x3a, 0x72, 0x27, 0xe0, 0x71, 0x36, 0x6b, 0x70, 0x68,
		0xc9, 0x9c, 0x99, 0xbc, 0x46, 0x99, 0x5b, 0x3e, 0x63, 0x91, 0x97, 0xe8, 0xe2, 0xe8, 0x64, 0xb3,
		0x60, 0xb0, 0x34, 0x62, 0xba, 0xa0, 0xc0, 0x04, 0x07, 0x2e, 0x48, 0xb0, 0xc1, 0xd2, 0x3a, 0x68,
		0x5a, 0x07, 0x0f, 0x3a, 0x88, 0x9a, 0x83, 0x49, 0x13, 0x54, 0xbb, 0xd9, 0x1e, 0xd6, 0x4b, 0x8e,
		0xb3, 0x73, 0x24, 0xa4, 0xba, 0x86, 0x98, 0x3a, 0x0b, 0x8a, 0x4b, 0x80, 0xe8, 0x57, 0x26, 0xe7,
		0xf1, 0xc3, 0x1f, 0x41, 0x26, 0x82, 0xb9, 0x2e, 0x79, 0xf0, 0xbd, 0x90, 0x60, 0x5f, 0x23, 0xa3,
		0xb9, 0x34, 0xec, 0x6f, 0xe6, 0x45, 0xbc, 0x1e, 0x12, 0x6b, 0xc7, 0x7d, 0x0a, 0x98, 0xab, 0x84,
		0x2f, 0x6f, 0xc5, 0x5c, 0xa8, 0x30, 0x9e, 0x18, 0x3c, 0x7e, 0x73, 0x86, 0x30, 0x05, 0x5b, 0xbd,
		0xba, 0x29, 0x06, 0xaf, 0x68, 0x8a, 0xde, 0x0b, 0x1a, 0xec, 0xd8, 0x23, 0x8c, 0x42, 0xec, 0xc9,
		0x16, 0xc7, 0x17, 0x63, 0x5a, 0xa9, 0x71, 0xaf, 0xdd, 0xf8, 0x06, 0x5f, 0xa4, 0x49, 0x7d, 0x82,
		0x2c, 0x02, 0x26, 0x2f, 0x9d, 0xd8, 0x6d, 0x4a, 0xec, 0x6f, 0x36, 0xb1, 0x4f, 0xb9, 0x2b, 0x16,
		0xcc, 0xbb, 0xba, 0x40, 0x24, 0x77, 0xdb, 0x01, 0xc8, 0x96, 0x76, 0xdc, 0x80, 0x4a, 0x02, 0xac,
		0xc5, 0x4e, 0x06, 0xaf, 0x1d, 0xcb, 0xb2, 0x5e, 0xd1, 0x1a, 0xc7, 0x8e, 0xd8, 0x2e, 0x12, 0xb1,
		0xdd, 0x97, 0x46, 0x6c, 0x87, 0x10, 0xfb, 0xcd, 0x22, 0x76, 0xa8, 0x02, 0x21, 0xe7, 0x18, 0xb8,
		0xbe, 0xee, 0x2a, 0x8e, 0xa7, 0xc8, 0x38, 0x9e, 0x12, 0xa5, 0x40, 0x71, 0x4c, 0x95, 0x07, 0x55,
		0x1e, 0x54, 0x79, 0x1c, 0xa8, 0xf2, 0xe0, 0x48, 0xc4, 0xe6, 0x84, 0xd8, 0x84, 0xd8, 0x5b, 0x3b,
		0x0b, 0xa9, 0xec, 0x2b, 0x04, 0x5a, 0x3b, 0x6f, 0x15, 0x77, 0xed, 0xb6, 0xb8, 0xfb, 0x7c, 0xac,
		0xb1, 0x4e, 0x09, 0x79, 0x4f, 0x8e, 0xa4, 0xd3, 0x6c, 0xa9, 0x3b, 0x11, 0xaa, 0x91, 0x52, 0x01,
		0x6c, 0x5b, 0xdd, 0x0b, 0xf9, 0xbb, 0xc7, 0xe3, 0x0d, 0x1f, 0xeb, 0x2a, 0x23, 0xcf, 0x03, 0xec,
		0x97, 0x7b, 0xb6, 0xc2, 0x0f, 0xfa, 0x2b, 0x98, 0xf2, 0x80, 0x4f, 0x6f, 0xd6, 0xd9, 0x90, 0xae,
		0xb2, 0xcb, 0x0c, 0x99, 0x5d, 0x66, 0x94, 0x5d, 0x28, 0xbb, 0x6c, 0xed, 0x3c, 0x11, 0x92, 0x05,
		0x6b, 0x44, 0x7a, 0xf9, 0x00, 0x10, 0xbd, 0xe3, 0x72, 0xae, 0xbe, 0x9f, 0x4a, 0x5d, 0xef, 0x50,
		0x7a, 0x39, 0x8c, 0x2d, 0x8e, 0xbd, 0xac, 0x9f, 0x23, 0x81, 0x77, 0x4e, 0xc0, 0x4b, 0xc0, 0xbb,
		0x03, 0x5e, 0xdf, 0xf7, 0x38, 0x93, 0x18, 0x1a, 0xc6, 0x6e, 0x1b, 0xc8, 0xa8, 0x1e, 0xa8, 0x91,
		0x94, 0xbe, 0x62, 0xf1, 0x9e, 0x6d, 0x6e, 0x85, 0x0a, 0xdd, 0xef, 0x7c, 0xc1, 0x96, 0x2c, 0x01,
		0x7a, 0xf3, 0x5c, 0xf1, 0x50, 0xd9, 0xe7, 0x69, 0xe7, 0xe0, 0x79, 0x63, 0x3f, 0x5b, 0x3a, 0x5a,
		0x05, 0x91, 0xab, 0x64, 0x66, 0x8c, 0x87, 0x78, 0xf0, 0xbf, 0x1f, 0xe3, 0xc1, 0xa3, 0xe4, 0x8f,
		0x33, 0xaa, 0xf6, 0x58, 0x59, 0x95, 0x0a, 0x35, 0x92, 0xdd, 0x66, 0x03, 0xda, 0xee, 0x32, 0xb9,
		0xe6, 0xb6, 0x3b, 0x8b, 0xda, 0xee, 0xf0, 0xbb, 0x0a, 0x17, 0x72, 0xda, 0xdd, 0x03, 0xa7, 0xe1,
		0x9f, 0xe8, 0xf7, 0x06, 0x19, 0x60, 0x7d, 0x02, 0x7b, 0x9b, 0x40, 0x80, 0x20, 0x2a, 0xf7, 0xee,
		0x72, 0xee, 0x25, 0x50, 0xbe, 0x45, 0xaa, 0xdd, 0xc0, 0xde, 0x7d, 0x3a, 0x57, 0xd1, 0xb6, 0x3a,
		0xd4, 0xb1, 0x25, 0x6a, 0x8e, 0x9f, 0x03, 0x40, 0x22, 0x04, 0xf5, 0xfd, 0x66, 0x72, 0xd4, 0xf7,
		0x7b, 0xec, 0x7d, 0xbf, 0x59, 0xa2, 0x02, 0x56, 0x7b, 0x89, 0x34, 0xd5, 0x7a, 0x54, 0xeb, 0xb5,
		0xff, 0x78, 0xfc, 0xee, 0xde, 0xb2, 0x2f, 0xe8, 0x25, 0x7b, 0x6b, 0x8a, 0x6b, 0x7a, 0xc7, 0x8e,
		0xd5, 0x52, 0xab, 0xfe, 0xd2, 0xff, 0xc9, 0x03, 0x38, 0xee, 0xee, 0x46, 0x10, 0xf6, 0x12, 0xf6,
		0xee, 0x9d, 0xa1, 0x40, 0x7d, 0x3f, 0xbb, 0xa2, 0xbe, 0x05, 0xa2, 0x37, 0xdf, 0x03, 0xf4, 0xa2,
		0x2a, 0xe4, 0xcf, 0x7c, 0xad, 0xa9, 0x6c, 0x61, 0xdf, 0xe1, 0x50, 0xdf, 0xdf, 0x0a, 0xdf, 0xdd,
		0x80, 0x59, 0x40, 0xd7, 0xda, 0x84, 0x01, 0xd5, 0x7d, 0x40, 0x5d, 0xb0, 0x55, 0x9f, 0x6f, 0x57,
		0x03, 0x00, 0x89, 0x36, 0x90, 0x9a, 0x83, 0xd3, 0x0b, 0xb3, 0x83, 0xac, 0x0a, 0xfa, 0x2a, 0xd9,
		0x25, 0x37, 0xd8, 0xf8, 0xce, 0xab, 0xe3, 0x06, 0xe3, 0x18, 0x43, 0x70, 0x83, 0x8d, 0xa7, 0x76,
		0x35, 0xca, 0x34, 0x29, 0x51, 0x75, 0x68, 0xb9, 0x7e, 0xd5, 0xf9, 0xe5, 0x3e, 0x2d, 0x6a, 0x6f,
		0x41, 0xe9, 0x91, 0xeb, 0x49, 0x3f, 0x54, 0x4c, 0x71, 0xcd, 0xc1, 0xec, 0xad, 0x14, 0xf2, 0x78,
		0xb6, 0x43, 0xc7, 0xb3, 0xab, 0x78, 0xe2, 0x86, 0x76, 0x5a, 0x58, 0x1b, 0x2d, 0x98, 0x27, 0x76,
		0x88, 0xa6, 0x79, 0x35, 0x9e, 0x58, 0x5b, 0xf5, 0x01, 0xaa, 0x3d, 0x60, 0x95, 0x77, 0x24, 0x2c,
		0xb1, 0x6d, 0x81, 0x7b, 0x2e, 0xdf, 0x32, 0x53, 0xec, 0x74, 0xac, 0xe6, 0xa1, 0xc8, 0xe2, 0x09,
		0x90, 0x2c, 0x9e, 0x3c, 0x97, 0x2c, 0x26, 0x14, 0xea, 0x9c, 0x2c, 0x16, 0x72, 0xca, 0x57, 0xf0,
		0x7a, 0x35, 0x15, 0x27, 0xca, 0x82, 0x28, 0x0b, 0xfa, 0xd9, 0x07, 0x58, 0x12, 0xa0, 0x86, 0xac,
		0xa7, 0x7c, 0x78, 0x79, 0x49, 0x6c, 0xf1, 0xb6, 0x96, 0x1f, 0x20, 0x8f, 0x78, 0x0e, 0xe8, 0x88,
		0x27, 0xc1, 0x2e, 0x7d, 0xa5, 0x23, 0xae, 0x98, 0xb8, 0xe2, 0x0e, 0xb8, 0xe2, 0xa6, 0xba, 0xb6,
		0x7b, 0xb2, 0x58, 0x23, 0x7c, 0x60, 0x1a, 0x34, 0xe3, 0xf0, 0xce, 0x1b, 0xdf, 0xe9, 0x9a, 0x69,
		0xc5, 0x9b, 0x6f, 0xf1, 0x13, 0x52, 0x46, 0xf4, 0xe6, 0xd0, 0x8c, 0x68, 0x35, 0x27, 0x09, 0x51,
		0x00, 0x42, 0x8e, 0xc6, 0xb9, 0x7a, 0xa4, 0x1e, 0xfc, 0xe5, 0x1d, 0xff, 0xc1, 0xbd, 0x7a, 0x7a,
		0xb4, 0x20, 0x57, 0x4d, 0x90, 0x5a, 0xf4, 0xfb, 0x95, 0x45, 0xbf, 0xd7, 0xe6, 0x4a, 0x7d, 0x6e,
		0x6c, 0xca, 0x85, 0xe6, 0x17, 0xa6, 0x14, 0x0f, 0x64, 0x6d, 0xf2, 0x33, 0x1f, 0x47, 0xfd, 0x7f,
		0xc6, 0xbf, 0x06, 0x9b, 0xfe, 0xa3, 0xd5, 0xff, 0x30, 0xfe, 0xad, 0xbc, 0xee, 0x71, 0x5d, 0x74,
		0xf4, 0xf6, 0xf4, 0xa8, 0x8b, 0x5b, 0x53, 0x84, 0x9f, 0xd8, 0x7f, 0xfc, 0xab, 0xef, 0x97, 0xad,
		0x57, 0x8c, 0x65, 0x73, 0xff, 0x56, 0x2e, 0x62, 0x6f, 0xd3, 0xdf, 0x5a, 0x4d, 0x27, 0xec, 0x6d,
		0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x3b, 0x7c, 0xd3, 0x8a,
		0x55, 0x00, 0x00,
	}
)


// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
var ΛEnumTypes = map[string][]reflect.Type{
}

