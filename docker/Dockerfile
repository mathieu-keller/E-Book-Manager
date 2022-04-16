FROM alpine:3.15.4
LABEL maintainer=MathieuKeller@gmx.de
WORKDIR /app
COPY backend/e-book-manager .
COPY frontend ./dist
RUN chmod 0777 e-book-manager
EXPOSE 8080
ENTRYPOINT ["/app/e-book-manager"]
