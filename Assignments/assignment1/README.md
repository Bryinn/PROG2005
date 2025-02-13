# Assignment1

- Render: https://prog2005.onrender.com/
  - PORT: 80
- Workspace: https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2025-workspace/bryinn/assignment1

## Extra info
I believe the only prominent problem would be if the dependant APIs went down. 
I am referring to the locally hosted APIs detailed in the assignment.

I did not find a good way to check the status of the APIs without performing a normal request.
The APIs did not have a status endpoint as far as I saw, so I went with making a normal request that was as slim as I saw possible, to perform the status checks.

## Endpoint spesification
The assignment has the following endpoints desribed in the assignment:
- ``/countryinfo/v1/info/{:two_letter_country_code}{?limit=10}``
- ``/countryinfo/v1/population/{:two_letter_country_code}{?limit={:startYear-endYear}}``
- ``/countryinfo/v1/status/``

Other than this, it checks thecks that the request is not bad as you would expect as far as I am concerned.
The sanitization checks for that the:
- request type is a GET request
- parameter is two characters exactly
- query is in the allowed format.

The only exception is that in my implementation, you can just write two years as filter in the population limit query, and the API will sort them automatically in acending order.
I chose to do it this way, instead of throwing an error, because of that it makes the API easier to use, and more flexible to imput in a non-hazardous way.
