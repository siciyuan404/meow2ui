# Multimodal Support Guide

## Overview

Multimodal support allows image/audio references to be attached to agent tasks.

## Request Shape

Use `/agent/run` with `media`:

```json
{
  "sessionId": "ssn-xxx",
  "prompt": "Generate UI from uploaded image",
  "media": [
    {
      "type": "image",
      "ref": "https://cdn.example.com/mock.png",
      "metadata": {"mime": "image/png"}
    }
  ]
}
```

## Validation Rules

- Local/internal URL prefixes are blocked.
- External URLs require allow-list hosts.
- `s3://`, `oss://`, `local://` refs are accepted.

## Storage

- `schema_version_assets` stores version-to-media references.
- Local storage backend is available at `pkg/media/storage`.
