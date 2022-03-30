FROM adrienaury/go-devcontainer:v0.6

USER root

ADD cgi_ca_root.crt /usr/local/share/ca-certificates/cgi_ca_root.crt
RUN chmod 644 /usr/local/share/ca-certificates/cgi_ca_root.crt && update-ca-certificates

RUN apk add --update --progress --no-cache make gomplate yarn

ARG VERSION_GOLICENSE=0.2.0
ARG VERSION_MILLER=5.10.2
RUN    wget -nv -O- https://github.com/mitchellh/golicense/releases/download/v${VERSION_GOLICENSE}/golicense_${VERSION_GOLICENSE}_linux_x86_64.tar.gz | tar xz -C /usr/bin golicense \
    && wget -nv -O/usr/bin/mlr https://github.com/johnkerl/miller/releases/download/v${VERSION_MILLER}/mlr.linux.x86_64 \
    && chmod +x /usr/bin/golicense /usr/bin/mlr

ENV http_proxy ${http_proxy:-}
ENV https_proxy ${https_proxy:-}
ENV no_proxy ${no_proxy:-}

USER vscode
