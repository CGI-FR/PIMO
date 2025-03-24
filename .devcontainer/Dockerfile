FROM adrienaury/go-devcontainer:v6.0

USER root

ADD --chmod=644 *.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

RUN apk add --update --progress --no-cache make gomplate yarn

ARG VERSION_GOLICENSE=0.2.0
ARG VERSION_MILLER=6.13.0
RUN    wget -nv -O- https://github.com/mitchellh/golicense/releases/download/v${VERSION_GOLICENSE}/golicense_${VERSION_GOLICENSE}_linux_x86_64.tar.gz | tar xz -C /usr/bin golicense \
    && wget -nv -O- https://github.com/johnkerl/miller/releases/download/v${VERSION_MILLER}/miller-${VERSION_MILLER}-linux-amd64.tar.gz | tar xz --strip-components 1 -C /usr/bin miller-${VERSION_MILLER}-linux-amd64/mlr \
    && chmod +x /usr/bin/golicense /usr/bin/mlr

ENV http_proxy ${http_proxy:-}
ENV https_proxy ${https_proxy:-}
ENV no_proxy ${no_proxy:-}

USER vscode
