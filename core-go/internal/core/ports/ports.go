/* Набор правил сделали = абстракция форма есть не важно кто 
	ядро знает только эти интерфесы - не Redis/OpenAI
	StateStore, AIClient, EventBus, Logger, ConfigProvider*/  

package ports
import (
		"context"
		"core-go/internal/core/model"
)

// ядро только знает дай состояние и сохрани состояние неа не redis, posgres
type StateStore interface {
	Load(ctx context.Context, key model.StateKey) (model.State, error)
	Save(ctx context.Context, key model.StateKey, state model.State) error
}
// ядро не знает, что такое chatgpt только одно генерируй
type AIClient interface {
	Generate(ctx context.Context, prompt string, opts AIOptions) (string, error)
}

type AIOptions struct {
	Temperature float32
	MaxTokens	int 
}
// ядро отправляет событие наружу например пользователь завершил сценарий через шину
type EventBus interface {
	Publish(ctx context.Context, topic string, event any) error
}
// я не зависим и в плане логов от zap
type Logger interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
	Debug(msg string, fields ...any)
}
// а да про env переменные не знаем
type ConfigProvider interface {
	Get(key string) string
}

// вроде чисто идем
