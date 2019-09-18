# go-restful-boilerplate

This is a boilerplate for creating RESTful API services in Golang. This includes,

* Authorization using JWT tokens embedded within Cookies
* CSRF protection
* Standardized JSON responses
* Logging using Logrus
* Configuration using Viper
* Commands using Cobra
* Example of structuring database logic/schemas


## Getting started with a breeze 😏
* Clone the project
* Copy `config.example.yaml` and rename the new version as `config.yaml`
* Configure `config.yaml` to your needs
* Get all of the dependencies by executing `go get ./...` in your project root directory & **RUN**!

## Roadmap
* `/api` is the app entry point, and defines the app's routes
* `/handlers` contains all handlers
* `/helpers` contains helper files & functions to assist handlers
* `/storage` consists of the database client, and database schemas
* `/util` contains the app's console commands and configuration


## Security considerations & recommendations
For this boilerplate it should be noted I have decided to use cookies to embed JWT tokens. The cookie is intended to only store a JWT token, since the more the cookie maintains a client's state, the more distant this becomes from being a RESTful API. However, if for whatever reason you wish to store insensitive data in this cookie, instead embed it within the JWT token to avoid clients tampering with the data.

### Cookies & XSS
By nature, cookies offer more security provided by browsers, and can be secured from XSS attacks through restricting access to cookies by only allowing them to be used by HTTP requests (`httpOnly: true`). This stops any JavaScript executed scripts from accessing these cookies.

### Man in the middle/Sniffing
The attacker may redirect the client to their own server, before the traffic is forwarded to the API. Alternatively, in an unsecured & opened network, an attacker could simply sniff the network. I would recommend using SSL (HTTPS) for transport so this data is encrypted. [Caddy Server](https://caddyserver.com/) will make life much easier.
