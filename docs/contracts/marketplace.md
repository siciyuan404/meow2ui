# API Contract: /api/v1/marketplace/*

## Methods

- `GET/POST /api/v1/marketplace/templates`
- `POST /api/v1/marketplace/review`
- `POST /api/v1/marketplace/ratings`
- `POST /api/v1/marketplace/apply`

## Create Template Request

```json
{
  "name": "Dashboard Template",
  "category": "dashboard",
  "tags": ["react", "saas"],
  "schema": "{}",
  "theme": "default",
  "owner": "user-1",
  "sessionId": "ssn-1",
  "versionId": "ver-1"
}
```

## Apply Request

```json
{
  "templateId": "tpl-1",
  "sessionId": "ssn-1"
}
```
