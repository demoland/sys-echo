FROM alpine:latest

RUN apk add --no-cache go git

ARG RELEASE
ARG TEST_MESSAGE="This is a test Message"
RUN wget https://github.com/demoland/sys-echo/releases/download/${RELEASE}/sys-echo-ubuntu-x86-64 -O /usr/local/bin/sys-echo \
    && chmod +x /usr/local/bin/sys-echo

EXPOSE 8080
CMD ["/usr/local/bin/sys-echo",  ${TEST_MESSAGE} ]
