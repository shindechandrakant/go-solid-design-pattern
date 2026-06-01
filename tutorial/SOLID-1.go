package main

//
//import (
//	"database/sql"
//	"fmt"
//	"net/smtp"
//	"sync"
//	"time"
//)
//
//type RateLimiterStrategy interface {
//	Allow(clientId string) bool
//}
//
//type FixedWindowRateLimiter struct {
//	mu       sync.Mutex
//	requests map[string][]time.Time
//	limit    int
//	window   time.Duration
//}
//
//func (fw *FixedWindowRateLimiter) Allow(clientID string) bool {
//	fw.mu.Lock()
//	defer fw.mu.Unlock()
//	now := time.Now()
//
//	// Clean old entries
//	valid := []time.Time{}
//	for _, t := range fw.requests[clientID] {
//		if now.Sub(t) <= fw.window {
//			valid = append(valid, t)
//		}
//	}
//	fw.requests[clientID] = valid
//
//	if len(fw.requests[clientID]) >= fw.limit {
//		return false
//	}
//	fw.requests[clientID] = append(fw.requests[clientID], now)
//	return true
//}
//
//type Db interface {
//	Insert(string, time.Time) error
//}
//
//type RateLimitService struct {
//	db *sql.DB
//}
//
//func (rs *RateLimitService) Insert(clientID string, now time.Time) error {
//	_, err := rs.db.Exec(
//		"INSERT INTO rate_limit_violations (client_id, timestamp) VALUES ($1, $2)",
//		clientID, now,
//	)
//	if err != nil {
//		return fmt.Errorf("log violation: %w", err)
//	}
//	return nil
//}
//
//type RateLimiter struct {
//	strategy     RateLimiterStrategy
//	DB           Db
//	notification Notification1
//}
//
//type Notification1 interface {
//	Send(string) error
//}
//
//type EmailNotification struct {
//	smtpHost string
//}
//
//func (en *EmailNotification) Send(clientID string) error {
//	now := time.Now()
//	body := fmt.Sprintf("Client %s exceeded rate limit at %s", clientID, now)
//	auth := smtp.PlainAuth("", "alerts@api.com", "pass", en.smtpHost)
//	err := smtp.SendMail(en.smtpHost+":587", auth, "alerts@api.com",
//		[]string{"ops@api.com"}, []byte(body))
//	if err != nil {
//		return fmt.Errorf("send alert: %w", err)
//	}
//	return nil
//}
//
//func NewRateLimiter(strategy RateLimiterStrategy, db Db, notification Notification1) *RateLimiter {
//	return &RateLimiter{
//		strategy:     strategy,
//		DB:           db,
//		notification: notification,
//	}
//}
//
//func (r *RateLimiter) Allow(clientID string) (bool, error) {
//	// Check limit
//	if isAllowed := r.strategy.Allow(clientID); isAllowed {
//		return true, nil
//	}
//
//	// Log to database
//	if err := r.DB.Insert(clientID, time.Now()); err != nil {
//		return false, err
//	}
//	// Send alert email
//	if err := r.notification.Send(clientID); err != nil {
//		return false, err
//	}
//	// Record request
//	return false, nil
//}
