# 2023-07: Errors in Repositories
My decision:
- An error when the application loses connection with the database or other unexpected error from the database, there will be a panic and the http server will return an internal error.
- An error when id UUID is busy is not a business error, there will be panic too and the http server will return an internal error.

## **Advantages**:
- make easy supporting application tests
- make easy application layer code
- creating and  maintaining these error clauses in application layer is a job with a small profit.

## **Disadvantages**:
- error when id UUID is busy, the application returns 500 internal server error instead of 409 conflict. It is not good because 500 errors mean that a service works incorrectly, but in this case it is not true.

But such errors happen very rarely with probability close to zero [1](#references).  Readable and maintainable code is more profitable instead of handling these error to avoid 500 error in such rare cases.

## References
1. https://en.wikipedia.org/wiki/Universally_unique_identifier