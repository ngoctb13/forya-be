#!/usr/bin/env sh

SCRIPTPATH="$(
  cd "$(dirname "$0")" || exit 1
  pwd -P
)"

CURRENT_DIR="$SCRIPTPATH"
ROOT_DIR="$(dirname "$CURRENT_DIR")"
PORT="8090"

INFRA_LOCAL_COMPOSE_FILE="$ROOT_DIR/build/docker-compose.dev.yaml"

local_infra() {
  docker compose -f "$INFRA_LOCAL_COMPOSE_FILE" "$@"
}

init() {
  cd "$CURRENT_DIR/.." || exit 1
  goimports -w ./..
  go fmt ./...
}

infra() {
  cmd="$1"
  shift 2>/dev/null || true
  case "$cmd" in
    up)
      local_infra up "$@"
      ;;
    down)
      local_infra down "$@"
      ;;
    build)
      local_infra build "$@"
      ;;
    *)
      echo "Usage: infra {up|down|build} [docker-compose args]"
      ;;
  esac
}

setup_env_variables() {
  set -a
  export $(grep -v '^#' "$ROOT_DIR/build/.base.env" | xargs -0) || true
  . "$ROOT_DIR/build/.base.env"
  set +a
  CONFIG_FILE="$ROOT_DIR/build/app.yaml"
  export CONFIG_FILE
  export PORT="$PORT"
}

api_start() {
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/main.go"
  go run "$ENTRY_FILE" --config-file="$CONFIG_FILE" --port="$PORT"
}

api() {
  cmd="$1"
  shift 2>/dev/null || true
  case "$cmd" in
    start)
      api_start
      ;;
    worker_start)
      worker_start
      ;;
    migrate)
      migrate_db "$@"
      ;;
    *)
      echo "Usage: api {start|worker_start|migrate}"
      ;;
  esac
}

case "$1" in
  init)
    shift
    init "$@"
    ;;
  infra)
    shift
    infra "$@"
    ;;
  api)
    shift
    api "$@"
    ;;
  migrate)
    shift
    migrate_db "$@"
    ;;
  *)
    echo "Usage: $0 {infra|api|migrate|init}"
    ;;
esac
