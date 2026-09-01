FROM ubuntu:latest
LABEL authors="amirb"

ENTRYPOINT ["top", "-b"]