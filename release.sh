#!/bin/bash
LAST_TAG=$(gh release view --json tagName --jq '.tagName')
NEXT_TAG=$(echo "${LAST_TAG}" | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}')
echo "Bumping from " "${LAST_TAG}" " -> " "${NEXT_TAG}"

confirm() {
    # call with a prompt string or use a default
    read -r -p "${1:-Are you sure? [y/N]} " response
    case "$response" in
        [yY][eE][sS]|[yY])
            true
            ;;
        *)
            false
            ;;
    esac
}

if confirm "Push tags and create GitHub release? [y/N] "; then
  # tag root module and sub-module at the SAME commit
  git tag -a "${NEXT_TAG}" -m "Release chalk-go ${NEXT_TAG}"
  git tag -a "gen/${NEXT_TAG}" -m "Release chalk-go/gen ${NEXT_TAG}"

  # push both tags
  git push origin "${NEXT_TAG}" "gen/${NEXT_TAG}"

  # create GitHub release for the root tag
  gh release create "${NEXT_TAG}" --generate-notes
fi