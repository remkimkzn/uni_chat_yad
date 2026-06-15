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

// Состояние диалога 
type State struct {
Step			string			'json:"step"'		// текущий шаг сценария
Data		map[string]any		'json:"data"'		// временные данные
Flags		map[string]bool		'json:"flags"'		// флаги (ожидаем телефон ...)
Context		map[string]any		'json:"context"'	// контекст для AI/логики
UpdatedAT 	time.Time			'json:"update_at"'	// для TTL/отладки
}

// Тип ответа (абстрактный, без специфики)
type ResponseType string

const {
	ResponseText 	ResponseType = "text"
	ResponseAction	ResponseType = "action"
	ResponseForm	ResponseType = "form"
	ResponseImage	ResponseType = "image"
}

// Действие (кнопка, выбор, интент)
type Action struct {
	ID			string			'json:"id"'
	Label 		string			'json:"label"'
	Payload		map[string]any	'json:"payload,omitempty"'
}

// Универсальный ответ ядра 
type Response struct {
	Type		ResponseType	'json:"type"'
	Content		string			'json:"content,omitempty"'
	Actions		[]Action		'json:"actions,omitempty"'
	Meta		map[string]any	'json:"meta,omitempty"'
}

// Обертка ответа для адаптера
type ResponseEnvelope struct {
	TenantID	TenantID		'json:"tenant_id"'
	BotID		BotID			'json:"bot_id"'
	UserID		UserID			'json:"user_id"'
	Channel		Channel			'json:"channel"'
	Items		[]Responce		'json:"items"'
	}

/* здесь ни слова про телеграм/почту
это чистые структуры, которые будут жить годами
все ядро дальше будет возится вокруг этих типов */

