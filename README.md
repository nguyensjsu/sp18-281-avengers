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

## Network Partition :
* We have acheived this by having a partition in the cluster, hit the endpoint using the Postman and read the stale data for a particular feature

## Creativity:
* We have deployed our GO API on the docker cloud by building the docker images and its respective compose file, thus making the process of deployment smoother
* We have a cluster of five Riak nodes which has an Elastic Load balancer to distribute the traffic among the nodes.  

## Agile Values
* [Bruce      -- Feedback](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Feedback_Bruce_Decker.md) 
* [Spandana   -- Courage](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Courage_Spandana_Padala.md)  
* [Raghvendra -- Communication](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Communication_Raghvendra_Dixit.md)  
* [Sahil 	  -- Commitment](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Commitment_Sahil_Sharma.md)  
* [Aaditya 	  -- Respect](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Respect_Aaditya_Deowanshi.md)

## Weekly Meeting Minutes
* Minutes of Meeting (MoM)](https://github.com/nguyensjsu/team281-avengers/blob/master/Project_documentation/Minutes_Meeting.md)
