# gcppub - extra tool for Google Cloud Pub/Sub

# This repo is archived, please use https://github.com/emicklei/gcloudx instead.

## requirements

Authenticated GCP user account with Pub/Sub permissions.

    gcloud auth application-default login

## publish file content

    gcppub -p my-project-id -t my-topic-short-id -f data-to-send.json

## receive file content

    gcppub -p my-project-id -s my-subscription-short-id -f data-received.json


&copy; 2021 ernestmicklei.com MIT License.
