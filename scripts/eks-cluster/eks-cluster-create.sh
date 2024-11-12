#!/bin/bash

# Create a EKS cluster
eksctl create cluster --name hms-monitor-cluster \
  --region us-east-1 \
  --zones=us-east-1a,us-east-1b \
  --with-oidc \
  --nodes=2 \
  --nodes-min=2 \
  --nodes-max=3 \
  --instance-types=t3a.small