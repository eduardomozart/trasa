#!/bin/bash
export NODE_OPTIONS=--openssl-legacy-provider
npm run build && \
rm -rdf /var/trasa/dashboard
mv build/ /var/trasa/dashboard/
