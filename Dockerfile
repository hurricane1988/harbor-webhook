FROM ubuntu:latest
LABEL authors="hurricane"

ENTRYPOINT ["top", "-b"]
