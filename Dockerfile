# First stage: build the executable.
FROM golang:1.12.1-alpine AS builder

# Uncomment to trigger a failed build due to failing tests
# ENV FAILED=true

# Create the user and group files that will be used in the running container to
# run the process an unprivileged user.
RUN mkdir /user && \
		echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
		echo 'nobody:x:65534:' > /user/group

# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Run all tests
RUN CGO_ENABLED=0 go test ./...

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app ./cmd/app/main.go

# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app
COPY --from=builder /src/web /web

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

ENTRYPOINT ["/app"]
