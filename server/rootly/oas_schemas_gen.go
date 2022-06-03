// Code generated by ogen, DO NOT EDIT.

package rootly

type BearerAuth struct {
	Token string
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// ScimUsersGetOK is response for ScimUsersGet operation.
type ScimUsersGetOK struct{}

func (*ScimUsersGetOK) scimUsersGetRes() {}

// ScimUsersGetUnauthorized is response for ScimUsersGet operation.
type ScimUsersGetUnauthorized struct{}

func (*ScimUsersGetUnauthorized) scimUsersGetRes() {}

// ScimUsersIDDeleteNoContent is response for ScimUsersIDDelete operation.
type ScimUsersIDDeleteNoContent struct{}

// ScimUsersIDGetOK is response for ScimUsersIDGet operation.
type ScimUsersIDGetOK struct{}

// ScimUsersIDPatchOK is response for ScimUsersIDPatch operation.
type ScimUsersIDPatchOK struct{}

// ScimUsersPostCreated is response for ScimUsersPost operation.
type ScimUsersPostCreated struct{}