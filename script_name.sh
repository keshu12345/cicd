#!/bin/bash

# Define the deployment name and namespace
DEPLOYMENT_NAME=my-app
NAMESPACE=default

# Get the current revision
REVISION=$(kubectl get deployments -n $NAMESPACE -o jsonpath='{.items[0].metadata.annotations.deployment\.kubernetes\.io/revision}')

# Rollback to the previous revision
kubectl rollout undo deployment $DEPLOYMENT_NAME -n $NAMESPACE --to-revision=$((REVISION-1))