FROM alpine:3.15.4
LABEL maintainer=Business@Mathieu.tech
RUN addgroup -g 1001 -S appuser && adduser -u 1001 -S appuser -G appuser
WORKDIR /home/appuser
COPY --chown=appuser:appuser backend/e-book-manager .
COPY --chown=appuser:appuser frontend ./bundles
RUN chmod 0500 e-book-manager && chmod 0600 -R bundles
USER appuser:appuser

EXPOSE 8080
ENTRYPOINT ["/home/appuser/e-book-manager"]
