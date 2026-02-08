# API Contract: /api/v1/flows

## Methods

- `GET /api/v1/flows`
- `POST /api/v1/flows`
- `POST /api/v1/flows/bind-session`

## Create Flow Request Example

```json
{
  "name": "default",
  "version": "v1",
  "definition": {
    "name": "default",
    "policy": {"parallelism": 1, "failure_mode": "fail_fast"},
    "nodes": [{"id": "plan", "type": "plan"}],
    "edges": []
  }
}
```

## Bind Session Request Example

```json
{
  "sessionId": "ssn-123",
  "templateId": "flow-123",
  "version": "v1"
}
```
