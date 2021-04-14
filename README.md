# gcppub - extra tool for Google Cloud Pub/Sub

## requirements

Authenticated GCP user account with Pub/Sub permissions.

    gcloud auth application-default login

## publish file content

    gcppub -p my-project-id -t my-topic-short-id -f data-to-send.json

TOPIC must be the short name.

## receive file content

    gcppub -p my-project-id -s my-subscription-short-id -f data-received.json


&copy 2021 ernestmicklei.com