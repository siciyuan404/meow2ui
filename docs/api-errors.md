# API Errors

统一错误响应结构：

```json
{
  "code": "NOT_FOUND",
  "message": "resource not found",
  "detail": "not found",
  "traceId": "trace-..."
}
```

## Error Codes

- `METHOD_NOT_ALLOWED`
- `NOT_FOUND`
- `CONFLICT`
- `BAD_REQUEST`
- `INTERNAL_ERROR`
