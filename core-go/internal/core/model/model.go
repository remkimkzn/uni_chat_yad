package model 

import "time"

// Канал общения (абстрактный (не телеграм, почта)
type Channel string

// Идентификаторы
type TenantID string
type BotID string
type UserID string
type SessionID string 

// Входящее сообщение в ядро (универсальное)
type Message struct {
	ID				string			'json:"id"'
	TenantID		TenantID		'json:"tenant_id"'
	BotID			BotID			'json:"bot_id"'
	UserID			UserID			'json:"user_id"'
	Channel			Channel			'json:"channel"'
	SessionID		SessionID		'json:"session_id,omitempty"'
	Type			string			'json:"type"' // "text", "event", "command" ...
	Content			string			'json:"content"' // text or payload
	Metadata		map[string]any	'json:"metadata,omitempty"'
	ReseivedAt		time.Time		'json:"reseived_at"'
	CorrelationID	string			'json:"correlation_id,omitempty"'
}

