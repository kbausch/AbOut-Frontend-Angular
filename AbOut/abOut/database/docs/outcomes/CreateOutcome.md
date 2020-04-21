# Create Outcome

## Name, IN, and OUT parameters

### **Name:** ```outcomes__create_outcome__sp```

### **IN:**

- "pre" - The prefix of the outcome to be created
- "idnt" - The identifier of the outcome to be created
- "txt" - The text of the outcome to be created

### **OUT:**

- "status" - Is set to 1 if an error occured, otherwise is set to 0
- "error_message" - A message briefly describing the error, NULL if no error occured

### **Possible Errors:**

- Invalid/Nonexistant prefix text
- An outcome with the input prefix and identifer already exists

## Description

Creates a new outcome in the system with the input prefix, identifer, and text. The two last parameters are status and error_message. When calling the stored procedure, those two parameters have their values changed after the stored procedure ends. To see what their new values are you have to query them. You can do that through a query such as, "SELECT status, error_message;".