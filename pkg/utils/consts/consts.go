// Package consts hold global types and constants
package consts

import (
	"fmt"
	"strings"
)

// PermissionTypes types of permissions that can be granted
type PermissionTypes struct {
	Create string
	Read   string
	Update string
	Delete string
	List   string
	Assign string
	Upload string
}

// Tables names of all the tables in the db
type Tables struct {
	Users       string
	Roles       string
	Permissions string
}

//type roles struct {
//Admin string
//User  string
//}

// Tablenames the names of the tables in the server
func Tablenames() Tables {
	return Tables{
		Users:       "users",
		Roles:       "roles",
		Permissions: "permissions",
	}
}

// Permissions has the types of permissions that can be assigned
func Permissions() PermissionTypes {
	return PermissionTypes{
		Create: "create:%s",
		Read:   "read:%s",
		Update: "update:%s",
		Delete: "delete:%s",
		List:   "list:%s",
		Assign: "assign:%s",
		Upload: "upload:%s",
	}
}

// FormatPermissionTag returns a string formatted action:entity permission
func FormatPermissionTag(action string, entity string) string {
	return fmt.Sprintf(action, entity)
}

// FormatPermissionDesc returns a string with the description of the
// action:entity permission
func FormatPermissionDesc(action string, entity string) string {
	return "Allows the user to " +
		strings.ReplaceAll(FormatPermissionTag(action, entity), ":", " ")
}
