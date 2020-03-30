FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/backend/graduate_project
WORKDIR /go/src/backend/graduate_project

RUN go build
RUN export GOPROXY=
CMD if [ ${APP_ENV} = production ]; \
        then \
        app; \
        else \
        go get github.com/pilu/fresh && \
        fresh; \
        fi

EXPOSE 18080