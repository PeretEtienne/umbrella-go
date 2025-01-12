package constants

type gender int

//go:generate goenums gender.go
const (
	genderMale gender = iota
	genderFemale
)
