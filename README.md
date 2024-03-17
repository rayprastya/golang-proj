Golang project structure


root project will be on shopifyx


the cmd will act as a main file

config will be stored as an env so we can use safely accross the app

internal
this will cover all the api works like

api/handler => the handler of api
model => the struct of the model that later will be use similar as serializer
repository => this will act as a file that hold many query or this file is the query action of the api function

pkg
this will act as a component or if any of the code has similar function, stored it here

scripts
this is just a folder to store automated script like migrating etc

tests
to store testcase