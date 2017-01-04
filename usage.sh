#!/bin/bash
set -e

echo "GET list of snacks:"
curl -X GET http://localhost:8080/snacks

echo "Create new coin slot:"
curl -X POST http://localhost:8080/coin_slots
#returns 201 status and creates new coin slot

echo "Update coin slot"
curl -X PUT http://localhost:8080/coin_slots/awesome_wright --data '{"Coin": 5}'
#no need to know secret ID. Everyone should be able to put coin

echo "GET information about slot"
curl -X GET http://localhost:8080/coin_slots/awesome_wright


#curl -X POST http://localhost:8080/snacks/43 -body "{"wallet": 123456, "secret": "asdf"}"
#gets single snack
