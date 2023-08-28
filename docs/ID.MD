# 2023-08: The implementation of ID value object in domain entities
## Solutions
💡 I chose the 3-d solution.

[1. For every entity create their own value object for ID](#1-for-every-entity-create-their-own-value-object-for-id)

[2. General value object for every entity](#2-general-value-object-for-every-entity)

[3. Use base generic error and create a custom ID type for every entity](#3-use-base-generic-error-and-create-a-custom-id-type-for-every-entity)

###  1. For every entity create their own value object for ID

❌ The solution requires writing a similar code for ID in every entity. It increases code support because we need to support code and tests for a similar logic in every entity.
✅ The solution allows to use own value object ID for every entity, which decreases mistakes in code. For example, we can't use program ID instead of version ID.

Example of the solution:
```go
package program

import (
	"github.com/google/uuid"
	"reference-application/internal/pkg/errors"
)

// ErrInvalidID is an errors for invalid program id.
var ErrInvalidID = errors.New("invalid program id", "INVALID_PROGRAM_ID")

// ID is a type for program id.
type ID struct {
	id uuid.UUID
}

// NewID is a constructor for ID.
func NewID(raw string) (ID, error) {
	id, err := uuid.Parse(raw)
	if err != nil {
		return ID{}, ErrInvalidID
	}
	return ID{id: id}, nil
}

// MustNewID is a constructor for ID.
// It panics if the given raw string is invalid.
func MustNewID(raw string) ID {
	id, err := NewID(raw)
	if err != nil {
		panic(err)
	}
	return id
}

// String returns a string representation of ID.
func (i ID) String() string { return i.id.String() }

```
In every new entity, we need to copy this code and change "program" to the name of a new entity.

[⬆ To top](#solutions)

###  2. General value object for every entity
❌ The problem with this solution is a compiler doesn't notice if a developer makes a mistake and uses, for example, program ID instead of version ID:
✅ We don't need to support code and tests for a similar logic in every entity.

```go
program = program.MustNew(id, code)
version = version.New(program.id, ....)
...

package version

func New(versionId ID, programID ID....)
...
```
[⬆ To top](#solutions)

### 3. Use base generic error and create a custom ID type for every entity
```go
func New(versionId VersionID, programID ProgramID....)
```
✅ The solution allows to use own value object ID for every entity like in the first solution.  
✅ Also, the solution decreases the code for creating custom IDs for every entity.

Example of custom ID:
```
package version

import (
	"reference-application/internal/pkg/id"
)

var (
	NewID     = id.New[*ID]
	MustNewID = id.MustNew[*ID]
)

// ID is a type for program version id.
type ID struct {
	id.Base
}
```
We can copy this code for another entity ID, and we don't need to change something in this code for the another entity.

[⬆ To top](#solutions)