FROM postgres:16-alpine3.20

HEALTHCHECK \
    --start-period=10s \
    --start-interval=1s \
    --interval=10s \
    --timeout=5s \
    --retries=3 \
    CMD pg_isready
