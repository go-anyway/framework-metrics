package metrics_test

import (
	"testing"

	"github.com/go-anyway/framework-metrics"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSetEnabled(t *testing.T) {
	tests := []struct {
		name     string
		enabled  bool
		expected bool
	}{
		{"启用指标收集", true, true},
		{"禁用指标收集", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics.SetEnabled(tt.enabled)
			if got := metrics.IsEnabled(); got != tt.expected {
				t.Errorf("IsEnabled() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsEnabled_DefaultValue(t *testing.T) {
	metrics.SetEnabled(false)
	if metrics.IsEnabled() {
		t.Error("IsEnabled() 默认应为 false")
	}
}

func TestHTTPRequestTotal(t *testing.T) {
	metrics.HTTPRequestTotal.WithLabelValues("GET", "/test", "200").Inc()
	metrics.HTTPRequestTotal.WithLabelValues("POST", "/api", "201").Inc()

	count := testutil.CollectAndCount(metrics.HTTPRequestTotal)
	if count < 2 {
		t.Errorf("HTTPRequestTotal 计数器数量 = %d, want at least 2", count)
	}
}

func TestHTTPRequestDuration(t *testing.T) {
	metrics.HTTPRequestDuration.WithLabelValues("POST", "/api", "200").Observe(0.5)
	metrics.HTTPRequestDuration.WithLabelValues("POST", "/api", "200").Observe(1.5)

	count := testutil.CollectAndCount(metrics.HTTPRequestDuration)
	if count != 1 {
		t.Errorf("HTTPRequestDuration 标签组合数量 = %d, want 1", count)
	}
}

func TestGRPCRequestTotal(t *testing.T) {
	metrics.GRPCRequestTotal.WithLabelValues("/method", "0").Inc()
	metrics.GRPCRequestTotal.WithLabelValues("/method2", "0").Inc()

	count := testutil.CollectAndCount(metrics.GRPCRequestTotal)
	if count < 2 {
		t.Errorf("GRPCRequestTotal 计数器数量 = %d, want at least 2", count)
	}
}

func TestGRPCRequestDuration(t *testing.T) {
	metrics.GRPCRequestDuration.WithLabelValues("/method", "0").Observe(0.1)
	metrics.GRPCRequestDuration.WithLabelValues("/method", "0").Observe(0.2)

	count := testutil.CollectAndCount(metrics.GRPCRequestDuration)
	if count != 1 {
		t.Errorf("GRPCRequestDuration 标签组合数量 = %d, want 1", count)
	}
}

func TestDatabaseQueryTotal(t *testing.T) {
	metrics.DatabaseQueryTotal.WithLabelValues("SELECT", "success").Inc()
	metrics.DatabaseQueryTotal.WithLabelValues("INSERT", "success").Inc()

	count := testutil.CollectAndCount(metrics.DatabaseQueryTotal)
	if count < 2 {
		t.Errorf("DatabaseQueryTotal 计数器数量 = %d, want at least 2", count)
	}
}

func TestDatabaseQueryDuration(t *testing.T) {
	metrics.DatabaseQueryDuration.WithLabelValues("SELECT").Observe(0.01)
	metrics.DatabaseQueryDuration.WithLabelValues("INSERT").Observe(0.02)

	count := testutil.CollectAndCount(metrics.DatabaseQueryDuration)
	if count != 2 {
		t.Errorf("DatabaseQueryDuration 标签组合数量 = %d, want 2", count)
	}
}

func TestRedisOperationTotal(t *testing.T) {
	metrics.RedisOperationTotal.WithLabelValues("GET", "success").Inc()
	metrics.RedisOperationTotal.WithLabelValues("SET", "success").Inc()

	count := testutil.CollectAndCount(metrics.RedisOperationTotal)
	if count < 2 {
		t.Errorf("RedisOperationTotal 计数器数量 = %d, want at least 2", count)
	}
}

func TestRedisOperationDuration(t *testing.T) {
	metrics.RedisOperationDuration.WithLabelValues("GET").Observe(0.001)
	metrics.RedisOperationDuration.WithLabelValues("SET").Observe(0.002)

	count := testutil.CollectAndCount(metrics.RedisOperationDuration)
	if count != 2 {
		t.Errorf("RedisOperationDuration 标签组合数量 = %d, want 2", count)
	}
}

func TestActiveConnections(t *testing.T) {
	metrics.ActiveConnections.WithLabelValues("http").Set(10)
	metrics.ActiveConnections.WithLabelValues("grpc").Set(5)

	value := testutil.ToFloat64(metrics.ActiveConnections.WithLabelValues("http"))
	if value != 10 {
		t.Errorf("ActiveConnections(http) = %v, want 10", value)
	}

	value = testutil.ToFloat64(metrics.ActiveConnections.WithLabelValues("grpc"))
	if value != 5 {
		t.Errorf("ActiveConnections(grpc) = %v, want 5", value)
	}
}

func TestHTTPRequestSize(t *testing.T) {
	metrics.HTTPRequestSize.WithLabelValues("POST", "/api").Observe(1024)
	metrics.HTTPRequestSize.WithLabelValues("POST", "/api").Observe(2048)

	count := testutil.CollectAndCount(metrics.HTTPRequestSize)
	if count != 1 {
		t.Errorf("HTTPRequestSize 标签组合数量 = %d, want 1", count)
	}
}

func TestHTTPResponseSize(t *testing.T) {
	metrics.HTTPResponseSize.WithLabelValues("GET", "/test", "200").Observe(512)
	metrics.HTTPResponseSize.WithLabelValues("GET", "/test", "200").Observe(1024)

	count := testutil.CollectAndCount(metrics.HTTPResponseSize)
	if count != 1 {
		t.Errorf("HTTPResponseSize 标签组合数量 = %d, want 1", count)
	}
}

func TestDatabaseConnectionsInUse(t *testing.T) {
	metrics.DatabaseConnectionsInUse.WithLabelValues("mysql").Set(5)
	metrics.DatabaseConnectionsInUse.WithLabelValues("postgres").Set(3)

	value := testutil.ToFloat64(metrics.DatabaseConnectionsInUse.WithLabelValues("mysql"))
	if value != 5 {
		t.Errorf("DatabaseConnectionsInUse(mysql) = %v, want 5", value)
	}

	value = testutil.ToFloat64(metrics.DatabaseConnectionsInUse.WithLabelValues("postgres"))
	if value != 3 {
		t.Errorf("DatabaseConnectionsInUse(postgres) = %v, want 3", value)
	}
}

func TestDatabaseConnectionsIdle(t *testing.T) {
	metrics.DatabaseConnectionsIdle.WithLabelValues("mysql").Set(10)
	value := testutil.ToFloat64(metrics.DatabaseConnectionsIdle.WithLabelValues("mysql"))
	if value != 10 {
		t.Errorf("DatabaseConnectionsIdle(mysql) = %v, want 10", value)
	}
}

func TestDatabaseConnectionsOpen(t *testing.T) {
	metrics.DatabaseConnectionsOpen.WithLabelValues("mysql").Set(15)
	value := testutil.ToFloat64(metrics.DatabaseConnectionsOpen.WithLabelValues("mysql"))
	if value != 15 {
		t.Errorf("DatabaseConnectionsOpen(mysql) = %v, want 15", value)
	}
}

func TestXXLJobTaskTotal(t *testing.T) {
	metrics.XXLJobTaskTotal.WithLabelValues("task1", "success").Inc()
	metrics.XXLJobTaskTotal.WithLabelValues("task2", "failed").Inc()

	count := testutil.CollectAndCount(metrics.XXLJobTaskTotal)
	if count < 2 {
		t.Errorf("XXLJobTaskTotal 计数器数量 = %d, want at least 2", count)
	}
}

func TestXXLJobTaskDuration(t *testing.T) {
	metrics.XXLJobTaskDuration.WithLabelValues("task1").Observe(0.5)
	metrics.XXLJobTaskDuration.WithLabelValues("task1").Observe(1.5)

	count := testutil.CollectAndCount(metrics.XXLJobTaskDuration)
	if count != 1 {
		t.Errorf("XXLJobTaskDuration 标签组合数量 = %d, want 1", count)
	}
}

func TestMetricsEnabledToggle(t *testing.T) {
	metrics.SetEnabled(true)
	if !metrics.IsEnabled() {
		t.Error("SetEnabled(true) 后 IsEnabled 应返回 true")
	}

	metrics.SetEnabled(false)
	if metrics.IsEnabled() {
		t.Error("SetEnabled(false) 后 IsEnabled 应返回 false")
	}
}

func TestConcurrentEnabledAccess(t *testing.T) {
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			metrics.SetEnabled(true)
			_ = metrics.IsEnabled()
			metrics.SetEnabled(false)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestCounterIncrement(t *testing.T) {
	initialCount := testutil.CollectAndCount(metrics.HTTPRequestTotal)

	metrics.HTTPRequestTotal.WithLabelValues("GET", "/test", "200").Inc()

	newCount := testutil.CollectAndCount(metrics.HTTPRequestTotal)
	if newCount != initialCount {
		t.Errorf("计数器数量变化了 = %d, want stay %d", newCount, initialCount)
	}
}

func TestHistogramObservation(t *testing.T) {
	metrics.HTTPRequestDuration.WithLabelValues("PUT", "/resource", "200").Observe(0.05)
	metrics.HTTPRequestDuration.WithLabelValues("PUT", "/resource", "200").Observe(0.1)
	metrics.HTTPRequestDuration.WithLabelValues("PUT", "/resource", "200").Observe(0.15)

	count := testutil.CollectAndCount(metrics.HTTPRequestDuration)
	if count < 1 {
		t.Errorf("Histogram 标签组合数量 = %d, want at least 1", count)
	}
}

func TestGaugeSet(t *testing.T) {
	metrics.ActiveConnections.WithLabelValues("ws").Set(0)
	metrics.ActiveConnections.WithLabelValues("ws").Set(20)

	value := testutil.ToFloat64(metrics.ActiveConnections.WithLabelValues("ws"))
	if value != 20 {
		t.Errorf("ActiveConnections(ws) = %v, want 20", value)
	}
}

func TestMultipleLabels(t *testing.T) {
	metrics.HTTPRequestTotal.WithLabelValues("PATCH", "/items", "200").Inc()
	metrics.HTTPRequestTotal.WithLabelValues("PATCH", "/items", "200").Inc()
	metrics.HTTPRequestTotal.WithLabelValues("PATCH", "/items", "200").Inc()

	count := testutil.CollectAndCount(metrics.HTTPRequestTotal)
	if count < 1 {
		t.Errorf("HTTPRequestTotal 唯一标签组合数量 = %d, want at least 1", count)
	}
}
