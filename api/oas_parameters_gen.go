// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

type CreateFunctionParams struct {
	// Project ref.
	Ref string
}

func unpackCreateFunctionParams(packed map[string]any) (params CreateFunctionParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeCreateFunctionParams(args [1]string, r *http.Request) (params CreateFunctionParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type CreateSecretsParams struct {
	// Project ref.
	Ref string
}

func unpackCreateSecretsParams(packed map[string]any) (params CreateSecretsParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeCreateSecretsParams(args [1]string, r *http.Request) (params CreateSecretsParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type DeleteFunctionParams struct {
	// Project ref.
	Ref string
	// Function slug.
	FunctionSlug string
}

func unpackDeleteFunctionParams(packed map[string]any) (params DeleteFunctionParams) {
	params.Ref = packed["ref"].(string)
	params.FunctionSlug = packed["function_slug"].(string)
	return params
}

func decodeDeleteFunctionParams(args [2]string, r *http.Request) (params DeleteFunctionParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	// Decode path: function_slug.
	{
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "function_slug",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FunctionSlug = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["/^[A-Za-z0-9_-]+$/"],
				}).Validate(string(params.FunctionSlug)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: invalid")
			}
		} else {
			return params, errors.New("path: function_slug: not specified")
		}
	}
	return params, nil
}

type DeleteSecretsParams struct {
	// Project ref.
	Ref string
}

func unpackDeleteSecretsParams(packed map[string]any) (params DeleteSecretsParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeDeleteSecretsParams(args [1]string, r *http.Request) (params DeleteSecretsParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type GetConfigParams struct {
	// Project ref.
	Ref string
}

func unpackGetConfigParams(packed map[string]any) (params GetConfigParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeGetConfigParams(args [1]string, r *http.Request) (params GetConfigParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type GetFunctionParams struct {
	IncludeBody OptBool
	// Project ref.
	Ref string
	// Function slug.
	FunctionSlug string
}

func unpackGetFunctionParams(packed map[string]any) (params GetFunctionParams) {
	if v, ok := packed["include_body"]; ok {
		params.IncludeBody = v.(OptBool)
	}
	params.Ref = packed["ref"].(string)
	params.FunctionSlug = packed["function_slug"].(string)
	return params
}

func decodeGetFunctionParams(args [2]string, r *http.Request) (params GetFunctionParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: include_body.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "include_body",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIncludeBodyVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotIncludeBodyVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IncludeBody.SetTo(paramsDotIncludeBodyVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: include_body: parse")
			}
		}
	}
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	// Decode path: function_slug.
	{
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "function_slug",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FunctionSlug = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["/^[A-Za-z0-9_-]+$/"],
				}).Validate(string(params.FunctionSlug)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: invalid")
			}
		} else {
			return params, errors.New("path: function_slug: not specified")
		}
	}
	return params, nil
}

type GetFunctionsParams struct {
	// Project ref.
	Ref string
}

func unpackGetFunctionsParams(packed map[string]any) (params GetFunctionsParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeGetFunctionsParams(args [1]string, r *http.Request) (params GetFunctionsParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type GetSecretsParams struct {
	// Project ref.
	Ref string
}

func unpackGetSecretsParams(packed map[string]any) (params GetSecretsParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeGetSecretsParams(args [1]string, r *http.Request) (params GetSecretsParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type UpdateConfigParams struct {
	// Project ref.
	Ref string
}

func unpackUpdateConfigParams(packed map[string]any) (params UpdateConfigParams) {
	params.Ref = packed["ref"].(string)
	return params
}

func decodeUpdateConfigParams(args [1]string, r *http.Request) (params UpdateConfigParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	return params, nil
}

type UpdateFunctionParams struct {
	// Project ref.
	Ref string
	// Function slug.
	FunctionSlug string
}

func unpackUpdateFunctionParams(packed map[string]any) (params UpdateFunctionParams) {
	params.Ref = packed["ref"].(string)
	params.FunctionSlug = packed["function_slug"].(string)
	return params
}

func decodeUpdateFunctionParams(args [2]string, r *http.Request) (params UpdateFunctionParams, _ error) {
	// Decode path: ref.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "ref",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ref = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    20,
					MinLengthSet: true,
					MaxLength:    20,
					MaxLengthSet: true,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Ref)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: ref: invalid")
			}
		} else {
			return params, errors.New("path: ref: not specified")
		}
	}
	// Decode path: function_slug.
	{
		param := args[1]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "function_slug",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FunctionSlug = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: parse")
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["/^[A-Za-z0-9_-]+$/"],
				}).Validate(string(params.FunctionSlug)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: function_slug: invalid")
			}
		} else {
			return params, errors.New("path: function_slug: not specified")
		}
	}
	return params, nil
}
