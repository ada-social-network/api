#!/usr/bin/env bash

IMAGE_LATEST=ghcr.io/ada-social-network/api:latest
IMAGE_VERSION=ghcr.io/ada-social-network/api:"${1}"

printf "Build image with TAG:\n"
printf "  - latest: %s\n" "${IMAGE_LATEST}"
printf "  - version: %s\n" "${IMAGE_VERSION}"

printf "\n--- Start Docker build\n\n"
docker build . -t ghcr.io/ada-social-network/api:latest -t ghcr.io/ada-social-network/api:"${1}" --build-arg VERSION="${1}"
printf "\n\n--- End Docker build\n"

printf "\nYou can now start using your image by running the following command\n\n"
printf "  docker run --rm %s" "${IMAGE_VERSION}"

printf "\n\nYou can push your images by running the following command:\n\n"
printf "  docker push %s\n" "${IMAGE_VERSION}"
printf "  docker push %s\n" "${IMAGE_LATEST}"
