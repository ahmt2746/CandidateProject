# CandidateProject

This is a project of FINFREE Candidate Project.

This project include 3 rest api;

  1. Login; all usr_info saved in code. This api takes username and pasword from request body as json. And check with data which storaged in code and returns JWT Token in response body.
  2. UsrInfo; this api gets request with token and check it. If token valid then returns usr_info(username, firstname, lastname) which storaged local in response body.
  3. UsrPasswordUpdate; this api takes request with token and check it. If token valid then read new password from request body and updates local password.

The postman collection includes three apis added.
