**Update Outcome**
----
This endpoint Updates a single outcome

**URL**

/outcomes/{prefix}/{identifier}

**Method:**
  
  `PUT`
  
**URL Params**

   **Required:**
 
   `prefix=[char(5)]`

   `identifier=[char(5)]`

**Data Params**

   `description=[varchar(300)] `

**Success Response:**
  
**Code:** 200 OK


 
**Error Response:**

* **Code:** 400 BAD REQUEST

  **Content:** `{ error : "Outcome not valid" }`

OR

 *  **Code:** 403 FORBIDDEN

    **Content:** `{ error : "Insuffecient permissions" }`
    
OR

* **Code:** 404 PAGE NOT FOUND

  **Content:** `{ error : "Outcome not found" }`

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**

 