FROM golang:1.24

RUN apt-get update && \
    apt-get install -y jq

LABEL "com.github.actions.name"="Lint"
LABEL "com.github.actions.description"="Lint go gocode"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/gabeduke/level"
LABEL "homepage"="http://github.com/actions"
LABEL "maintainer"="Octocat <octocat@github.com>"

ENV GO111MODULE=on

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]