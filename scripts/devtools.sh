#!/bin/bash

install_docker() {
  if ! command -v docker >/dev/null 2>&1; then
    echo "Installing Docker..."
    # For Linux
    curl -fsSL https://get.docker.com -o get-docker.sh
    sh get-docker.sh
    rm get-docker.sh
  else
    echo "Docker already installed."
  fi
}
install_docker_compose() {
  if ! command -v docker-compose >/dev/null 2>&1; then
    echo "Installing Docker Compose..."
    pip install docker-compose
  else
    echo "Docker Compose already installed."
  fi
}
install_migrate() {
  if ! command -v migrate >/dev/null 2>&1; then
    echo "Installing migrate CLI..."
  else
    echo "migrate already installed."
  fi
}
case $1 in
  all)
    install_docker
    install_docker_compose
    install_migrate
    ;;
  docker) install_docker ;;
  compose) install_docker_compose ;;
  migrate) install_migrate ;;
  *) echo "Usage: $0 [all|docker|compose|migrate]"; exit 1 ;;
esac
