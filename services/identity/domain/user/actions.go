package user

import "time"

// --- ESTADOS VITALES ---

func (user *User) Activate() {
	user.Status = Active
	user.UpdatedAt = time.Now()
}

func (user *User) Deactivate() {
	user.Status = Inactive
	user.UpdatedAt = time.Now()
}

func (user *User) Suspend() {
	user.Status = Suspended
	user.UpdatedAt = time.Now()
}

func (user *User) SoftDelete() {
	user.Status = Deleted
	now := time.Now()
	user.UpdatedAt = now
	user.DeletedAt = &now
}

// --- GESTIÓN DE GRUPOS (Nombres Descriptivos) ---

func (user *User) AddGroup(groupID string) {
	for _, existingGroupID := range user.Groups {
		if existingGroupID == groupID {
			return // Ya pertenece al grupo
		}
	}
	user.Groups = append(user.Groups, groupID)
	user.UpdatedAt = time.Now()
}

func (user *User) RemoveGroup(groupID string) {
	var remainingGroups []string
	for _, existingGroupID := range user.Groups {
		if existingGroupID != groupID {
			remainingGroups = append(remainingGroups, existingGroupID)
		}
	}
	user.Groups = remainingGroups
	user.UpdatedAt = time.Now()
}

// --- GESTIÓN DE ROLES (Agregado para cumplir requerimiento) ---

func (user *User) AddRole(roleID string) {
	for _, existingRoleID := range user.Roles {
		if existingRoleID == roleID {
			return // Ya tiene el rol
		}
	}
	user.Roles = append(user.Roles, roleID)
	user.UpdatedAt = time.Now()
}

func (user *User) RemoveRole(roleID string) {
	var remainingRoles []string
	for _, existingRoleID := range user.Roles {
		if existingRoleID != roleID {
			remainingRoles = append(remainingRoles, existingRoleID)
		}
	}
	user.Roles = remainingRoles
	user.UpdatedAt = time.Now()
}

// --- ACCESO ---

func (user *User) UpdateLastLogin() {
	user.LastLogin = time.Now()
}

func (user *User) ChangePassword(hashedPassword string) {
	user.Password = hashedPassword
	user.UpdatedAt = time.Now()
}
