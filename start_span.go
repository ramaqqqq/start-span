package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context, tracer trace.Tracer, spanName string) (context.Context, trace.Span) {
	if tracer == nil {
		return ctx, trace.SpanFromContext(ctx)
	}
	return tracer.Start(ctx, spanName)
}
