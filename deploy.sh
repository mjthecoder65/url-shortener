#!/bin/bash


gcloud run deploy url-shortener \
    --image asia-northeast3-docker.pkg.dev/vegax-429008/url-shortener/url-shortener:v1 \
    --platform managed \
    --region asia-northeast3 \
    --concurrency=80 \
    --service-account=cloud-run-runtime@vegax-429008.iam.gserviceaccount.com \
    --allow-unauthenticated \
    --verbosity=debug
