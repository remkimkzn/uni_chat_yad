/* Здесь валидация входного Message, определение bot_id, tenant_id
вызов logicEngine, сборка ответа.
Никаких HTTP - только функция */

func (r *Router) HandleMessage(ctx context.Context, msg model.Message) (model.ResponseEnvelope, error)


