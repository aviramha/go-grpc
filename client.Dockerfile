FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS build-env

WORKDIR /app
COPY ./client /app

ARG TARGETARCH
RUN GOARCH=$TARGETARCH go build -o /main main.go

FROM  gcr.io/distroless/static-debian11

COPY --from=build-env /main /main

CMD [ "/main" ]
