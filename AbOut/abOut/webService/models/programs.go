package models

// ProgramsRepository defines the interactions with the database for programs
type ProgramsRepository interface {
	GetPrograms() (Programs, error)
	DisassociateOutcome(a string, p string, i string) error
}

// ProgramsRepo implements ProgramsRepository for use in the endpoints
type ProgramsRepo struct {
}

// Program is a model for the Programs table
type Program struct {
	Abbrev          string `json:"abbrev"`
	Name            string `json:"name"`
	CurrentSemester string `json:"current_semester"`
}

// Programs is a list of Program
type Programs []Program

// GetPrograms returns a list of programs
func (ProgramsRepo) GetPrograms() (Programs, error) {
	rows, err := db.Query("SELECT * FROM programs__list_current__vw")
	if err != nil {
		return nil, err
	}
	programs := []Program{}
	for rows.Next() {
		var abbrev string
		var name string
		var currentSemester string
		err2 := rows.Scan(&abbrev, &name, &currentSemester)
		if err2 != nil {
			return nil, err2
		}
		program := Program{abbrev, name, currentSemester}
		programs = append(programs, program)
	}
	return programs, nil
}

//DisassociateOutcome removes an outcome-program association
func (ProgramsRepo) DisassociateOutcome(a string, p string, i string) error {
	_, err := db.Query("program_outcomes__disassociate_outcome__sp(?,?,?)", a, p, i)
	if err != nil {
		return err
	} else {
		return nil
	}

}
