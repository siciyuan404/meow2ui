# API Contract: /agent/run

## Method

- `POST /agent/run`

## Request Example

```json
{
  "sessionId": "ssn-123",
  "prompt": "生成仪表盘",
  "onlyArea": "main",
  "media": [
    {
      "type": "image",
      "ref": "https://cdn.example.com/ui.png",
      "metadata": {"mime": "image/png"}
    }
  ]
}
```

## Response Example

```json
{
  "runId": "run-123",
  "versionId": "ver-123",
  "schemaJSON": "{...}",
  "repaired": false
}
```
