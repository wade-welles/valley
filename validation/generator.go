package validation

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"io"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/seeruk/valley/valley"
)

// Generator is a type used to generate validation code.
type Generator struct {
	constraints map[string]valley.ConstraintGenerator

	cb   *bytes.Buffer
	ipts map[valley.Import]struct{}
}

// NewGenerator returns a new Generator instance.
func NewGenerator(constraints map[string]valley.ConstraintGenerator) *Generator {
	return &Generator{
		constraints: constraints,
		cb:          &bytes.Buffer{},
		ipts:        make(map[valley.Import]struct{}),
	}
}

// Generate attempts to generate the code (returned as bytes) to validate code in the given package,
// using the given configuration.
func (g *Generator) Generate(config valley.Config, file valley.Source) ([]byte, error) {
	typeNames := make([]string, 0, len(config.Types))
	for typeName := range config.Types {
		typeNames = append(typeNames, typeName)
	}

	// Ensure we generate methods in the same order each time.
	sort.Strings(typeNames)

	for _, typeName := range typeNames {
		err := g.generateType(config, file, typeName)
		if err != nil {
			// TODO: Wrap.
			return nil, err
		}
	}

	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, "// Code generated by valley. DO NOT EDIT.")
	fmt.Fprintf(buf, "package %s\n", file.Package)
	fmt.Fprintln(buf)
	fmt.Fprintln(buf, "import \"strconv\"")
	fmt.Fprintln(buf, "import \"github.com/seeruk/valley/valley\"")

	for ipt := range g.ipts {
		if ipt.Alias != "" {
			fmt.Fprintf(buf, "import %s \"%s\"\n", ipt.Alias, ipt.Path)
		} else {
			fmt.Fprintf(buf, "import \"%s\"\n", ipt.Path)
		}
	}

	fmt.Fprintln(buf)

	_, err := io.Copy(buf, g.cb)
	if err != nil {
		return nil, errors.New("failed to copy code buffer contents to generate buffer")
	}

	return buf.Bytes(), nil
}

// generateType generates the entire Validate method for a particular type (found in the given
// package with the given type name.
func (g *Generator) generateType(config valley.Config, file valley.Source, typeName string) error {
	typ := config.Types[typeName]

	s, ok := file.Structs[typeName]
	if !ok {
		// TODO: Is this correct?
		return nil
	}

	// Figure out an "okay" receiver name, based on the first letter of the type.
	firstRune, _ := utf8.DecodeRuneInString(typeName)
	receiver := strings.ToLower(string(firstRune))

	mm, ok := file.Methods[typeName]
	if ok && len(mm) > 0 {
		receiver = mm[0].Receiver
	}

	g.wc("// Validate validates this %s.\n", typeName)
	g.wc("// This method was generated by Valley.\n")
	g.wc("func (%s %s) Validate() []valley.ConstraintViolation {\n", receiver, typeName)
	g.wc("	var violations []valley.ConstraintViolation\n")
	g.wc("\n")
	g.wc("	path := valley.NewPath()\n")
	g.wc("	path.Write(\".\")\n\n")

	ctx := valley.Context{
		FileSet:  file.FileSet,
		TypeName: typeName,
		Receiver: receiver,
		VarName:  receiver,
	}

	for _, constraint := range typ.Constraints {
		value := valley.Value{
			Name: s.Name,
			Type: s.Node,
		}

		err := g.generateConstraint(ctx, constraint, value)
		if err != nil {
			// TODO: Wrap.
			return err
		}
	}

	fieldNames := make([]string, 0, len(typ.Fields))
	for fieldName := range typ.Fields {
		fieldNames = append(fieldNames, fieldName)
	}

	// Ensure that each field's validation is generated in the same order each time.
	sort.Strings(fieldNames)

	for _, fieldName := range fieldNames {
		fieldConfig := typ.Fields[fieldName]

		f, ok := s.Fields[fieldName]
		if !ok {
			return fmt.Errorf("field %q does not exist in Go source", fieldName)
		}

		ctx.FieldName = fieldName
		ctx.VarName = fmt.Sprintf("%s.%s", receiver, fieldName)
		ctx.Path = fmt.Sprintf("\"%s\"", fieldName)
		ctx.BeforeViolation = fmt.Sprintf("size := path.Write(%s)", ctx.Path)
		ctx.AfterViolation = "path.TruncateRight(size)"

		err := g.generateField(ctx, fieldConfig, f)
		if err != nil {
			// TODO: Wrap?
			return err
		}
	}

	g.wc("	return violations\n")
	g.wc("}\n\n")

	return nil
}

