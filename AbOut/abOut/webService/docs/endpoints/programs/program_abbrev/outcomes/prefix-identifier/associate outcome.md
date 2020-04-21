**Associate Outcome with program**
----
  modifies the program_outcomes table and adds a new association between program and outcome

* **URL**

  /programs/{program abbrev}/outcomes/{prefix-identifier}

* **Method:**

   `POST` 
  
*  **URL Params**

   **Required:**
 
   `program abbrev=[char(5)]`

   `prefix-identifier=[char(11)]`

* **Data Params**

  **Required:**

   `start_semseter=[char(11)]`

  **optional:**

   `end_semseter=[char(11)]`
  

* **Success Response:**
  
  * **Code:** 201 CREATED <br />
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "invalid program abbrev" }`

  OR

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "invalid prefix-identifier" }`






* **Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

* **Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 