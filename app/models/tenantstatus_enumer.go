// Code generated by "enumer -type=TenantStatus -json -transform=snake"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TenantStatusName = "tenant_status_createdtenant_status_activetenant_status_deleted"

var _TenantStatusIndex = [...]uint8{0, 21, 41, 62}

const _TenantStatusLowerName = "tenant_status_createdtenant_status_activetenant_status_deleted"

func (i TenantStatus) String() string {
	if i < 0 || i >= TenantStatus(len(_TenantStatusIndex)-1) {
		return fmt.Sprintf("TenantStatus(%d)", i)
	}
	return _TenantStatusName[_TenantStatusIndex[i]:_TenantStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TenantStatusNoOp() {
	var x [1]struct{}
	_ = x[TenantStatusCreated-(0)]
	_ = x[TenantStatusActive-(1)]
	_ = x[TenantStatusDeleted-(2)]
}

var _TenantStatusValues = []TenantStatus{TenantStatusCreated, TenantStatusActive, TenantStatusDeleted}

var _TenantStatusNameToValueMap = map[string]TenantStatus{
	_TenantStatusName[0:21]:       TenantStatusCreated,
	_TenantStatusLowerName[0:21]:  TenantStatusCreated,
	_TenantStatusName[21:41]:      TenantStatusActive,
	_TenantStatusLowerName[21:41]: TenantStatusActive,
	_TenantStatusName[41:62]:      TenantStatusDeleted,
	_TenantStatusLowerName[41:62]: TenantStatusDeleted,
}

var _TenantStatusNames = []string{
	_TenantStatusName[0:21],
	_TenantStatusName[21:41],
	_TenantStatusName[41:62],
}

// TenantStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TenantStatusString(s string) (TenantStatus, error) {
	if val, ok := _TenantStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TenantStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TenantStatus values", s)
}

// TenantStatusValues returns all values of the enum
func TenantStatusValues() []TenantStatus {
	return _TenantStatusValues
}

// TenantStatusStrings returns a slice of all String values of the enum
func TenantStatusStrings() []string {
	strs := make([]string, len(_TenantStatusNames))
	copy(strs, _TenantStatusNames)
	return strs
}

// IsATenantStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TenantStatus) IsATenantStatus() bool {
	for _, v := range _TenantStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TenantStatus
func (i TenantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TenantStatus
func (i *TenantStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TenantStatus should be a string, got %s", data)
	}

	var err error
	*i, err = TenantStatusString(s)
	return err
}