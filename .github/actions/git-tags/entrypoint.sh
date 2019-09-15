#!/bin/sh -l

TAG=$(cat /github/workspace/.release)
printf -v DATA '{"ref": "refs/tags/%s","sha": "%s"}' "$TAG" "$GITHUB_SHA"

if [ -z "$TAG" ]
then
    echo "Release not found.."
else
    echo "Tag: $TAG"
    echo "DATA: ${DATA}"

    # POST a new ref to repo via Github API
    curl -s -X POST https://api.github.com/repos/$GITHUB_REPOSITORY/git/refs \
    -H "Authorization: token $GITHUB_TOKEN" \
    -d "${DATA}"
fi