FROM alpine:3.14

# Install dependencies
RUN apk add --no-cache \
    wget \
    jq

# Download sys-echo binary
ARG SYS_ECHO_VERSION=v1.0.0
RUN wget -O /usr/local/bin/sys-echo "https://github.com/demoland/sys-echo/releases/download/${SYS_ECHO_VERSION}/sys-echo"

# Make sys-echo binary executable
RUN chmod +x /usr/local/bin/sys-echo

# Expose port 8080
EXPOSE 8080

# Start sys-echo
ENTRYPOINT ["/usr/local/bin/sys-echo", "-msg"]
