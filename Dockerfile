FROM  golang:1.22.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o /api main.go


FROM build-stage AS run-test-stage 
RUN go test ./Test 

FROM scratch AS build-release-stage 

WORKDIR /

COPY --from=build-stage /api /api 

ENV API_PORT=8080
ENV MONGODB_URI="mongodb+srv://admin:root@cluster0.w4lxs.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
ENV SECRET_KEY="your-secret-key"

EXPOSE 8080 

ENTRYPOINT [ "/api" ]
