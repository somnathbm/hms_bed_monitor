package metrics

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const appName = "hms-bm-svc"

var meter = otel.Meter(appName)

// emg bed counter
func emgBedCounterMetric() (metric.Int64Counter, error) {
	emgbedCnt, emgBedCnterr := meter.Int64Counter("beds.emg",
		metric.WithDescription("Total emergency beds"),
	)
	if emgBedCnterr != nil {
		return nil, emgBedCnterr
	}
	return emgbedCnt, nil
}

// IPD bed counter
func ipdBedCounterMetric() (metric.Int64Counter, error) {
	ipdbedCnt, ipdBedCnterr := meter.Int64Counter("beds.ipd",
		metric.WithDescription("Total IPD beds"),
	)
	if ipdBedCnterr != nil {
		return nil, ipdBedCnterr
	}
	return ipdbedCnt, nil
}

// OPD bed counter
func opdBedCounterMetric() (metric.Int64Counter, error) {
	opdbedCnt, opdBedCnterr := meter.Int64Counter("beds.opd",
		metric.WithDescription("Total OPD beds"),
	)
	if opdBedCnterr != nil {
		return nil, opdBedCnterr
	}
	return opdbedCnt, nil
}

func GetAllCounterMetrics() map[string]metric.Int64Counter {
	allCounterMetricsMap := make(map[string]metric.Int64Counter)

	emgBedCounterMetric, emgBedCounterMetricErr := emgBedCounterMetric()
	if emgBedCounterMetricErr != nil {
		panic(emgBedCounterMetricErr)
	}
	ipdBedCounterMetric, ipdBedCounterMetricErr := ipdBedCounterMetric()
	if ipdBedCounterMetricErr != nil {
		panic(ipdBedCounterMetricErr)
	}
	opdBedCounterMetric, opdBedCounterMetricErr := opdBedCounterMetric()
	if opdBedCounterMetricErr != nil {
		panic(opdBedCounterMetricErr)
	}
	allCounterMetricsMap["EmgBedCountermetric"] = emgBedCounterMetric
	allCounterMetricsMap["IPDBedCountermetric"] = ipdBedCounterMetric
	allCounterMetricsMap["OPDBedCountermetric"] = opdBedCounterMetric
	return allCounterMetricsMap
}
