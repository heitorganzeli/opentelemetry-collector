// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/processor/batchprocessor")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/processor/batchprocessor")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                              metric.Meter
	mu                                 sync.Mutex
	registrations                      []metric.Registration
	ProcessorBatchBatchSendSize        metric.Int64Histogram
	ProcessorBatchBatchSendSizeBytes   metric.Int64Histogram
	ProcessorBatchBatchSizeTriggerSend metric.Int64Counter
	ProcessorBatchMetadataCardinality  metric.Int64ObservableUpDownCounter
	ProcessorBatchTimeoutTriggerSend   metric.Int64Counter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// RegisterProcessorBatchMetadataCardinalityCallback sets callback for observable ProcessorBatchMetadataCardinality metric.
func (builder *TelemetryBuilder) RegisterProcessorBatchMetadataCardinalityCallback(cb metric.Int64Callback) error {
	reg, err := builder.meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		cb(ctx, &observerInt64{inst: builder.ProcessorBatchMetadataCardinality, obs: o})
		return nil
	}, builder.ProcessorBatchMetadataCardinality)
	if err != nil {
		return err
	}
	builder.mu.Lock()
	defer builder.mu.Unlock()
	builder.registrations = append(builder.registrations, reg)
	return nil
}

type observerInt64 struct {
	embedded.Int64Observer
	inst metric.Int64Observable
	obs  metric.Observer
}

func (oi *observerInt64) Observe(value int64, opts ...metric.ObserveOption) {
	oi.obs.ObserveInt64(oi.inst, value, opts...)
}

// Shutdown unregister all registered callbacks for async instruments.
func (builder *TelemetryBuilder) Shutdown() {
	builder.mu.Lock()
	defer builder.mu.Unlock()
	for _, reg := range builder.registrations {
		reg.Unregister()
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.ProcessorBatchBatchSendSize, err = builder.meter.Int64Histogram(
		"otelcol_processor_batch_batch_send_size",
		metric.WithDescription("Number of units in the batch"),
		metric.WithUnit("{units}"),
		metric.WithExplicitBucketBoundaries([]float64{10, 25, 50, 75, 100, 250, 500, 750, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000, 100000}...),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorBatchBatchSendSizeBytes, err = builder.meter.Int64Histogram(
		"otelcol_processor_batch_batch_send_size_bytes",
		metric.WithDescription("Number of bytes in batch that was sent. Only available on detailed level."),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries([]float64{10, 25, 50, 75, 100, 250, 500, 750, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 50000, 100000, 200000, 300000, 400000, 500000, 600000, 700000, 800000, 900000, 1e+06, 2e+06, 3e+06, 4e+06, 5e+06, 6e+06, 7e+06, 8e+06, 9e+06}...),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorBatchBatchSizeTriggerSend, err = builder.meter.Int64Counter(
		"otelcol_processor_batch_batch_size_trigger_send",
		metric.WithDescription("Number of times the batch was sent due to a size trigger"),
		metric.WithUnit("{times}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorBatchMetadataCardinality, err = builder.meter.Int64ObservableUpDownCounter(
		"otelcol_processor_batch_metadata_cardinality",
		metric.WithDescription("Number of distinct metadata value combinations being processed"),
		metric.WithUnit("{combinations}"),
	)
	errs = errors.Join(errs, err)
	builder.ProcessorBatchTimeoutTriggerSend, err = builder.meter.Int64Counter(
		"otelcol_processor_batch_timeout_trigger_send",
		metric.WithDescription("Number of times the batch was sent due to a timeout trigger"),
		metric.WithUnit("{times}"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