// generateField generates all of the code for a specific field.
func (g *Generator) generateField(ctx valley.Context, fieldConfig valley.FieldConfig, value valley.Value) error {
	err := g.generateFieldConstraints(ctx, fieldConfig, value)
	if err != nil {
		// TODO: Wrap.
		return err
	}

	err = g.generateFieldElementsConstraints(ctx, fieldConfig, value)
	if err != nil {
		// TODO: Wrap.
		return err
	}

	g.wc("\n")

	return nil
}

// generateFieldConstraints generates the code for constraints that apply directly to a specific
// field (i.e. not including things like the element constraints).
func (g *Generator) generateFieldConstraints(ctx valley.Context, fieldConfig valley.FieldConfig, value valley.Value) error {
	// Generate the constraint code for the field as a whole.
	for _, constraintConfig := range fieldConfig.Constraints {
		err := g.generateConstraint(ctx, constraintConfig, value)
		if err != nil {
			// TODO: Wrap.
			return err
		}
	}

	return nil
}

// generateFieldElementsConstraints generate the constraint code for each element of an
// array/map/slice field.
// TODO: Only works with array/slice currently, need more info on Context to handle maps?
func (g *Generator) generateFieldElementsConstraints(ctx valley.Context, fieldConfig valley.FieldConfig, value valley.Value) error {
	ctx.Path = fmt.Sprintf("\"%s.[\" + strconv.Itoa(i) + \"]\"", ctx.FieldName)
	ctx.BeforeViolation = fmt.Sprintf("size := path.Write(%s)", ctx.Path)

	if len(fieldConfig.Elements) == 0 {
		return nil
	}

	g.wc("	for i, element := range %s {\n", ctx.VarName)

	for _, constraintConfig := range fieldConfig.Elements {
		elementValue := ctx.Clone()
		elementValue.VarName = "element"

		arrayType, ok := value.Type.(*ast.ArrayType)
		if !ok {
			return errors.New("config for elements applied to non-iterable type")
		}

		elementField := valley.Value{
			Name: value.Name,
			Type: arrayType.Elt,
		}

		err := g.generateConstraint(elementValue, constraintConfig, elementField)
		if err != nil {
			// TODO: Wrap.
			return err
		}
	}

	g.wc("	}\n\n")

	return nil
}

// generateConstraint attempts to generate the code for a particular constraint.
func (g *Generator) generateConstraint(ctx valley.Context, constraintConfig valley.ConstraintConfig, value valley.Value) error {
	constraint, ok := g.constraints[constraintConfig.Name]
	if !ok {
		return fmt.Errorf("unknown validation constraint: %q", constraintConfig.Name)
	}

	output, err := constraint(ctx, value.Type, constraintConfig.Opts)
	if err != nil {
		selector := ctx.TypeName
		if ctx.FieldName != "" {
			selector += "." + ctx.FieldName
		}

		pos := ctx.FileSet.Position(constraintConfig.Pos)

		return fmt.Errorf("failed to generate code for %s's %q constraint on line %d, col %d: %v",
			selector, constraintConfig.Name, pos.Line, pos.Column, err)
	}

	for _, ipt := range output.Imports {
		g.ipts[ipt] = struct{}{}
	}

	g.wc(output.Code)
	g.wc("\n")

	return nil
}

// wc writes code to the code buffer.
func (g *Generator) wc(format string, a ...interface{}) {
	fmt.Fprintf(g.cb, format, a...)
}
