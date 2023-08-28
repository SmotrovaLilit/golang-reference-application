# Decisions about testing the application
## 2023-08: Application layer with real database or with mocking the repository layer
I chose to use a real database in application layer tests.
- It makes it easy to do refactoring. We can change repository interface, and it doesn't require changing in application tests.
- Maintaining these tests are easy using http://github.com/testcontainers/testcontainers-go

## 2023-08: How to check test assertion in application layer
An example. We are testing the create program command. We need to know that after command completed the program was created in database.   
Two ways to check it:
1. We can get the created program using orm or execute SQL query and then check that the program was created in database correctly.
```go
cmd := createprogram.NewCommand(
    program.MustNewID("35A530CF-91F3-49DC-BB6D-AC423563541C"),
    program.AndroidPlatformCode,
)
handler.Handle(context.Background(), cmd)

_program = &program.Program{}
err := db.Find(_program, cmd.ID).Error
require.NotNil(t, _program)
require.Equal(t, cmd.ID, _program.ID())
require.Equal(t, cmd.PlatformCode, _program.PlatformCode())
```
2. We can call the open repository method for example findById and check that the program was created correctly.
```go
cmd := createprogram.NewCommand(
    program.MustNewID("35A530CF-91F3-49DC-BB6D-AC423563541C"),
    program.AndroidPlatformCode,
)
handler.Handle(context.Background(), cmd)

_program := programRepository.FindByID(context.Background(), cmd.ID)
require.NotNil(t, _program)
require.Equal(t, cmd.ID, _program.ID())
require.Equal(t, cmd.PlatformCode, _program.PlatformCode())
```
In the first option we use information about the internal realisation of the repository layer. If we change the way data is stored in the database we need to change application layer tests too. So we increase coupling between layers.
For example, a Program can be a difficult aggregate with entities and value objects.
Matching columns in tables with fields in aggregates' entities is the responsibility of the repository layer. If we do it in application tests, it will increase coupling.

The second option requires to create in a repository open methods for testing needs. I think it is better than an increasing coupling between layers.
## 2023-08: Using real database or mocking in repository layer tests
### With mocking
- It doesn't affect refactoring complexity because we can change orm without big changes in the repository tests.
- Tests with mocking a database are easier and more maintainable than tests with a real database, because with a real database we need to prepare data.
Example of such tests [program_repository_test.go](..%2Finternal%2Finfrastructure%2Frepositories%2Fprogram_repository_test.go)
### With real database
- It doesn't affect refactoring complexity because we can change orm without changes in the repository tests.
- Creating and maintaining these are more difficult than tests with mocking a database, because we need to prepare data.
  
An example of such tests is in the repository layer see the solution of https://github.com/SmotrovaLilit/golang-reference-application/issues/1

## 2023-08: Transport layer tests
Decoding request and encoding response I test using unit tests without mocking. This logic is independent of other components.

Test that the handler is settled up correctly I test with a real build and real database. See example here [integration](test%2Fintegration)

**Advantages**
- It makes it easy to do refactoring. The test doesn't need to be changed if we change the application layer realization.
- Maintaining these tests are easy because we don't have a lot of test cases.
