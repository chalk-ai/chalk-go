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
confirm "Submit to GitHub (Y/n)?" && echo "Submitting..." && gh release create "${NEXT_TAG}" --generate-notes
sed -i '' -e "s/chalk-go.*/chalk-go ${NEXT_TAG}/g" ../chalk-private/go-api-server/go.mod ../chalk-private/inttest/go.mod
cd ../chalk-private/go-api-server && go mod tidy && cd ../inttest && go mod tidy
