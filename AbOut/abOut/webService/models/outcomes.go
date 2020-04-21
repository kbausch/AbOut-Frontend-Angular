package models

// OutcomesRepository defines the interactions with the database for outcomes.
type OutcomesRepository interface {
	GetOutcomes() (Outcomes, error)
	GetOutcome(p string, i string) (Outcome, error)
	GetOutcomesInProgram(a string) (Outcomes, error)
	CreateOutcome(p string, i string, c string) error
	UpdateOutcome(p string, i string, c string) error
	DeleteOutcome(p string, i string) error
}

// OutcomesRepo implements OutcomesRepository for use in the endpoints.
type OutcomesRepo struct {
}

// Outcome is a model for the Outcomes table.
type Outcome struct {
	Prefix     string `json:"prefix"`     //The outcome prefix.
	Identifier string `json:"identifier"` //The identifier of the outcome.
	Text       string `json:"text"`       //The description of the outcome.
	Begin      string `json:"begin"`      //The name of the begin semester.
	End        string `json:"end"`        //The name of the end semester.

}

// Outcomes is a list of Outcome.
type Outcomes []Outcome

// GetOutcomes returns a list of outcomes
func (OutcomesRepo) GetOutcomes() (Outcomes, error) {
	rows, err := db.Query("SELECT * FROM outcomes__list_all__vw")
	if err != nil {
		return nil, err
	}
	outcomes := []Outcome{}
	for rows.Next() {
		var prefix string
		var identifier string
		var text string
		var begin string
		var end string
		err2 := rows.Scan(&prefix, &identifier, &text, &begin, &end)
		if err2 != nil {
			return nil, err2
		}
		outcome := Outcome{prefix, identifier, text, begin, end}
		outcomes = append(outcomes, outcome)
	}
	return outcomes, nil
}

// GetOutcomesInProgram returns a list of outcomes
func (OutcomesRepo) GetOutcomesInProgram(a string) (Outcomes, error) {
	rows, err := db.Query("programs_outcomes__outcomes_in_program__sp(?)", a)
	if err != nil {
		return nil, err
	}
	outcomes := []Outcome{}
	for rows.Next() {
		var prefix string
		var identifier string
		var text string
		var begin string
		var end string
		err2 := rows.Scan(&prefix, &identifier, &text, &begin, &end)
		if err2 != nil {
			return nil, err2
		}
		outcome := Outcome{prefix, identifier, text, begin, end}
		outcomes = append(outcomes, outcome)
	}
	return outcomes, nil
}

// GetOutcome returns a outcome
func (OutcomesRepo) GetOutcome(p string, i string) (Outcome, error) {
	outcome := Outcome{}
	row, err := db.Query("CALL outcomes__list_one__sp(?, ?)", p, i)
	if err != nil {
		return outcome, err
	}
	var prefix string
	var identifier string
	var text string
	var begin string
	var end string
	row.Next()
	err2 := row.Scan(&prefix, &identifier, &text, &begin, &end)
	if err2 != nil {
		return outcome, err2
	}
	outcome = Outcome{prefix, identifier, text, begin, end}
	return outcome, nil
}

// CreateOutcome adds a outcome to the database
func (OutcomesRepo) CreateOutcome(p string, i string, c string) error {
	_, err := db.Query("CALL outcomes__create_outcome__sp(?,?,?)", p, i, c)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteOutcome removes an outcome from the database
func (OutcomesRepo) DeleteOutcome(p string, i string) error {
	_, err := db.Query("CALL outcomes__delete_outcome__sp(?,?)", p, i)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateOutcome updates a outcome in the database
func (OutcomesRepo) UpdateOutcome(p string, i string, c string) error {
	_, err := db.Query("CALL outcomes__update_outcome__sp(?,?,?)", p, i, c)
	if err != nil {
		return err
	} else {
		return nil
	}
}
