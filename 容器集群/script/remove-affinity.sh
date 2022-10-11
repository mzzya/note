#! /bin/bash

rm deploy.json
rm -rf temp/
mkdir temp/

node remove-affinity.js

ls temp/

kubectl apply  -f temp/