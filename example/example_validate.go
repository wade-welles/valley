// Code generated by valley. DO NOT EDIT.
package main

import "fmt"
import "strconv"
import "github.com/seeruk/valley/valley"
import math "math"

// Reference imports to suppress errors if they aren't otherwise used
var _ = fmt.Sprintf
var _ = strconv.Itoa

// Validate validates this Example.
// This method was generated by Valley.
func (e Example) Validate(path *valley.Path) []valley.ConstraintViolation {
	var violations []valley.ConstraintViolation

	path.Write(".")

	{
		// MutuallyExclusive uses it's own block to lock down nonEmpty's scope.
		var nonEmpty []string

		if !(len(e.Text) == 0) {
			nonEmpty = append(nonEmpty, "Text")
		}

		if !(len(e.Texts) == 0) {
			nonEmpty = append(nonEmpty, "Texts")
		}

		if len(nonEmpty) > 1 {

			violations = append(violations, valley.ConstraintViolation{
				Field:   path.String(),
				Message: "fields are mutually exclusive",
				Details: map[string]interface{}{
					"fields": nonEmpty,
				},
			})

		}
	}

	if e.Adults < 1 {
		size := path.Write("Adults")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "minimum value not met",
			Details: map[string]interface{}{
				"minimum": 1,
			},
		})
		path.TruncateRight(size)
	}

	if e.Adults > 9 {
		size := path.Write("Adults")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "maximum value exceeded",
			Details: map[string]interface{}{
				"maximum": 9,
			},
		})
		path.TruncateRight(size)
	}

	if len(e.Chan) > 12 {
		size := path.Write("Chan")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "maximum length exceeded",
			Details: map[string]interface{}{
				"maximum": 12,
			},
		})
		path.TruncateRight(size)
	}

	if e.Children < 0 {
		size := path.Write("Children")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "minimum value not met",
			Details: map[string]interface{}{
				"minimum": 0,
			},
		})
		path.TruncateRight(size)
	}

	if e.Children > int(math.Max(float64(8-(e.Adults-1)), 0)) {
		size := path.Write("Children")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "maximum value exceeded",
			Details: map[string]interface{}{
				"maximum": int(math.Max(float64(8-(e.Adults-1)), 0)),
			},
		})
		path.TruncateRight(size)
	}

	if e.Int == 0 {
		size := path.Write("Int")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	if e.Int2 == nil {
		size := path.Write("Int2")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	if e.Int2 == nil {
		size := path.Write("Int2")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "value must not be nil",
		})
		path.TruncateRight(size)
	}

	if e.Int2 != nil && *e.Int2 < 0 {
		size := path.Write("Int2")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "minimum value not met",
			Details: map[string]interface{}{
				"minimum": 0,
			},
		})
		path.TruncateRight(size)
	}

	if len(e.Ints) == 0 {
		size := path.Write("Ints")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	if len(e.Ints) > 3 {
		size := path.Write("Ints")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "maximum length exceeded",
			Details: map[string]interface{}{
				"maximum": 3,
			},
		})
		path.TruncateRight(size)
	}

	for i, element := range e.Ints {

		if element == 0 {
			size := path.Write("Ints.[" + strconv.Itoa(i) + "]")
			violations = append(violations, valley.ConstraintViolation{
				Field:   path.String(),
				Message: "a value is required",
			})
			path.TruncateRight(size)
		}

		if element < 0 {
			size := path.Write("Ints.[" + strconv.Itoa(i) + "]")
			violations = append(violations, valley.ConstraintViolation{
				Field:   path.String(),
				Message: "minimum value not met",
				Details: map[string]interface{}{
					"minimum": 0,
				},
			})
			path.TruncateRight(size)
		}

	}

	if e.Nested == nil {
		size := path.Write("Nested")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	if e.Nested != nil {
		size := path.Write("Nested")
		violations = append(violations, e.Nested.Validate(path)...)
		path.TruncateRight(size)
	}

	for i, element := range e.Nesteds {
		if element != nil {
			size := path.Write("Nesteds.[" + strconv.Itoa(i) + "]")
			violations = append(violations, element.Validate(path)...)
			path.TruncateRight(size)
		}

	}

	if len(e.Text) == 0 {
		size := path.Write("Text")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	if len(e.Text) > 12 {
		size := path.Write("Text")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "maximum length exceeded",
			Details: map[string]interface{}{
				"maximum": 12,
			},
		})
		path.TruncateRight(size)
	}

	if len(e.TextMap) == 0 {
		size := path.Write("TextMap")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	for i, element := range e.TextMap {

		if len(element) == 0 {
			size := path.Write("TextMap.[" + fmt.Sprintf("%v", i) + "]")
			violations = append(violations, valley.ConstraintViolation{
				Field:   path.String(),
				Message: "a value is required",
			})
			path.TruncateRight(size)
		}

	}

	path.TruncateRight(1)

	return violations
}

// Validate validates this NestedExample.
// This method was generated by Valley.
func (n NestedExample) Validate(path *valley.Path) []valley.ConstraintViolation {
	var violations []valley.ConstraintViolation

	path.Write(".")

	if len(n.Text) == 0 {
		size := path.Write("Text")
		violations = append(violations, valley.ConstraintViolation{
			Field:   path.String(),
			Message: "a value is required",
		})
		path.TruncateRight(size)
	}

	path.TruncateRight(1)

	return violations
}
