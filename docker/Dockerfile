FROM alpine:3.18.3
RUN addgroup -g 1001 -S appuser && adduser -u 1001 -S appuser -G appuser
WORKDIR /home/appuser
COPY src/e-book-manager .
COPY sql ./sql
RUN mkdir "upload" && \
    chown appuser:appuser -R . && \
    chmod 775 -R .
USER appuser

EXPOSE 8080
ENTRYPOINT ["/home/appuser/e-book-manager"]
