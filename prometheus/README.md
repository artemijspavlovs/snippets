`compose.yaml` - docker compose file to run a prometheus service locally

`iac` - directory containing Terraform code to deploy an example VM to GCP. The example is intended to use the public IP
of this VM for the prometheus service to make it easier to test it locally. A production version shouldn't use the
public IP for exporting the metrics


