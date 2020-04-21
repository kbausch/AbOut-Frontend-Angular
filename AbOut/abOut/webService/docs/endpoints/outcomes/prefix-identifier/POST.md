**Create Outcome**
----
This endpoint creates a single outcome

**URL**

/outcomes/{prefix}/{identifier}

**Method:**
  
  `POST`
  
**URL Params**

   **Required:**
 
   `prefix=[char(5)]`

   `identifier=[char(5)]`

**Data Params**

   `description=[varchar(300)] `

**Success Response:**
  
**Code:** 201 CREATED


 
**Error Response:**

* **Code:** 400 BAD REQUEST

  **Content:** `{ error : "Outcome not valid" }`

OR

 *  **Code:** 403 FORBIDDEN

    **Content:** `{ error : "Insuffecient permissions" }`

OR

* **Code:** 409 CONFLICT

  **Content:** `{ error : "Outcome already exists" }`

  

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**
