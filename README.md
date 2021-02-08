# Stringcompare
Stringcompare takes 2 texts in a POST request and returns a percentage of common words.

## About

Stringcompare reponds to a POST request with curls default header which includes `Content-Type: application/x-www-form-urlencoded`
This application could be upgraded later to also support other types of requests such as JSON. 

This application compares the similarity of the text based on word matching of each. Sentences with a high percentage of matching words, regardless of order, will be considered similar. 

Ways to improve this application:

- Giving weight to consecutive matched words would provide a better indicator of similarity.
- Allow matching of conjunctions to their non-conjoined pair with weight penalty. (can't vs can not)
- Incorporate all types of post requests.

## Installation

Stringcompare can be installed as a Golang Package, executable, or as a docker container.


### Golang Package

Make sure you have Golang installed per the instructions found [here](https://golang.org/doc/install).

```bash
git clone https://github.com/psheets/stringcompare
cd ./stringcompare
go run main.go
```

### Executable

Download appropriate executable for your operating system from [here](https://github.com/psheets/stringcompare/releases/tag/V1.0).

Run executable per operating systems norm. 

### Docker

```bash
git clone https://github.com/psheets/stringcompare
cd ./stringcompare
docker build -t stringcompare .
docker run -p 8080:8080 stringcompare
```

or pull from Dockerhub

```bash
docker pull psheets88/stringcompare-app:latest
docker run -p 8080:8080 stringcompare
```

## Usage

Stringcompare will respond to a post request with a post request containing two sample texts in x-www-form-urlencoded format. 

Example using CURL
```
curl -d "sample_one=<SAMPLE ONE HERE>&sample_two=<SAMPLE TWO HERE>" http://localhost:8080
```
Example using Postman

![](https://philsheets.com/content/images/2021/02/Capture.PNG)

### Demo

You can test this application using https://stringcompare-igftb.ondigitalocean.app

```
curl -d "sample_one=Test number one.&sample_two=Test number two." https://stringcompare-igftb.ondigitalocean.app
```