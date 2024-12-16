# Structure

+ build
+ config
+ deployment
- cmd
  - http_server
  - db_migrations
  - telegram_bot
+ core
  + account
  + crm
  - identity
    - commands
    - queries
    - internal
      - user
    - adapters
      - postgres
        - migrations
- infrastructure
  postgres



# Cons
- конфликты (напр в миграциях) при одновременной работе
