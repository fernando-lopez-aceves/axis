package models

import (
	"errors"
	"fmt"
	"time"
)

type AuditStatus string
type AuditService string
type Severity int

const (
	StatusSuccess AuditStatus = "success"
	StatusFailure AuditStatus = "failure"
	StatusDenied  AuditStatus = "access_denied"

	ServiceIdentity AuditService = "identity"
	ServiceHR       AuditService = "hr"
	ServiceAudit    AuditService = "audit"
	ServiceFinance  AuditService = "finance"

	SeverityLow      Severity = 1
	SeverityMedium   Severity = 2
	SeverityHigh     Severity = 3
	SeverityCritical Severity = 4
)

type AuditEntry struct {
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	Service   AuditService           `json:"service"`
	UserID    string                 `json:"user_id"`
	Action    string                 `json:"action"`
	Resource  string                 `json:"resource"`
	OldValue  map[string]interface{} `json:"old_value"`
	NewValue  map[string]interface{} `json:"new_value"`
	Status    AuditStatus            `json:"status"`
	Severity  Severity               `json:"severity"`
	IPAddress string                 `json:"ip_address"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// Validate realiza una inspección profunda (Versión Completa Original).
func (e *AuditEntry) Validate() error {
	if e.ID == "" {
		return errors.New("audit_validation: missing unique event id (uuid)")
	}
	if e.Timestamp.IsZero() {
		return errors.New("audit_validation: timestamp cannot be empty")
	}
	if e.Service == "" {
		return errors.New("audit_validation: service origin must be defined")
	}
	if e.UserID == "" {
		return errors.New("audit_validation: user_id is required for accountability")
	}
	if len(e.Action) < 5 {
		return fmt.Errorf("audit_validation: action name '%s' is too short or missing", e.Action)
	}
	if e.Status == "" {
		return errors.New("audit_validation: operation status is required")
	}

	// Verificamos que si hay un cambio, se haya registrado al menos un mapa vacío, no un nulo
	if e.OldValue == nil || e.NewValue == nil {
		return errors.New("audit_validation: old_value and new_value must be initialized maps (not nil)")
	}

	return nil
}
