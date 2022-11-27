#!/bin/bash
echo -e "Generating mocks using Mockery..."

echo -e "\nGenerating mocks for repositories..."
cd "${0%/*}/src/repository"

mockery  -name UserRepository -case underscore

echo -e "\nMocks generation complete"