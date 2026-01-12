// Copyright 2025 zampo.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// @contact  zampo3380@gmail.com

package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// enabled 指标收集是否启用
	enabled bool
	mu      sync.RWMutex
)

// SetEnabled 设置指标收集是否启用
// 应该在应用启动时根据配置调用此函数
func SetEnabled(e bool) {
	mu.Lock()
	defer mu.Unlock()
	enabled = e
}

// IsEnabled 返回指标收集是否启用
func IsEnabled() bool {
	mu.RLock()
	defer mu.RUnlock()
	return enabled
}

var (
	// HTTPRequestTotal HTTP 请求总数
	HTTPRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// HTTPRequestDuration HTTP 请求耗时
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	// GRPCRequestTotal gRPC 请求总数
	GRPCRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "code"},
	)

	// GRPCRequestDuration gRPC 请求耗时
	GRPCRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "gRPC request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "code"},
	)

	// DatabaseQueryTotal 数据库查询总数
	DatabaseQueryTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"operation", "status"},
	)

	// DatabaseQueryDuration 数据库查询耗时
	DatabaseQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "database_query_duration_seconds",
			Help:    "Database query duration in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"operation"},
	)

	// RedisOperationTotal Redis 操作总数
	RedisOperationTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "redis_operations_total",
			Help: "Total number of Redis operations",
		},
		[]string{"operation", "status"},
	)

	// RedisOperationDuration Redis 操作耗时
	RedisOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "redis_operation_duration_seconds",
			Help:    "Redis operation duration in seconds",
			Buckets: []float64{.0001, .0005, .001, .005, .01, .025, .05, .1, .25, .5, 1},
		},
		[]string{"operation"},
	)

	// ActiveConnections 活跃连接数
	ActiveConnections = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
		[]string{"type"},
	)

	// HTTPRequestSize HTTP 请求大小（字节）
	HTTPRequestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "HTTP request size in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7), // 100B to 100MB
		},
		[]string{"method", "path"},
	)

	// HTTPResponseSize HTTP 响应大小（字节）
	HTTPResponseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "HTTP response size in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7), // 100B to 100MB
		},
		[]string{"method", "path", "status"},
	)

	// DatabaseConnectionsInUse 数据库连接池使用中的连接数
	DatabaseConnectionsInUse = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "database_connections_in_use",
			Help: "Number of database connections currently in use",
		},
		[]string{"database"},
	)

	// DatabaseConnectionsIdle 数据库连接池空闲连接数
	DatabaseConnectionsIdle = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "database_connections_idle",
			Help: "Number of idle database connections",
		},
		[]string{"database"},
	)

	// DatabaseConnectionsOpen 数据库连接池总连接数
	DatabaseConnectionsOpen = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "database_connections_open",
			Help: "Number of open database connections",
		},
		[]string{"database"},
	)

	// XXLJobTaskTotal XXL-JOB 任务执行总数
	XXLJobTaskTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "xxljob_tasks_total",
			Help: "Total number of XXL-JOB task executions",
		},
		[]string{"task_name", "status"},
	)

	// XXLJobTaskDuration XXL-JOB 任务执行耗时
	XXLJobTaskDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "xxljob_task_duration_seconds",
			Help:    "XXL-JOB task execution duration in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30, 60},
		},
		[]string{"task_name"},
	)

	// OSSOperationTotal OSS 操作总数
	OSSOperationTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "oss_operations_total",
			Help: "Total number of OSS operations",
		},
		[]string{"provider", "operation", "status"},
	)

	// OSSOperationDuration OSS 操作耗时
	OSSOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "oss_operation_duration_seconds",
			Help:    "OSS operation duration in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30, 60},
		},
		[]string{"provider", "operation"},
	)
)
