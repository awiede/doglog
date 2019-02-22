package doglog

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"log"
	"strconv"
)

// buildSuffixFromContext creates a suffix to a log string that contains trace_id and span_id in a format that Datadog can consume
func buildSuffixFromContext(ctx ddtrace.SpanContext) string {
	return " dd.trace_id=" + strconv.FormatUint(ctx.TraceID(), 10) + " dd.span_id=" + strconv.FormatUint(ctx.SpanID(), 10)
}

// updateFormat updates a format string by appending the trace_id and span_id in a format that Datadog can consume
func updateFormat(ctx ddtrace.SpanContext, format string) string {
	return format + buildSuffixFromContext(ctx)
}

// Printf is a wrapper around log.Printf that appends trace_id and span_id to the format so that Datadog can correlate logs to traces
func Printf(ctx ddtrace.SpanContext, format string, v ...interface{}) {
	log.Printf(updateFormat(ctx, format), v)
}

// Println is a wrapper around log.Println that appends trace_id and span_id to the log statement so that Datadog can correlate logs to traces
func Println(ctx ddtrace.SpanContext, v ...interface{}) {
	log.Println(v, buildSuffixFromContext(ctx))
}

// Fatal is a wrapper around log.Fatal that appends trace_id and span_id to the fatal log statement so that Datadog can correlate logs to traces
func Fatal(ctx ddtrace.SpanContext, v ...interface{}) {
	log.Fatal(v, buildSuffixFromContext(ctx))
}

// Fatalf a wrapper around log.Fatalf that appends trace_id and span_id to the format so that Datadog can correlate logs to traces
func Fatalf(ctx ddtrace.SpanContext, format string, v ...interface{}) {
	log.Fatalf(updateFormat(ctx, format), v)
}

// Fatalln is a wrapper around log.Fatalln that appends trace_id and span_id to the log statement so that Datadog can correlate logs to traces
func Fatalln(ctx ddtrace.SpanContext, v ...interface{})  {
	log.Fatalln(v, buildSuffixFromContext(ctx))
}

// Panic is a wrapper around log.Panic that appends trace_id and span_id to the panic text so that Datadog can correlate logs to traces
func Panic(ctx ddtrace.SpanContext, v ...interface{}) {
	log.Panic(v, buildSuffixFromContext(ctx))
}

// Panicf a wrapper around log.Panicf that appends trace_id and span_id to the format so that Datadog can correlate logs to traces
func Panicf(ctx ddtrace.SpanContext, format string, v ...interface{}) {
	log.Panicf(updateFormat(ctx, format), v)
}

// Panicln is a wrapper around log.Panicln that appends trace_id and span_id to the log statement so that Datadog can correlate logs to traces
func Panicln(ctx ddtrace.SpanContext, v ...interface{}) {
	log.Panicln(v, buildSuffixFromContext(ctx))
}
