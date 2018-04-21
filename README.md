# Cmpe281 Hackathon Project

## Counter Burger
A Counter Burger Application developed for placing burger orders from the Catalog. Applications covers following functionalities:

1. User functionality    - [Bruce Decker](https://github.com/Bruce-Decker)
2. Catalog functionality - [Spandana](https://github.com/spandana7)
3. Cart functionality    - [Sahil Sharma](https://github.com/Sahil12S)
4. Payment functionality - [Aaditya Deowanshi](https://github.com/iamdeowanshi)
5. Review functionality  - [Raghvendra Dixit](https://github.com/raghvendra1218)

## A Brief Description
Our Online Counter Burger application has it's front end written in **nodejs**, deployed on **Heroku** (web server). As a team of five we have developed **five GO APIs** for the functionality mentioned above which are deployed on the **Docker cloud** hosted on individual EC2 instance.  

## AFK cube implementation:
* **X axis** - Cloning - We have achieved Cloning by replicating the data on the five nodes for each of the functionality mentioned above.  
* **Y axis** - Microservices - We have achieved partioning of functionality by deploying **five microservices** deployed on Docker cloud running/hosted on their respective individual EC2 instances.
* **Z axis** - Sharding - We have achieved Sharding by adding Redis as a cache server for Cart and Payment APIs.  