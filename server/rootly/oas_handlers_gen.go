// Code generated by ogen, DO NOT EDIT.

package rootly

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
)

// HandleScimUsersGetRequest handles  operation.
//
// GET /scim/Users
func (s *Server) handleScimUsersGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ScimUsersGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityBearerAuth(ctx, "ScimUsersGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "ScimUsersGet",
			Security:  "BearerAuth",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeScimUsersGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "ScimUsersGet",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ScimUsersGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeScimUsersGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleScimUsersIDDeleteRequest handles  operation.
//
// DELETE /scim/Users/{id}
func (s *Server) handleScimUsersIDDeleteRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ScimUsersIDDelete",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityBearerAuth(ctx, "ScimUsersIDDelete", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "ScimUsersIDDelete",
			Security:  "BearerAuth",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeScimUsersIDDeleteParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "ScimUsersIDDelete",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ScimUsersIDDelete(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeScimUsersIDDeleteResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleScimUsersIDGetRequest handles  operation.
//
// GET /scim/Users/{id}
func (s *Server) handleScimUsersIDGetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ScimUsersIDGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityBearerAuth(ctx, "ScimUsersIDGet", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "ScimUsersIDGet",
			Security:  "BearerAuth",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeScimUsersIDGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "ScimUsersIDGet",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ScimUsersIDGet(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeScimUsersIDGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleScimUsersIDPatchRequest handles  operation.
//
// PATCH /scim/Users/{id}
func (s *Server) handleScimUsersIDPatchRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ScimUsersIDPatch",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityBearerAuth(ctx, "ScimUsersIDPatch", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "ScimUsersIDPatch",
			Security:  "BearerAuth",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	params, err := decodeScimUsersIDPatchParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			Operation: "ScimUsersIDPatch",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ScimUsersIDPatch(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeScimUsersIDPatchResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleScimUsersPostRequest handles  operation.
//
// POST /scim/Users
func (s *Server) handleScimUsersPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ScimUsersPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	ctx, err = s.securityBearerAuth(ctx, "ScimUsersPost", r)
	if err != nil {
		err = &ogenerrors.SecurityError{
			Operation: "ScimUsersPost",
			Security:  "BearerAuth",
			Err:       err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.ScimUsersPost(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeScimUsersPostResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	span trace.Span,
	otelAttrs []attribute.KeyValue,
	err error,
) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	s.cfg.ErrorHandler(ctx, w, r, err)
}