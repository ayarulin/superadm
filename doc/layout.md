## Structure

```
- docs
- build
- config
- deployments
- system
	- users
		- cmd
			- users_app
		- internal
			- registration
			- login
			- password_recovery
		- api
			- users_query_api
			- users_command_api
	- clients
		- cmd
			- yclients_worker
			- period_est_worker
			- clients_app
		- internal
			- core
				- entity
				- command
				- queries
				- services

		- fetch data from YCliens
		- calculate cycles
		- manage appointments

- internal
	- adapters
		- yclients
		- pg_database

- go.mod
```

## 2 

```
- docs
- users
    - cmd
        - users_app
    - deployments
    - config
    - internal
        - registration
        - login
        - password_recovery
    - api
        - users_query_api
        - users_command_api
- clients
    - cmd
        - yclients_worker
        - period_est_worker
        - clients_app
    - internal
        - core
            - entity
            - command
            - queries
            - services

    - fetch data from YCliens
    - calculate cycles
    - manage appointments

- infrastructure
	- adapters
		- yclients
		- pg_database

- go.mod
