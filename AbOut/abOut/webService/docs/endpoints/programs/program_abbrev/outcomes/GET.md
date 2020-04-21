**View all outcomes in a program**
----
Lists all current outcomes with a program.

**URL**

  /programs/{program abbrev}/outcomes

**Method:**

`GET` 

*  **URL Params**

   **Required:**
 
   `program abbrev=[char(5)]`

**Success Response:**

**Code:** 200

**Error Response:**

 * **Code:** 404 PAGE NOT FOUND

   **Content:** `{ error : "Program not found" }`


**Content:** 
```json
[
  {
    prefix : "CAC",
    identifier : "K",
    text : "Students have developed software requirements",
    begin : "Fall2016",
    end : ""
  },
   ...
]
```

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 