package constants

type patientstatuses int

//go:generate goenums patient_statuses.go
const (
	processing patientstatuses = iota
	in_progress
	aborted
	done
)
