FROM golang:1.14.12-alpine3.12 as build-env

ARG BUILD_DIR=/build
ADD . $BUILD_DIR

RUN cd ${BUILD_DIR} && go build -mod=vendor -o bush


FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /build/bush /app
COPY --from=build-env /build/conf /app/conf
CMD [ "./bush"]