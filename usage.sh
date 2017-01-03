curl -X GET http://localhost:8080/snacks
#returns list of snacks

curl -X POST http://localhost:8080/coin_slots
#returns 201 status 1234567 and secret access key

#no need to know secret ID. Everyone should be able to put coin
#curl -X PUT http://localhost:8080/coin_slots/1234567/ -body "{"coin": 5}"

#GET information about slot
#curl -X GET http://localhost:8080/coin_slots/1234556


#curl -X POST http://localhost:8080/snacks/43 -body "{"wallet": 123456, "secret": "asdf"}"
#gets single snack
