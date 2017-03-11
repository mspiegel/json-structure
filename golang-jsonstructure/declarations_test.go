package jsonstructure

import (
	"encoding/json"
	"testing"

	multierror "github.com/mspiegel/go-multierror"
	"github.com/shopspring/decimal"
)

var negOneInt = -1
var negOneDec = decimal.NewFromFloat(-1.0)
var zero = 0
var posOneInt = 1
var affirmative = true
var foobar = "foobar"

func TestUnmarshalJSON(t *testing.T) {
	var td TypeDecl
	text := `{
		"type": "number",
		"default": 0
	}`
	err := json.Unmarshal([]byte(text), &td)
	if err != nil {
		t.Error("Unmarshal error", err)
	}
	if td.Default != json.Number("0") {
		t.Error("Default value unmarshal error", td.Default)
	}
}

func TestValidateJSONStructureSuccess(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "struct"
	structure.Definition.Main.Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Fields["foo"] = &TypeDecl{}
	structure.Definition.Main.Fields["foo"].Type = "boolean"
	err := structure.ValidateStructure()
	if err != nil {
		t.Error("Validation failure", err)
	}
	pattern := "[0-9]+"
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "struct"
	structure.Definition.Main.Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Fields["foo"] = &TypeDecl{}
	structure.Definition.Main.Fields["foo"].Type = "string"
	structure.Definition.Main.Fields["foo"].PatternRaw = &pattern
	err = structure.ValidateStructure()
	if err != nil {
		t.Error("Validation failure", err)
	}
}

func TestValidateDefaultSuccess(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	structure.Definition.Main.DefaultRaw = []byte("[]")
	structure.Definition.Main.Default = []interface{}{}
	structure.Definition.Main.Items = &TypeDecl{}
	structure.Definition.Main.Items.Type = "integer"
	structure.Definition.Main.Items.DefaultRaw = []byte("1")
	structure.Definition.Main.Items.Default = json.Number("1")
	err := structure.ValidateStructure()
	if err != nil {
		t.Error("Validation failure", err)
	}
}

func TestValidateDefaultFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	structure.Definition.Main.DefaultRaw = []byte("[]")
	structure.Definition.Main.Default = []interface{}{}
	structure.Definition.Main.Items = &TypeDecl{}
	structure.Definition.Main.Items.Type = "integer"
	structure.Definition.Main.Items.DefaultRaw = []byte(`"hello world"`)
	structure.Definition.Main.Items.Default = "hello world"
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateJSONStructureFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Types = make(map[string]*TypeDecl)
	structure.Definition.Types["foo"] = nil
	structure.Definition.Types["bar"] = nil
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure when 'type' is empty")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "foo"
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure when 'type' is foo")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "foo"
	structure.Definition.Main.Minimum = &decimal.Zero
	structure.Definition.Main.Nullable = &affirmative
	structure.Definition.Types = make(map[string]*TypeDecl)
	structure.Definition.Types["foo"] = &TypeDecl{}
	structure.Definition.Types["foo"].Type = "integer"
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure when 'type' is foo")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "boolean"
	structure.Definition.Main.MultipleOf = &decimal.Zero
	structure.Definition.Main.Minimum = &decimal.Zero
	structure.Definition.Main.Maximum = &decimal.Zero
	structure.Definition.Main.ExclusiveMinimum = &decimal.Zero
	structure.Definition.Main.ExclusiveMaximum = &decimal.Zero
	structure.Definition.Main.PatternRaw = &foobar
	structure.Definition.Main.MinLength = &zero
	structure.Definition.Main.MaxLength = &zero
	structure.Definition.Main.Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Items = &TypeDecl{}
	structure.Definition.Main.MinItems = &zero
	structure.Definition.Main.MaxItems = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	if len(err.(*multierror.Error).Errors) != 12 {
		t.Error("Expected 12 errors")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	structure.Definition.Main.MultipleOf = &negOneDec
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateNumberFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	structure.Definition.Main.Minimum = &decimal.Zero
	structure.Definition.Main.Maximum = &decimal.Zero
	structure.Definition.Main.ExclusiveMinimum = &decimal.Zero
	structure.Definition.Main.ExclusiveMaximum = &decimal.Zero
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	structure.Definition.Main.Minimum = &decimal.Zero
	structure.Definition.Main.Maximum = &negOneDec
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateStringFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.MinLength = &negOneInt
	structure.Definition.Main.MaxLength = &negOneInt
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.MinLength = &posOneInt
	structure.Definition.Main.MaxLength = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	badRegex := "(()"
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.PatternRaw = &badRegex
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateStructFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "struct"
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateUnionFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "union"
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateArrayFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	structure.Definition.Main.Items = &TypeDecl{}
	structure.Definition.Main.Items.Type = "boolean"
	structure.Definition.Main.MinItems = &negOneInt
	structure.Definition.Main.MaxItems = &negOneInt
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	structure.Definition.Main.Items = &TypeDecl{}
	structure.Definition.Main.Items.Type = "boolean"
	structure.Definition.Main.MinItems = &posOneInt
	structure.Definition.Main.MaxItems = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateCycleFailure(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Types = make(map[string]*TypeDecl)
	structure.Definition.Types["a"] = &TypeDecl{}
	structure.Definition.Types["b"] = &TypeDecl{}
	structure.Definition.Types["c"] = &TypeDecl{}
	structure.Definition.Types["a"].Type = "b"
	structure.Definition.Types["b"].Type = "a"
	structure.Definition.Types["c"].Type = "c"
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "boolean"
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}

func TestValidateFormatFailure(t *testing.T) {
	format := "foobar"
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	structure.Definition.Main.Format = &format
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	format = "hostname"
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	structure.Definition.Main.Format = &format
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	pattern := "\\d+"
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.PatternRaw = &pattern
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Validation failure", err)
	}
	t.Log(err)
}
