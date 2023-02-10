# Grafana

### Run Locally

```shell
docker compose up
```

### Provided Example Config files

`grafana-google-oauth.ini` - config file to enable authentication with Google via `auth.google`

replace `client_id`, `client_secret` with the values from the [Created GCP OAuth Client](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/google/#create-google-oauth-keys)

> The `Authorized JavaScript origins` in the GCP Credential created for OAuth authentication should contain a `https://domain.com` record

> The `Authorized redirect URIs` in the GCP Credential created for OAuth authentication should contain a `https://domain.com/login/google` record

---

`grafana-generic-google-oauth-rolemapping.ini` - config file that assigns `Admin` role to everyone that logs in with a specific domain. Uses `auth.generic_oauth`

replace `client_id`, `client_secret` with the values from the [Created GCP OAuth Client](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/google/#create-google-oauth-keys) domain under `role_attribute_path` with the domain of your organization.

> The `Authorized JavaScript origins` in the GCP Credential created for OAuth authentication should contain a `https://domain.com` record

> The `Authorized redirect URIs` in GCP Credential created for OAuth authentication should contain a `domain.com/login/generic_oauth` record