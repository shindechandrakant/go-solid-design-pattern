package main

import (
	"context"
	"fmt"
)

type LogEntry struct {
	Timestamp int64
	Level     string
	Message   string
}

type AppLogger struct {
	level string
	hooks []func(entry LogEntry)
}

func (l *AppLogger) Info(ctx context.Context, msg string)             { /* ... */ }
func (l *AppLogger) Error(ctx context.Context, msg string, err error) { /* ... */ }
func (l *AppLogger) Debug(ctx context.Context, msg string)            { /* ... */ }
func (l *AppLogger) SetLevel(level string)                            { l.level = level }
func (l *AppLogger) AddHook(fn func(entry LogEntry))                  { l.hooks = append(l.hooks, fn) }
func (l *AppLogger) Flush(ctx context.Context) error                  { /* ... */ return nil }
func (l *AppLogger) ExportLogs(ctx context.Context, from, to int64) ([]LogEntry, error) { /* ... */
	return nil, nil
}

// Every consumer takes the fat interface
type LogInfoLogger interface {
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string, err error)
}

type MockLogInfoLogger struct {
}

func (m *MockLogInfoLogger) Info(ctx context.Context, msg string) {
	fmt.Printf("Info: %s", msg)
}

func (m *MockLogInfoLogger) Error(ctx context.Context, msg string, err error) {
	fmt.Printf("Error: %s", msg)
}

func TestOrderService() {
	logger := &OrderService{
		logger: &MockLogInfoLogger{},
	}
	logger.logger.Info(context.Background(), "Test method")
}

type OrderService struct {
	logger LogInfoLogger // only calls Info and Error
}

type DebugSetLevelLogger interface {
	Debug(ctx context.Context, msg string)
	SetLevel(level string)
}

type DebugTool struct {
	logger DebugSetLevelLogger // only calls Debug and SetLevel
}

type ExportLogFlushLog interface {
	Flush(ctx context.Context) error
	ExportLogs(ctx context.Context, from, to int64) ([]LogEntry, error)
}
type MonitoringJob struct {
	logger ExportLogFlushLog // only calls ExportLogs and Flush
}

type ErrorHookLogger interface {
	Error(ctx context.Context, msg string, err error)
	AddHook(fn func(entry LogEntry))
}
type AlertService struct {
	logger ErrorHookLogger // only calls Error and AddHook
}
