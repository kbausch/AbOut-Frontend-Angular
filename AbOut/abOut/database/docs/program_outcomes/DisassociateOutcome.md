# Disassociate Outcome

## Name, IN, and OUT parameters

### **Name:** ```program_outcomes__disassociate_outcome__sp```

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
- The input program and outcome are not already associated

## Description

Removes an association between a program and an outcome.