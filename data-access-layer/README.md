# Data Access Layer

The idea behind this design pattern is to decouple the control plane and the data plane so that the logic that handles the operations on database is completely abstracted. Thus, giving the engineering team the flexibility to adopt different databases at any given point of time without having to touch the control layer of the software. 

This design pattern also gives a lot of benefits to test the data layer independently without the involvement or interference of the control layer. Thus adding to the quality of the software.

## Example

In this example implementation, a simple go application is coded that performs read and write on "Students" table in a "College" database. The idea is to implement it with DAL (Data Access Layer) Design Pattern so that we can use any DB anytime without having to touch the control plane.

The example uses PostgreSQL and MongoDB for the purpose.

## Usage

1. First deploy the PSQL and MongoDB temporary deployments (these are not attached with volumes hence the data won't be persisted) & services

    ```bash
    k apply -f manifests/
    ```

2. Then port-forward the services so that be accessible from outside the K8s local cluster (note that, you will have to run these below commands in different terminals or run them in the background)

    ```bash
    k port-forward svc/postgres 5432:5432
    k port-forward svc/mongodb 27017:27017
    ```

3. Once, the pods are running, services are up and ports are forwarded, go program is good to be run.

    ```bash
    # to run the example with PostgreSQL DB
    go run cmd/dal.go --db-provider=psql

    # to run the example with MongoDB
    go run cmd/dal.go --db-provider=mongodb
    ```