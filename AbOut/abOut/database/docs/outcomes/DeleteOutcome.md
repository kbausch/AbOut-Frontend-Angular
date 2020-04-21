# Delete Outcome

## Name, IN, and OUT parameters

### **Name:** ```outcomes__delete_outcome__sp```

### **IN:**

- "pre" - The prefix of the outcome to be deleted
- "idnt" - The identifier of the outcome to be deleted

### **OUT:**

- "status" - Is set to 1 if an error occured, otherwise is set to 0
- "error_message" - A message briefly describing the error, NULL if no error occured

### **Possible Errors:**

- An outcome with the input prefix and identifer has an association with a program. Deletion of such outcomes is not allowed.
- No outcome exists in the system with the input prefix and identifer.

## Description

Deletes an outcome from the outcomes table. The procedure will only do so if that outcome exists already, and doesn't have an association with a program. The two last parameters are status and error_message. When calling the stored procedure, those two parameters have their values changed after the stored procedure ends. To see what their new values are you have to query them. You can do that through a query such as, "SELECT status, error_message;".