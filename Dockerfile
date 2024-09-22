# Specifies a parent image
FROM golang:1.19.2-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /devops-test
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build -o /godocker
 
# Tells Docker which network port your container listens on
EXPOSE 5000

# Installs Kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin
 
# Specifies the executable command that runs when the container starts
CMD [ “/godocker” ]
