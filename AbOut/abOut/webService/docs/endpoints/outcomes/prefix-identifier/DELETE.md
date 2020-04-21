**Delete Outcome**
----
This endpoint deletes a single outcome

**URL**

/outcomes/{prefix}/{identifier}

**Method:**
  
  `DELETE`
  
**URL Params**

   **Required:**
 
   `prefix=[char(5)]`

   `identifier=[char(5)]`

**Success Response:**
  
  

**Code:** 200

 
**Error Response:**

 *  **Code:** 403 FORBIDDEN

    **Content:** `{ error : "Insuffecient permissions" }`

OR

 * **Code:** 404 PAGE NOT FOUND

   **Content:** `{ error : "Outcome not found" }`

OR

 *  **Code:** 405 METHOD NOT ALLOWED

    **Content:** `{ error : "Cannot delete an outcome that is associated to a program" }`

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**
