#!/bin/sh -l

TAG=$(cat /github/workspace/.release)

if [ -z "$TAG" ]
then
    echo "Release not found.."
else
    echo "Tag: $TAG"

    # POST a new ref to repo via Github API
    curl -s -X POST https://api.github.com/repos/$GITHUB_REPOSITORY/git/refs \
    -H "Authorization: token $GITHUB_TOKEN" \
    -d '{"ref": "refs/tags/'"$TAG"'","sha": '"$GITHUB_SHA"'"}'
fi