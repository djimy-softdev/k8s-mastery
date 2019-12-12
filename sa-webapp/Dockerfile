FROM maven:ibmjava-alpine as builder
WORKDIR /var
COPY . .
RUN mvn install

FROM openjdk:8-jdk-alpine
COPY  --from=builder /var/target/sentiment-analysis-web-0.0.1-SNAPSHOT.jar /
EXPOSE 8080
ENTRYPOINT java -jar sentiment-analysis-web-0.0.1-SNAPSHOT.jar --sa.logic.api.url=${SA_LOGIC_API_URL}