# API Contract: /api/v1/debug/runs*

## Methods

- `GET /api/v1/debug/runs`
- `GET /api/v1/debug/runs/{id}`
- `GET /api/v1/debug/runs/{id}/cost`

## List Response Example

```json
{
  "runs": [
    {
      "runId": "run-1",
      "sessionId": "ssn-1",
      "status": "completed",
      "startedAt": "2026-02-08T12:00:00Z",
      "durationMs": 1234
    }
  ]
}
```

## Cost Response Example

```json
{
  "totalTokens": 240,
  "totalCost": 0.00048,
  "byModel": {"mock-text": 0.00048},
  "byProvider": {"mock-provider": 0.00048}
}
```
