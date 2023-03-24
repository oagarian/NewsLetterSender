# NewsLetterSender

## How to use

* First, add a .env with your credentials
```ssh
.env
SMTP_SERVER=smtp.gmail.com
SMTP_PORT=587
SMTP_EMAIL=yourEmail
SMTP_PASSWORD=youPassword
```

* To the next step, change the variables subject and body in the code;

* Finally, run the following commands in the terminal:

```ssh
go mod tidy

go run main .go
```
