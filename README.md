<p align="center">
  <h1 align="center">
      Url Shortener With Gin & Mongodb & Jwt
  </h1>
  <h3 align="center">
     Version 1.0
  </h3>
  <h4 align="center">
      If you would like to report an issue or request a feature. Direct In My <a href="https://instagram.com/nima._.ism">Instagram</a> or create an issue.
  </h4>
</p>

<br/>
<br/>



## Routes
#### `localhost:8080/{YOUR-LINK-NAME}`:
Send a GET request and redirected :)


#### `localhost:8080/api/v1/CreateLink`:
POST a JSON object like below and in return, get the generated short link:
```JSON
{
    "Name":"Nimaism",
    "Url":"https://nimaism.ir"
}
```

#### Result:
```JSON
{
    "ResMessage": "Success",
    "Data": "localhost:8080/Nimaism"
}
```

#### `localhost:8080/account/v1/Register`:
POST a JSON object like below and account has created:
```JSON
{
    "UserName":"Nimaism",
    "Email":"nimasa036@gmail.com",
    "Pass":"12345678"
}
```

#### Result:
```JSON
{
    "ResMessage": "Success",
    "Data": null
}
```

#### `localhost:8080/account/v1/Login`:
POST a JSON object like below and show token for account:
```JSON
{
    "Email":"nimasa036@gmail.com",
    "Pass":"12345678"
}
```

#### Result:
```JSON
{
    "ResMessage": "Success",
    "Data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im5pbWFzZGEwMzZAZ21haWwuY29tIiwiZXhwIjoxNjYwMzkyMTE0fQ.rP9PCbMS39X5aeJryO5h7pLa_j_AYT-rs4Z2uGN-Pk0"
}
```




## Support
* Direct In My [Instagram](https://instagram.com/nima._.ism) for Support or open an issue via GitHub.
