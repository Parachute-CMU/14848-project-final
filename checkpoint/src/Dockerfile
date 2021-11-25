FROM openjdk:8-jdk-alpine
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
RUN javac Driver.java
CMD ["java", "Driver"]