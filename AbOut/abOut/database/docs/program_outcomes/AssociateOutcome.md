# Associate Outcome

## Name, IN, and OUT parameters

### **Name:** ```program_outcomes__associate_outcome__sp```

### **IN:**

- "program_abbrev" - The abbreviation of the program
- "outcome_prefix" - The outcomes associated prefix
- "outcome_identifier" - The outcomes associated identifier number

### **OUT:**

- "status" - is set to 1 if an error occured, otherwise is set to 0
- "error_message" - a message briefly describing the error, NULL if no error occured

### **Possible Errors:**

- Invalid/Nonexistant program abbreviation
- Invalid/Nonexistant outcome prefix and identifier combination
- The input program and outcome are already associated

## Description

Associates an existing program with an existing outcome.  It does this by inserting the program and outcome id's into the program_outcomes table along with the current interval.