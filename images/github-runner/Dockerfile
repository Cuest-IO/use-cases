FROM golang:1.18.4 AS build

WORKDIR /app   

COPY main.go go.mod go.sum ./

RUN go mod download

# build runner
RUN go build -o runner

FROM ubuntu:18.04

RUN mkdir -p actions-runner 

WORKDIR /actions-runner

COPY --from=build /app/runner .

RUN apt-get update \
    && apt-get install -y sudo curl

ARG USERNAME=github
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

CMD [ "./runner" ]