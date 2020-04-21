# Update Outcome

## Name, IN, and OUT parameters

### **Name:** ```outcomes__update_outcome__sp```

### **IN:**

- "prefix" - The prefix of the outcome to be updated
- "identifier" - The identifier of the outcome to be updated
- "new_text" - The new text of the specified outcome

### **OUT:**

- "status" - Is set to 1 if an error occured, otherwise is set to 0
- "error_message" - A message briefly describing the error, NULL if no error occured

### **Possible Errors:**

- specified outcome doesn't exist, the prefix identifier combination does not exist

## Description

Updates the text of the specified outcome.  Locates the specified outcome by joining the outcomes and prefix tables.