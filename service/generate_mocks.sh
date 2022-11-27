#!/bin/bash
echo -e "Generating mocks using Mockery..."

echo -e "\nGenerating mocks for repository..."
cd "${0%/*}/src/repository"
mockery  -all

echo -e "\nMocks generation complete"