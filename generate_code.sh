#!/bin/bash
echo -e "Generating project code...\n"

${0%/*}/service/generate_mocks.sh

echo -e "\nFinished generating code!"