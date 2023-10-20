# Logging and panic recovery in handlers

## Overview

- Application logging setting in [main.go](..%2Fcmd%2Fserver%2Fmain.go)
- Adding attributes from the context in [log.go](..%2Finternal%2Fpkg%2Flog%2Flog.go)
- Panic recovery is here [panic_recovery_test.go](..%2Ftest%2Fintegration%2Fpanic_recovery_test.go)

## Panic recovery
💡I used middleware for an endpoint to recover panic that can happen in endpoint handlers. Because

✅ It allows us to log panic errors with full context information. 

💡 I also used panic recovery in the transport layer to recover panic that can happen in transport layers, because endpoint middleware doesn't catch panic that can happen in the transport layer. 

## References
1. https://betterstack.com/community/guides/logging/logging-in-go/#creating-and-using-child-loggers
2. 
