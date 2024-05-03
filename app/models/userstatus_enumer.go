// Code generated by "enumer -type=UserStatus -json -transform=snake"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _UserStatusName = "user_status_createduser_status_activeuser_status_deleted"

var _UserStatusIndex = [...]uint8{0, 19, 37, 56}

const _UserStatusLowerName = "user_status_createduser_status_activeuser_status_deleted"

func (i UserStatus) String() string {
	if i < 0 || i >= UserStatus(len(_UserStatusIndex)-1) {
		return fmt.Sprintf("UserStatus(%d)", i)
	}
	return _UserStatusName[_UserStatusIndex[i]:_UserStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UserStatusNoOp() {
	var x [1]struct{}
	_ = x[UserStatusCreated-(0)]
	_ = x[UserStatusActive-(1)]
	_ = x[UserStatusDeleted-(2)]
}

var _UserStatusValues = []UserStatus{UserStatusCreated, UserStatusActive, UserStatusDeleted}

var _UserStatusNameToValueMap = map[string]UserStatus{
	_UserStatusName[0:19]:       UserStatusCreated,
	_UserStatusLowerName[0:19]:  UserStatusCreated,
	_UserStatusName[19:37]:      UserStatusActive,
	_UserStatusLowerName[19:37]: UserStatusActive,
	_UserStatusName[37:56]:      UserStatusDeleted,
	_UserStatusLowerName[37:56]: UserStatusDeleted,
}

var _UserStatusNames = []string{
	_UserStatusName[0:19],
	_UserStatusName[19:37],
	_UserStatusName[37:56],
}

// UserStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UserStatusString(s string) (UserStatus, error) {
	if val, ok := _UserStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _UserStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UserStatus values", s)
}

// UserStatusValues returns all values of the enum
func UserStatusValues() []UserStatus {
	return _UserStatusValues
}

// UserStatusStrings returns a slice of all String values of the enum
func UserStatusStrings() []string {
	strs := make([]string, len(_UserStatusNames))
	copy(strs, _UserStatusNames)
	return strs
}

// IsAUserStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UserStatus) IsAUserStatus() bool {
	for _, v := range _UserStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for UserStatus
func (i UserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserStatus
func (i *UserStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("UserStatus should be a string, got %s", data)
	}

	var err error
	*i, err = UserStatusString(s)
	return err
}
