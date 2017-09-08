# Base system is alpine linux for small size
FROM alpine:latest

# Install build dependencies
RUN apk -U add bash openssl ca-certificates go git musl-dev

ENV GOPATH /go
ENV LILYPAD_INSTALL_DIR /opt/lilypad
ENV LILYPAD_BUILDDIR $GOPATH/src/github.com/LilyPad/GoLilyPad
ENV LILYPAD_CONNECT_BIN connect
ENV LILYPAD_PROXY_BIN proxy

# Copy source
COPY . $LILYPAD_BUILDDIR

# Build and install lilypad connect
COPY .docker/install.sh /app/install.sh
RUN chmod 755 /app/install.sh
RUN /app/install.sh
# Cleanup
RUN rm /app/install.sh

# Add initscript
ADD .docker/init.sh /app/init.sh
# Fix permissions
RUN chmod 755 /app/init.sh

EXPOSE 5091 25565
VOLUME ["/data"]

WORKDIR /data

ENTRYPOINT ["/app/init.sh"]
CMD ["app:help"]

