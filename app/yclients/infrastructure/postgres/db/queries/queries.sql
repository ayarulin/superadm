-- name: GetActiveIntegration :one
SELECT * FROM active_integrations
WHERE id = $1 LIMIT 1;

-- name: CreateActiveIntegration :one
INSERT INTO active_integrations (
  account_id, ext_company_id, company_name, confirmation_time
) VALUES (
  $1, $2, $3, $4
)
 RETURNING id;

