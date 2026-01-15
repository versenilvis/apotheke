#!/bin/sh
set -e

# Apotheke installer
# Usage: curl -sS https://raw.githubusercontent.com/verse/apotheke/main/install.sh | sh

REPO="versenilvis/apotheke"
BIN_DIR="${BIN_DIR:-$HOME/.local/bin}"

main() {
    echo "Installing apotheke..."

    arch=$(get_arch)
    echo "Detected architecture: ${arch}"

    tmp_dir=$(mktemp -d)
    trap "rm -rf ${tmp_dir}" EXIT
    cd "${tmp_dir}"

    download_url=$(get_download_url "${arch}")
    if [ -z "${download_url}" ]; then
        err "Could not find release for architecture: ${arch}"
    fi
    echo "Downloading: ${download_url}"

    if command -v curl >/dev/null 2>&1; then
        curl -sLO "${download_url}"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "${download_url}"
    else
        err "curl or wget is required"
    fi

    archive=$(basename "${download_url}")
    case "${archive}" in
        *.tar.gz) tar -xzf "${archive}" ;;
        *.zip) unzip -q "${archive}" ;;
        *) err "Unknown archive format: ${archive}" ;;
    esac

    mkdir -p "${BIN_DIR}"
    cp apotheke "${BIN_DIR}/apotheke"
    chmod +x "${BIN_DIR}/apotheke"

    echo ""
    echo "Installed apotheke to ${BIN_DIR}/apotheke"
    
    if ! echo ":${PATH}:" | grep -q ":${BIN_DIR}:"; then
        echo ""
        echo "Note: ${BIN_DIR} is not in your PATH."
        echo "Add this to your shell config:"
        echo "  export PATH=\"\${PATH}:${BIN_DIR}\""
    fi

    echo ""
    echo "To enable the 'a' shortcut, add to your shell config:"
    echo "  eval \"\$(apotheke init bash)\"   # for bash"
    echo "  eval \"\$(apotheke init zsh)\"    # for zsh"
    echo "  apotheke init fish | source      # for fish"
}

get_arch() {
    os=$(uname -s | tr '[:upper:]' '[:lower:]')
    arch=$(uname -m)

    case "${os}" in
        linux) os="linux" ;;
        darwin) os="darwin" ;;
        mingw* | msys* | cygwin*) os="windows" ;;
        *) err "Unsupported OS: ${os}" ;;
    esac

    case "${arch}" in
        x86_64 | amd64) arch="amd64" ;;
        aarch64 | arm64) arch="arm64" ;;
        *) err "Unsupported architecture: ${arch}" ;;
    esac

    echo "${os}_${arch}"
}

get_download_url() {
    arch="$1"
    
    if command -v curl >/dev/null 2>&1; then
        releases=$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest")
    elif command -v wget >/dev/null 2>&1; then
        releases=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest")
    else
        err "curl or wget is required"
    fi

    echo "${releases}" | grep "browser_download_url" | grep "${arch}" | head -1 | cut -d '"' -f 4
}

err() {
    echo "Error: $1" >&2
    exit 1
}

main "$@"
