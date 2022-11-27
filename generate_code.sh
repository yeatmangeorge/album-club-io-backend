#!/bin/bash
echo -e "Generating project code...\n"

${0%/*}/domain/generate_mocks.sh

echo -e "\nFinished generating code!"