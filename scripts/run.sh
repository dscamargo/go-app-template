#!/bin/bash

set -eo pipefail

echo "Running migrations..."
./app db migrate

echo "Running app"
./app app run