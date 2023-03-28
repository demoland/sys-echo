FROM alpine:latest

RUN apk add --no-cache go git

ARG SYS_ECHO_VERSION
ARG TEST_MESSAGE="This is a test Message"
RUN wget https://github.com/demoland/repo/releases/download/v${SYS_ECHO_VERSION}/sys-echo -O /usr/local/bin/sys-echo \
    && chmod +x /usr/local/bin/sys-echo

EXPOSE 8080
CMD ["/usr/local/bin/sys-echo",  ${TEST_MESSAGE} ]
