# Copyright (C) 2021 CGI France
#
# This file is part of PIMO.
#
# PIMO is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# PIMO is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with PIMO.  If not, see <http:#www.gnu.org/licenses/>.

FROM gcr.io/distroless/base

COPY bin/pimo /usr/bin/pimo

WORKDIR /home/pimo

ENTRYPOINT [ "/usr/bin/pimo" ]

ARG BUILD_DATE
ARG VERSION
ARG REVISION

# https://github.com/opencontainers/image-spec/blob/master/annotations.md
LABEL org.opencontainers.image.created       "${BUILD_DATE}"
LABEL org.opencontainers.image.authors       "CGI Lino <lino.fr@cgi.com>"
LABEL org.opencontainers.image.url           "https://hub.docker.com/r/cgifr/pimo"
LABEL org.opencontainers.image.documentation "https://github.com/CGI-FR/PIMO/blob/main/README.md"
LABEL org.opencontainers.image.source        "https://github.com/CGI-FR/PIMO.git"
LABEL org.opencontainers.image.version       "${VERSION}"
LABEL org.opencontainers.image.revision      "${REVISION}"
LABEL org.opencontainers.image.vendor        "CGI France"
LABEL org.opencontainers.image.licenses      "GPL-3.0-only"
LABEL org.opencontainers.image.ref.name      "cgi-pimo"
LABEL org.opencontainers.image.title         "CGI PIMO"
LABEL org.opencontainers.image.description   "PIMO Private Input, Masked Output is a tool for data masking. It can mask data from a JSONline stream and return another JSONline stream thanks to a masking configuration contained in a yaml file."
