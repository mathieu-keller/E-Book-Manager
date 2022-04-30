FROM alpine:3.15.4
LABEL maintainer=Business@Mathieu.tech
RUN addgroup -g 1001 -S appuser && adduser -u 1001 --create-home -S appuser -G appuser
WORKDIR /home/appuser
COPY --chown=1001:1001 backend/e-book-manager .
COPY --chown=1001:1001 frontend ./bundles
USER appuser
EXPOSE 8080
ENTRYPOINT ["/home/appuser/e-book-manager"]
