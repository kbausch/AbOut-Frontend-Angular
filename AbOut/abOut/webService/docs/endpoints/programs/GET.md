**View Programs**
----
This endpoint fetches all of the programs associated with the AbOut system and displays them in a list form.

**URL**

/programs

**Method:**

`GET`

**Success Response:**

**Code:** 200

**Content:**
```json
[
  {
    abbrev : "CS",
    name: "Computer Science",
    current_semester: "Fall2020"
  },
  {
    abbrev: "EE",
    name: "Electrical Engineering",
    current_semester: "Fall2020"
   },
   ...
]
```

**Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

**Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 