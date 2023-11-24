FROM golang:latest as builder

EXPOSE 8080

#Make a directory
RUN mkdir /app
#Add source code to the directory
ADD ./application /app
#Set the working directory
WORKDIR /app
#Build the application
RUN go build -o main .


FROM golang:latest as production

#Make a directory
RUN mkdir /app

#Create a user
RUN useradd --user-group --system --create-home --no-log-init app
#Set the default user
USER app
#Copy the binary file from builder to production
COPY --from=builder --chown=app:app /app/main /app/main

WORKDIR /app

#Expose the port
CMD ["./main"]

