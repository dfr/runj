package jail

import (
	"go.sbk.wtf/runj/state"
)

// Jail represents an existing jail
type Jail interface {
	// Return the jail's JID
	JID() ID

	// Attach attaches the current running process to the jail
	Attach() error

	// Destroy the jail
	Remove() error
}

type jail struct {
	id ID
}

// Return the jail's JID
func (j *jail) JID() ID {
	return j.id
}

// FromID queries the OS for a jail with the specified id.
func FromID(id ID) Jail {
	return &jail{id: id}
}

// FromState queries the OS for the jail associated with a state
// record
func FromState(s *state.State) Jail {
	return FromID(ID(s.JID))
}

// FromName queries the OS for a jail with the specified name
func FromName(name string) (Jail, error) {
	id, err := find(name)
	if err != nil {
		return nil, err
	}
	return &jail{id: id}, nil
}

// Attach attaches the current running process to the jail
func (j *jail) Attach() error {
	return attach(j.id)
}

// Destroy the jail
func (j *jail) Remove() error {
	return remove(j.id)
}
