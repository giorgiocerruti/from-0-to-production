FROM golang:latest as builder

RUN mkdir /application
WORKDIR /application

COPY ./application .

RUN go build -o app .

FROM golang:latest as prod

RUN mkdir /application

RUN useradd --user-group --system --create-home --no-log-init app
USER app

COPY --from=builder --chown=app:app /application/app /application/app

WORKDIR /application

CMD [ "./app" ]
