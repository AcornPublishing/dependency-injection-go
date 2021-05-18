package get

import (
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/modules/data"
)

var (
	// error thrown when the requested person is not in the database
	errPersonNotFound = errors.New("person not found")
)

// Getter will attempt to load a person.
// It can return an error caused by the data layer or when the requested person is not found
type Getter struct {
}

// Do will perform the get
func (g *Getter) Do(ID int) (*data.Person, error) {
	// load person from the data layer
	person, err := data.Load(ID)
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPersonNotFound
		}
		return nil, err
	}

	return person, err
}
