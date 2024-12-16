# Structure

+ build
+ env
+ deployment
- cmd
  - identity_http_server
  - identity_telegram_bot
+ systems
  + account
  + crm
  - identity
    - internal
      - user
      - database
        - migrations
    exec_connect_user.go
    query_list_users.go

---
# Subsystems

## authorization
  login
  registration

## accounts
## crm

 ## users
  - registration
  - password recovery
  - login

 Account
 User
    accountId
    fullname
    password
    phoneNumber
    type: ACCOUNT_MANAGER | USER

# clients
  - period calculation
  - visits history
  - notifications

  Client
      accountId
      fullname

  Visit



# ycapp
  register user
  receive all clients
  keep up client history

YCUser
  name
  accountId

YCSalon
    ycSalonId
    ycUserId
    name

YCSalonClient
    ycSalonId
    clientId


YCEvent



### Folders

ycapp
  cmd
    yc_app
  internal
    core
    application
      services
      http_client
  api
mainapp
  cmd
  internal
    core
    application

# ---

handlers
systems/modules
yclients
auth


```
registration
  register_user_by_yclients_app()

yclientsapp
  <- registration_api?
  register_user(id, password, salon_id)
  [ - registration.create_account()
    - registration.register_user()
    - yclientsapp.add_salon()
  ]
```

# handler

1. parse/validate params
2. authenticate user
3. execute app logic

```go
	func handler(request) error {
	  params, err = cast(request.params, schema?)

	  tx = txManager.NewTransaction(ctx)

	  usersApp = NewUsersApp()
	  regYclientsApp = NewRegisterYClientsAppFeature(tx)

	  appId, err := regYclientApp.Call(params)
	  err := usersApp.RegisterUser(params + appId)

	  return nil
	}
```

# application

salonsRepo = NewYclients
salonClientsRepo = NewYclients
