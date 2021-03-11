# restapi - Go Language V-1.14
Author: Silambarasan

Programming Lang: RestAPI developped using Go language using version 1.14.

Features:

1. /api/version - To fetch the version of go language you are running. GET Method should be used with it
2. /api/encrypt - To encrypt the json file this api can be used. POST Method should be used with it
3. /api/decrypt - TO decrypt the file which was encrypted by api "/api/encrypt". GET Method should be used with it.

Webserver: The router/middleware used for this is mux can be get it from "https://github.com/gorilla/mux".

Input file: Input file named "input.json" can be used to with your POST call

When you call encrypt api with your input json file, it will create new file with name "encrypted-file.json" and store all the encrypted content

Decrypt API use the file "encrypted-file.json" and take the content out of it to decrypt it and return that value in your console.








