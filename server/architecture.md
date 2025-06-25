## PR Review Processing Flow (Redis + Asynq + Worker)

```mermaid
flowchart LR
    ghwebhook[GitHub Webhook PR Open or Update] --> httpapi[HTTP Server API]
    httpapi --> redis[Redis Asynq DB]
    redis --> worker[Asynq Worker]
    worker --> ai[AI Review Engine LLM]
    ai --> ghapi[GitHub API Post Comments]

    httpapi --> validate[Validate Webhook]
    validate --> enqueue[Enqueue Task Asynq]
    enqueue --> redis

    worker --> dequeue[Dequeues Task]
    dequeue --> fetch[Fetch PR Diff]
    fetch --> buildprompt[Build Prompt]
    buildprompt --> callai[Call AI API]
    callai --> postcomment[Post Comments]
    postcomment --> ghapi

```

+----------------------+ +------------------------+ +------------------+
| GitHub Webhook | -----> | HTTP Server (API) | | AI Review |
| (PR Open/Update) | | - Validates Webhook | | Generator (LLM)|
+----------------------+ | - Enqueues Task (Asynq)| +------------------+
+------------------------+
|
v
+------------------+
| Redis (Asynq DB)|
| - Stores Queued |
| Tasks |
+------------------+
|
v
+-----------------------------+
| Asynq Worker (Goroutine) |
| - Dequeues Task |
| - Fetches PR Diff |
| - Builds Prompt |
| - Calls AI API |
| - Posts Comments via GitHub |
+-----------------------------+

## Understand graceful shutdown of services

```mermaid
flowchart TD
    main[main function] --> init[Initialize Dependencies]
    init --> sig[SIGINT Listener Goroutine]
    init --> runworker[Run Worker Goroutine]
    init --> startserver[Start Server Goroutine]
    sig --> wait[Wait for Shutdown]
    runworker --> wait
    startserver --> wait


```

┌──────────────┐
│ main() │
└────┬─────────┘
│
├─ initializeDependencies()
├─ signal listener goroutine ───┐
├─ RunWorker() goroutine │ (listens for ctx cancel)
├─ startServer() goroutine │ (listens for ctx cancel)
└─ wg.Wait() <──────────────────┘ waits for both
