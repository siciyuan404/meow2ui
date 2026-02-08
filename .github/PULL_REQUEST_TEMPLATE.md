## Summary

- What problem does this PR solve?
- What changed and why?

## Risks

- Potential impact areas
- Rollback strategy

## Verification

- [ ] `go test ./... -coverprofile=coverage.out`
- [ ] `bash scripts/check-coverage.sh 40`
- [ ] `npm --prefix web run test -- --run` (if web changed)
- [ ] `bash scripts/integration-postgres.sh` (if storage/API changed)
- [ ] `bash scripts/smoke.sh` (if runtime path changed)

## Related Issue

- Fixes #
