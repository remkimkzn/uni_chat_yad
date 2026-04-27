/* Набор правил сделали = абстракция форма есть не важно кто 
	ядро знает только эти интерфесы - не Redis/OpenAI
	StateStore, AIClient, EventBus, Logger, ConfigProvider*/  

package ports
import "context"

type StateStore interface {
	Load(ctx context.Context, key StateKey) (State, error)
	Save(ctx context.Context, key StateKey, state State) error
}

type AIClient interface {
	Generate(ctx context.Context, prompt string, opts AIOptions) (string, error)
}

type EventBus interface {
	Publish(ctx context.Context, topic string, event any) error
}
