# CreditCardValidatory



Use curl cmd by passing request body to the GET method

curl -X GET -d '{"value":"0,9,9,2,7,3,9,8,7,9,3"}'  "http://127.0.0.1:8089/cred/"
{"Status":"FAIL"}

curl -X GET -d '{"value":"8,9,9,2,7,3,9,8,7,9,3"}'  "http://127.0.0.1:8089/cred/"
{"Status":"OK"}
