// Code generated by ogen, DO NOT EDIT.

package rootly

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

type ScimUsersGetParams struct {
	Filter     string
	StartIndex OptInt
	Count      OptInt
}

func decodeScimUsersGetParams(args [0]string, r *http.Request) (params ScimUsersGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: filter.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "filter",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Filter = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: filter: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: startIndex.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "startIndex",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartIndexVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotStartIndexVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.StartIndex.SetTo(paramsDotStartIndexVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: startIndex: parse")
			}
		}
	}
	// Decode query: count.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "count",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCountVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotCountVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Count.SetTo(paramsDotCountVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: count: parse")
			}
		}
	}
	return params, nil
}

type ScimUsersIDDeleteParams struct {
	ID int
}

func decodeScimUsersIDDeleteParams(args [1]string, r *http.Request) (params ScimUsersIDDeleteParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

type ScimUsersIDGetParams struct {
	ID int
}

func decodeScimUsersIDGetParams(args [1]string, r *http.Request) (params ScimUsersIDGetParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

type ScimUsersIDPatchParams struct {
	ID int
}

func decodeScimUsersIDPatchParams(args [1]string, r *http.Request) (params ScimUsersIDPatchParams, _ error) {
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}