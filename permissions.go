package chalk

import (
	"fmt"
	"strings"

	authv1 "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
	"google.golang.org/protobuf/proto"
)

var permissionNameToEnum map[string]authv1.Permission = func() map[string]authv1.Permission {
	m := make(map[string]authv1.Permission)
	permissionEnum := authv1.File_chalk_auth_v1_permissions_proto.Enums().ByName("Permission")
	values := permissionEnum.Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		// Skip PERMISSION_UNSPECIFIED (value 0)
		if value.Number() == 0 {
			continue
		}
		if slug := proto.GetExtension(value.Options(), authv1.E_Slug); slug != nil {
			m[slug.(string)] = authv1.Permission(value.Number())
		}
	}
	return m
}()

func PermissionsFromStrings(permissions []string) ([]authv1.Permission, error) {
	var converted []authv1.Permission
	var invalidPermissions []string

	for _, p := range permissions {
		if enum, valid := permissionNameToEnum[p]; valid {
			converted = append(converted, enum)
		} else {
			invalidPermissions = append(invalidPermissions, p)
		}
	}

	if len(invalidPermissions) > 0 {
		// Quote each invalid permission for clarity
		quotedPerms := make([]string, len(invalidPermissions))
		for i, perm := range invalidPermissions {
			quotedPerms[i] = fmt.Sprintf("%q", perm)
		}

		if len(invalidPermissions) == 1 {
			return nil, fmt.Errorf("invalid permission: %s", quotedPerms[0])
		}
		return nil, fmt.Errorf("invalid permissions: %s", strings.Join(quotedPerms, ", "))
	}

	return converted, nil
}
