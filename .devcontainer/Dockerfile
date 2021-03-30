FROM qmcgaw/godevcontainer:debian

USER root

ADD cgi_ca_root.crt /usr/local/share/ca-certificates/cgi_ca_root.crt
RUN chmod 644 /usr/local/share/ca-certificates/cgi_ca_root.crt && update-ca-certificates

RUN echo "Acquire::http::Proxy \"${http_proxy:-false}\";" > /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::https::Proxy \"${https_proxy:-false}\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::http::Pipeline-Depth \"0\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::http::No-Cache \"True\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::http::BrokenProxy \"True\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::https::Pipeline-Depth \"0\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::https::No-Cache \"True\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::https::BrokenProxy \"True\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    echo "Acquire::Check-Date \"False\";" >> /etc/apt/apt.conf.d/proxy.conf && \
    apt-get update && \
    apt-get install -y --no-install-recommends make wget jq miller && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -r /var/cache/* /var/lib/apt/lists/*

ENV http_proxy ${http_proxy:-}
ENV https_proxy ${https_proxy:-}
ENV no_proxy ${no_proxy:-}

RUN go get -v github.com/ramya-rao-a/go-outline &&\
    rm -rf /go/pkg/* /go/src/* /root/.cache/go-build

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /bin -d ${GOLANGCI_LINT_VERSION}
RUN wget -O- -nv https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
RUN wget -O- -nv https://install.goreleaser.com/github.com/caarlos0/svu.sh | bash -s -- -b /usr/local/bin

ADD https://github.com/ovh/venom/releases/download/v0.27.0/venom.linux-amd64 /bin/venom

RUN chmod +x /bin/venom
