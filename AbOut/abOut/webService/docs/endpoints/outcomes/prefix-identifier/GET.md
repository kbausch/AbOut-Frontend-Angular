**View Outcome**
----
This endpoint fetches a single outcome

**URL**

/outcomes/{prefix}/{identifier}

**Method:**
  
  `GET`
  
**URL Params**

   **Required:**
 
   `prefix=[char(5)]`

   `identifier=[char(5)]`

**Success Response:**
  
  <_What should the status code be on success and is there any returned data? This is useful when people need to to know what their callbacks should expect!_>

**Code:** 200

**Content:**
```json
{
  prefix: "CAC",
  identifier: "K",
  text: "Students have developed software requirements",
  begin : "Fall2016",
  end : ""
}
```
 
**Error Response:**

**Code:** 404 PAGE NOT FOUND

**Content:** `{ error : "Outcome not found" }`

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 