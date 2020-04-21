**Disassociate Outcome with program**
----
  modifies the program_outcomes table and removes a current association between program and outcome

* **URL**

  /programs/{program abbrev}/outcomes/{prefix}/{identifier}

* **Method:**

   `DELETE` 
  
*  **URL Params**

   **Required:**
 
   `program abbrev=[char(5)]`

   `prefix-identifier=[char(11)]`
  
* **Success Response:**
  
  * **Code:** 200 OK <br />
 
* **Error Response:**

 *  **Code:** 403 FORBIDDEN

    **Content:** `{ error : "Insuffecient permissions" }`

OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "invalid program abbrev" }`

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "invalid prefix-identifier" }`

  






* **Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

* **Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 