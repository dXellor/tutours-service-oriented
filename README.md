# Tours marketplace system designed as a service-oriented architecture

## About
This repository contains the source code of the project which is the part of the **Service Oriented Architectures** course. The goal of this project is to migrate existing [Tutours](https://github.com/RA2020PSW8/tourism-api) project from **Modular monolith architecture** to the **Service-oriented architecture** by extracting some of the core modules and adding some new modules as services. By utilizing **Docker** our team was able to pack different parts of the system into containers and implement logging and monitoring throughout the system. Communication between services is optimized by using **GRPC** and the **HTTP** requests from the frontend side are translated into the **GRPC** calls via the **GRPC Gateway** implemented in **GO**.

## Architecture diagram
![Architecture diagram](/docs/architecture.png)

## Technologies
* **Languages/Frameworks:** GO, Python, DotNetCore, Angular
* **Database management systems:** PostgreSql, Neo4J, MongoDB
* **External services:** Docker, FluentBit, Prometheus, Grafana, Loki

## Running project
### Requirements
You must have software from this list installed on your system to be able to run the project
* Docker
* Angular 16.x
* NodeExporter

### Running project
After you setup all the required software follow these steps to run the application
* enter the [/src](/src/) and run ``docker-compose -f docker-compose-grpc.yml up`` OR ``docker-compose up`` to start all the services 
* enter the [/src/ui/Explorer](/src/ui/Explorer/) and run ``ng serve``

## Authors

* [Nikola Simić](https://github.com/dXellor)
* [Anastasija Radić](https://github.com/anastasijaradic)
* [Jelena Vujović](https://github.com/zanyaIO)
* [Srđan Petronijević](https://github.com/srdjanpetronijevic)