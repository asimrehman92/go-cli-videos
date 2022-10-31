FROM golang:1.15-alpine as dev 

WORKDIR /work

# WORKDIR /videos
# COPY ./videos/* /videos/
# RUN go build -o videos

# FROM golang:1.17-alpine as build

# FROM alpine as runtime 
# COPY --from=build /videos/videos /usr/local/bin/videos
# COPY ./videos/videos.json  /
# COPY run.sh /
# RUN chmod +x /run.sh
# ENTRYPOINT [ "./run.sh" ]
