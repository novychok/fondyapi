# Fondy payment system integration example on Go language
1) Payment Accept documentation (https://docs.fondy.eu/ru/docs/page/3/#chapter-3-1)

# How to use
1) Run service/service.go
   - It will run a simple service on port :8080 that will listen for 127.0.0.1/callback

2) Use Ngrok to make Fondy Server Callback URL have acces to your 127.0.0.1/callback
   - You need to set Ngrok generated Link in your Request ServerCallbackURL

3) Run client/client.go
   - It will make a Post request to Fondy APIs, and generate you a response, you can find there the link
     (e.g checkout_url:https://pay.fondy.eu/merchants/5ad6b888f4becb0c33d543d54e57d86c/default/index.html?token=3b7aff32dc40be1437f5217c2c06a0a8bd5bf8fd) that you need to follow to make your payment.
     After payment will done, you can face a Callback in your service/service.go console

# Used libraries
1) Google UUID (github.com/google/uuid)
   - Generating unique uuid for the Order id
2) Structs (github.com/fatih/structs)
   - A good choice to reach the struct fields names, instead of using reflect directly