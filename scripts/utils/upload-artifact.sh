#!/usr/bin/env bash
set -exuo pipefail

BINARY_NAME="mercury"
DIST_DIR="dist"
FILENAME="dist.zip"

mapfile -d '' files < <(
  find "$DIST_DIR" -regextype posix-extended -type f \
    -regex ".*/[^/]*(amd64|arm64)[^/]*/${BINARY_NAME}(\\.exe)?$" -print0
)

if [[ ${#files[@]} -eq 0 ]]; then
  echo -e "\033[31mNo binaries found for packaging.\033[0m"
  exit 1
fi

rm -f "${DIST_DIR}/${FILENAME}"

while IFS= read -r -d '' dir; do
  printf "Remove the quarantine attribute before running the executable:\n\nxattr -d com.apple.quarantine %s\n" \
    "$BINARY_NAME" >"${dir}/README.txt"
done < <(find "$DIST_DIR" -type d -name '*macos*' -print0)

relative_files=()
for file in "${files[@]}"; do
  relative_files+=("${file#"${DIST_DIR}"/}")
done

(cd "$DIST_DIR" && zip -r "$FILENAME" "${relative_files[@]}")

RESPONSE=$(curl -X POST "$URL?filename=$FILENAME" \
  -H "Authorization: Bearer $AUTH" \
  -H "Content-Type: application/json")

SIGNED_URL=$(echo "$RESPONSE" | jq -r '.url')

if [[ "$SIGNED_URL" == "null" ]]; then
  echo -e "\033[31mFailed to get signed URL.\033[0m"
  exit 1
fi

UPLOAD_RESPONSE=$(curl -v -X PUT \
  -H "Content-Type: application/zip" \
  --data-binary "@${DIST_DIR}/${FILENAME}" "$SIGNED_URL" 2>&1)

if echo "$UPLOAD_RESPONSE" | grep -q "HTTP/[0-9.]* 200"; then
  echo -e "\033[32mUploaded build to Stainless storage.\033[0m"
  echo -e "\033[32mInstallation: Download and unzip: 'https://pkg.stainless.com/s/mercury-cli/$SHA/$FILENAME'. On macOS, run `xattr -d com.apple.quarantine {executable name}.`\033[0m"
else
  echo -e "\033[31mFailed to upload artifact.\033[0m"
  exit 1
fi
