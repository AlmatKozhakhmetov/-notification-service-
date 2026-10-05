# Notification Service (Go Worker Pool)

Микросервис асинхронной рассылки уведомлений (Email/Telegram), написанный на Go с использованием паттерна Worker Pool и Clean Architecture.

## Особенности
- **Асинхронная обработка:** Сервер мгновенно возвращает ответ `202 Accepted`, а задачи отправляются в буферизированный канал.
- **Worker Pool:** Ограничение и контроль количества одновременно выполняющихся фоновых горутин.
- **Clean Architecture:** Четкое разделение на `domain`, `worker`, `handler` и `cmd`.

## Архитектура
```text
  [HTTP Handler]  --->  [ Buffered Channel ]  --->  [ Worker 1 ]
  (Status 202)            (Task Queue)               [ Worker 2 ]
                                                     [ Worker 3 ]