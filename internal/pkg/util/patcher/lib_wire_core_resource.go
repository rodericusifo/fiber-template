package patcher

import (
	lib_wire_core_resource_user "github.com/rodericusifo/fiber-template/lib/wire/core/resource/user"
	lib_wire_core_resource_role "github.com/rodericusifo/fiber-template/lib/wire/core/resource/role"
	lib_wire_core_resource_permission "github.com/rodericusifo/fiber-template/lib/wire/core/resource/permission"
	lib_wire_core_resource_role_permission "github.com/rodericusifo/fiber-template/lib/wire/core/resource/role_permission"
)

var (
	UserResource = lib_wire_core_resource_user.UserResource
	RoleResource = lib_wire_core_resource_role.RoleResource
	PermissionResource = lib_wire_core_resource_permission.PermissionResource
	RolePermissionResource = lib_wire_core_resource_role_permission.RolePermissionResource
)
