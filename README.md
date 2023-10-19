# Golang Reference Application

## Overview
This repository serves as a reference application that demonstrates the Domain-Driven Design (DDD) approach using Go. 
The code aims to offer a hands-on example of implementing Domain-Driven Design (DDD) in Go, while also showcasing the methodologies I've developed in my work.

## Status
⚠️ This repository is a work in progress. 

While the codebase is functional, there are known issues and improvements that are actively being addressed. 
For more information, see the [Issues](https://github.com/SmotrovaLilit/golang-reference-application/issues) section.

Every controversial decision, along with its pros and cons, is documented in the  [Decision Log.md](docs/Decision%20Log.md)

## Features
- Domain-Driven Design: The architecture follows DDD principles, separating the domain logic, application logic, and interfaces.
- Test Coverage: Unit tests and integration tests to ensure code quality.
- Example of query [Get approve programs](internal%2Fapplication%2Fqueries%2Fapprovedprograms%2Fhandler.go)
- Example of command [Create program](internal%2Fapplication%2Fcommands%2Fcreateprogram%2Fhandler.go)

## Why This Repository?

It is intended both as a portfolio piece for potential employers and as a personal knowledge base of practical coding solutions that can be re-used in future projects.

## Plans
- [ ] Add version handler
- [ ] Add example of using telegram bot
- [ ] Add example of using gRPC
- [ ] Add example of using bus
- [ ] Add example of pessimistic and optimistic locking of the aggregate
- [ ] Add example of using event sourcing
- [ ] Add example of using Saga
- [ ] Compare domain events approach with procedure approach https://enterprisecraftsmanship.com/posts/domain-events-simple-reliable-solution/ 
- [ ] Add example of using https://opensource.googleblog.com/2023/03/introducing-service-weaver-framework-for-writing-distributed-applications.html https://serviceweaver.dev/
- [ ] Add example of using https://github.com/zitadel/zitadel

## Documentation
[README.md](docs%2FREADME.md)

## References
- DDD:
  - https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example
  - Vladimir Khorikov:
    - https://github.com/vkhorikov/DddWorkshop
    - https://enterprisecraftsmanship.com/posts/domain-model-purity-completeness/
    - https://enterprisecraftsmanship.com/posts/domain-events-simple-reliable-solution/
    - https://enterprisecraftsmanship.com/posts/always-valid-domain-model/
  - Implementing Domain-Driven Design - Vaughn Vernon
  - Domain-Driven Design: Tackling Complexity in the Heart of Software - Eric Evans
- Tests:
  - Unit Testing book - Vladimir Khorikov https://enterprisecraftsmanship.com/book
- Go best practices
  - https://dave.cheney.net/2019/07/09/clear-is-better-than-clever
  - https://dave.cheney.net/2016/08/20/solid-go-design
  - https://go.dev/doc/effective_go