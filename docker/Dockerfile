FROM alpine:3.15.4
LABEL maintainer=Business@Mathieu.tech
RUN addgroup -g 1001 -S appuser && adduser -u 1001 -S appuser -G appuser
WORKDIR /home/appuser
COPY e-book-manager .
COPY frontend ./bundles
RUN mkdir "upload" && \
    chown appuser:appuser -R . && \
    chmod 775 -R .
USER appuser

EXPOSE 8080
ENTRYPOINT ["/home/appuser/e-book-manager"]
