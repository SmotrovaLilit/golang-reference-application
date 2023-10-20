# Logging and panic recovery in handlers

## Overview

- Application logging setting in [main.go](..%2Fcmd%2Fserver%2Fmain.go)
- Adding attributes from context in [log.go](..%2Finternal%2Fpkg%2Flog%2Flog.go)
- Panic recovery is here [panic_recovery_test.go](..%2Ftest%2Fintegration%2Fpanic_recovery_test.go)

## Example of testing that application log information with fields from context [handler_test.go](..%2Finternal%2Fapplication%2Fcommands%2Fapproveprogramversion%2Fhandler_test.go)

💡I did it for one command for example. It can be done for all commands and queries, but not all projects need 100% coverage. It depends on the project.

💡If the project require test logging too, it can be done by providing logger in http and application layers. Using slog.Default doesn't allow to test logging.

## Panic recovery
💡I used middleware for endpoint to recover panic that can happen in endpoint handlers. Because

✅ It allows to log panic error with full context information. 

💡 I also used panic recovery in transport layer to recover panic that can happen in transport layers, because endpoint middleware doesn't catch panic that can happen in transport layer. 
## References
1. https://betterstack.com/community/guides/logging/logging-in-go/#creating-and-using-child-loggers
2. 