# Golang Reference Application

## Decision Log

### Errors in repositories
Errors in repositories will be panic in service and logged and don't be handled in the application layer.

Advantages:
- make easy supporting application tests
- errors in repositories are not business errors
- error when id is busy is not business error

TODO: see about gorm save