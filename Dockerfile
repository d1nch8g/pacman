# 2023 FMNX team.
# Use of this code is governed by GNU General Public License.
# Additional information can be found on official web page: https://fmnx.su/
# Contact email: help@fmnx.su

FROM docker.io/archlinux/archlinux:base-devel

# This dockerfile contains sudo environment for library testing.

LABEL maintainer="dancheg <dancheg@fmnx.su>"

RUN pacman -Sy --noconfirm --noprogressbar go
RUN useradd --system --create-home usr
RUN echo "usr ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/usr
USER usr
WORKDIR /home/usr

COPY go.mod /home/usr/
COPY go.sum /home/usr/
RUN go mod download

COPY . /home/usr/
