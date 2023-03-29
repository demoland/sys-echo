FROM alpine:3.14

# Install dependencies
RUN apk add --no-cache \
    wget \
    jq

# Copy sys-echo binary from local file system to Docker image
COPY sys-echo-linux /usr/local/bin/sys-echo

# Make sys-echo binary executable
RUN chmod +x /usr/local/bin/sys-echo

# Expose port 8080
EXPOSE 8080

# Start sys-echo
ENTRYPOINT ["/usr/local/bin/sys-echo", "-msg"]
