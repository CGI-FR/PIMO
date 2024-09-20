FROM adrienaury/go-devcontainer:v4.0

USER root

ADD cgi_ca_root.crt /usr/local/share/ca-certificates/cgi_ca_root.crt
RUN chmod 644 /usr/local/share/ca-certificates/cgi_ca_root.crt && update-ca-certificates

ADD misc-sni-google-com.crt /usr/local/share/ca-certificates/misc-sni-google-com.crt
RUN chmod 644 /usr/local/share/ca-certificates/misc-sni-google-com.crt && update-ca-certificates

RUN apk add --update --progress --no-cache make gomplate yarn

ARG VERSION_GOLICENSE=0.2.0
ARG VERSION_MILLER=6.7.0
RUN    wget --no-check-certificate -nv -O- https://github.com/mitchellh/golicense/releases/download/v${VERSION_GOLICENSE}/golicense_${VERSION_GOLICENSE}_linux_x86_64.tar.gz | tar xz -C /usr/bin golicense \
    && wget --no-check-certificate -nv -O- https://github.com/johnkerl/miller/releases/download/v${VERSION_MILLER}/miller-${VERSION_MILLER}-linux-amd64.tar.gz | tar xz --strip-components 1 -C /usr/bin miller-${VERSION_MILLER}-linux-amd64/mlr \
    && chmod +x /usr/bin/golicense /usr/bin/mlr

ENV http_proxy ${http_proxy:-}
ENV https_proxy ${https_proxy:-}
ENV no_proxy ${no_proxy:-}

USER vscode
