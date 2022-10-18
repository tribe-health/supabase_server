// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s CreateFunctionBody) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["/^[A-Za-z0-9_-]+$/"],
		}).Validate(string(s.Slug)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "slug",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateProjectBody) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Plan.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "plan",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Region.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "region",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateProjectBodyPlan) Validate() error {
	switch s {
	case "free":
		return nil
	case "pro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s CreateProjectBodyRegion) Validate() error {
	switch s {
	case "us-east-1":
		return nil
	case "us-west-1":
		return nil
	case "ap-southeast-1":
		return nil
	case "ap-northeast-1":
		return nil
	case "ap-northeast-2":
		return nil
	case "ap-southeast-2":
		return nil
	case "eu-west-1":
		return nil
	case "eu-west-2":
		return nil
	case "eu-central-1":
		return nil
	case "ca-central-1":
		return nil
	case "ap-south-1":
		return nil
	case "sa-east-1":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s CreateSecretBody) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["/^(?!SUPABASE_).*/"],
		}).Validate(string(s.Value)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "value",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FunctionResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Version)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "version",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.CreatedAt)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "created_at",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.UpdatedAt)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "updated_at",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FunctionResponseStatus) Validate() error {
	switch s {
	case "ACTIVE":
		return nil
	case "REMOVED":
		return nil
	case "THROTTLED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s FunctionSlugResponse) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.Version)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "version",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.CreatedAt)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "created_at",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.UpdatedAt)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "updated_at",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FunctionSlugResponseStatus) Validate() error {
	switch s {
	case "ACTIVE":
		return nil
	case "REMOVED":
		return nil
	case "THROTTLED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s GetFunctionsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s GetOrganizationsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s GetSecretsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}