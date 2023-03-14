package constant

type ModifierRole = string

const (
	UserModifierRole   = ModifierRole("USER")
	AdminModifierRole  = ModifierRole("ADMIN")
	SystemModifierRole = ModifierRole("SYSTEM")
)

type SubjectType int8

const (
	UserSubjectType = SubjectType(iota)
	SystemSubjectType
)
