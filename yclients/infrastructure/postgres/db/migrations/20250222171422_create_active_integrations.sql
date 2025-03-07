-- migrate:up
CREATE TABLE active_integrations (
  id serial PRIMARY KEY,
  account_id integer NOT NULL,
  ext_company_id varchar(255) NOT NULL,
  company_name varchar(255) NOT NULL,
  confirmation_time timestamp NOT NULL
);
-- migrate:down
DROP TABLE active_integrations;

