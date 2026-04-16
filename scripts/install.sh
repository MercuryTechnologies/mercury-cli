#!/bin/sh
# Mercury CLI installer
# Usage: curl -fsSL https://cli.mercury.com/install.sh | sh
#
# Environment variables:
#   MERCURY_VERSION     - Pin a specific version (e.g. "0.3.0"). Default: latest.
#   MERCURY_INSTALL_DIR - Override install directory. Default: ~/.local/bin

set -eu

BASE_URL="https://cli.mercury.com"
BINARY_NAME="mercury"
DEFAULT_INSTALL_DIR="$HOME/.local/bin"

# --- Helpers ---

info() {
  printf '%s\n' "$@"
}

error() {
  printf 'Error: %s\n' "$@" >&2
  exit 1
}

need_cmd() {
  if ! command -v "$1" > /dev/null 2>&1; then
    error "required command not found: $1"
  fi
}

# --- Platform detection ---

detect_os() {
  case "$(uname -s)" in
    Darwin) echo "macos" ;;
    Linux)  echo "linux" ;;
    *)      error "unsupported operating system: $(uname -s)" ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64)   echo "amd64" ;;
    aarch64|arm64)   echo "arm64" ;;
    *)               error "unsupported architecture: $(uname -m)" ;;
  esac
}

# --- Checksum verification ---

verify_checksum() {
  archive_file="$1"
  checksums_file="$2"
  archive_basename="$(basename "$archive_file")"

  expected="$(grep "$archive_basename" "$checksums_file" | awk '{print $1}')"
  if [ -z "$expected" ]; then
    error "checksum not found for $archive_basename in checksums file"
  fi

  if command -v sha256sum > /dev/null 2>&1; then
    actual="$(sha256sum "$archive_file" | awk '{print $1}')"
  elif command -v shasum > /dev/null 2>&1; then
    actual="$(shasum -a 256 "$archive_file" | awk '{print $1}')"
  else
    error "no sha256sum or shasum command found — cannot verify checksum"
  fi

  if [ "$expected" != "$actual" ]; then
    error "checksum mismatch for $archive_basename (expected $expected, got $actual)"
  fi
}

# --- Archive extraction ---

extract_archive() {
  archive_file="$1"
  dest_dir="$2"

  case "$archive_file" in
    *.tar.gz) tar -xzf "$archive_file" -C "$dest_dir" ;;
    *.zip)    unzip -oq "$archive_file" -d "$dest_dir" ;;
    *)        error "unknown archive format: $archive_file" ;;
  esac
}

# --- PATH configuration ---

add_to_path() {
  install_dir="$1"

  # Check if already in PATH
  case ":$PATH:" in
    *":$install_dir:"*) return ;;
  esac

  shell_name="$(basename "${SHELL:-sh}")"
  profile=""

  case "$shell_name" in
    zsh)
      profile="$HOME/.zshrc"
      ;;
    bash)
      # Prefer .bashrc on Linux, .bash_profile on macOS
      if [ -f "$HOME/.bashrc" ]; then
        profile="$HOME/.bashrc"
      elif [ -f "$HOME/.bash_profile" ]; then
        profile="$HOME/.bash_profile"
      else
        profile="$HOME/.profile"
      fi
      ;;
    fish)
      # fish uses a different config mechanism
      fish_config="$HOME/.config/fish/conf.d/mercury.fish"
      mkdir -p "$(dirname "$fish_config")"
      printf 'fish_add_path %s\n' "$install_dir" > "$fish_config"
      info "Added $install_dir to fish PATH via $fish_config"
      return
      ;;
    *)
      profile="$HOME/.profile"
      ;;
  esac

  if [ -n "$profile" ]; then
    printf '\n# Mercury CLI\nexport PATH="%s:$PATH"\n' "$install_dir" >> "$profile"
    info "Added $install_dir to PATH in $profile"
  fi
}

# --- Main ---

main() {
  need_cmd curl
  need_cmd uname

  os="$(detect_os)"
  arch="$(detect_arch)"

  # Determine archive format
  case "$os" in
    macos) ext="zip"; need_cmd unzip ;;
    linux) ext="tar.gz"; need_cmd tar ;;
  esac

  # Determine version
  if [ -n "${MERCURY_VERSION:-}" ]; then
    version="$MERCURY_VERSION"
  else
    info "Fetching latest version..."
    version="$(curl -fsSL "$BASE_URL/VERSION")"
  fi

  if [ -z "$version" ]; then
    error "could not determine version to install"
  fi

  info "Installing $BINARY_NAME v${version} (${os}/${arch})..."

  # Determine install directory
  install_dir="${MERCURY_INSTALL_DIR:-$DEFAULT_INSTALL_DIR}"

  # Set up temp directory
  tmp_dir="$(mktemp -d)"
  trap 'rm -rf "$tmp_dir"' EXIT

  # Download archive and checksums
  archive_name="${BINARY_NAME}_${version}_${os}_${arch}.${ext}"
  checksums_name="${BINARY_NAME}_${version}_checksums.txt"

  archive_url="$BASE_URL/v${version}/${archive_name}"
  checksums_url="$BASE_URL/v${version}/${checksums_name}"

  info "Downloading ${archive_name}..."
  curl -fsSL "$archive_url" -o "$tmp_dir/$archive_name"

  info "Downloading checksums..."
  curl -fsSL "$checksums_url" -o "$tmp_dir/$checksums_name"

  # Verify checksum
  info "Verifying checksum..."
  verify_checksum "$tmp_dir/$archive_name" "$tmp_dir/$checksums_name"

  # Extract
  extract_archive "$tmp_dir/$archive_name" "$tmp_dir"

  # Install binary
  mkdir -p "$install_dir"
  cp "$tmp_dir/$BINARY_NAME" "$install_dir/$BINARY_NAME"
  chmod +x "$install_dir/$BINARY_NAME"

  # Update PATH if needed
  add_to_path "$install_dir"

  info ""
  info "Mercury CLI v${version} installed to ${install_dir}/${BINARY_NAME}"
  info ""

  # Check if binary is reachable
  if command -v "$BINARY_NAME" > /dev/null 2>&1; then
    info "Run 'mercury --help' to get started."
  else
    info "To get started, restart your shell or run:"
    info "  export PATH=\"${install_dir}:\$PATH\""
    info ""
    info "Then run 'mercury --help'."
  fi
}

main "$@"
