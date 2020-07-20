#!/bin/bash
terraform validate
terraform init
terraform plan -var-file=minikube.tfvars -out planfile
terraform apply -auto-approve planfile
rm planfile
