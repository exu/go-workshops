# Kubernetes deployment on local machine

We'll try now put our code in working Kubernetes cluster with use of Kubernetes and Minikube on our local machine.

Kubernetes looks great but doing quick developemnt flow could be plain in the ass, if you don't have access to infrastructure of AWS or GCE with prepared pipelines to pass your code to valid infrastructure.

What if we want do develop our containers on local machine? I did'n found the out-of-the box solution but there is quite nice workaround for managing your own registry in [Sharing a local registry with minikube](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615) article.


## Getting started

0. Install kubectl and minikube
1. using docker registry runned on minikube cluster + proxy
     ([source](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615))

            ❯ kubectl create -f kubernetes/kube-registry.yaml

            # forwarding ports is temporary
            ❯ kubectl port-forward --namespace kube-system \
            $(kubectl get po -n kube-system | grep kube-registry-v0 | \
            awk '{print $1;}') 5000:5000


3. Now it's time to build our app. New docker have ability to build multi-stage builds.

    We don't need go in our system, this two step build docker will build binary in one step and put it in second step in small container without dependency.


            ❯ docker build -t localhost:5000/goapp:latest .
            ❯ docker push localhost:5000/goapp

    Our app is now ready to go in our local registry (on our cluster).

4. Now it's time to deploy! We'll use declarative method of managing kubernetes cluster with use of yml files.

    First step: create deployment (i've created file `deployment.yml` in `kubernetes` directory):

    ```yml
        apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
        kind: Deployment
        metadata:
        name: goapp-deployment
        spec:
        selector:
            matchLabels:
            app: goapp
        replicas: 2 # tells deployment to run 2 pods matching the template
        template: # create pods using pod definition in this template
            metadata:
            labels:
                app: goapp
            spec:
            containers:
            - name: goapp
                image: localhost:5000/goapp:latest
                ports:
                - containerPort: 8080
    ```

    I've used `NodePort` method for exposing

    Now if our deployment is prepared (image: from our local repository), we're ready do sync our definition with kubernetes cluster:

        ❯ kubectl create -f kubernetes/deployment.yml


## Play a little with our cluster

    our deployment is ready, we can play a little with it.

    - get pods:

            ❯ kubectl get pods -l app=goapp

            NAME                               READY     STATUS    RESTARTS   AGE
            goapp-deployment-684d96ff7-27hct   1/1       Running   0          1h
            goapp-deployment-684d96ff7-ltl7h   1/1       Running   0          1h

    - get deployments

            ❯  kubectl get deployments -l app=goapp
            NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
            goapp-deployment   2         2         2            2           1h

    - exec something on pod

            ❯ kubectl exec -it goapp-deployment-684d96ff7-27hct sh
            /app # ps aux
            PID   USER     TIME   COMMAND
                1 root       0:00 /bin/sh -c ./goapp
                5 root       0:00 ./goapp
            61 root       0:00 sh
            65 root       0:00 ps aux
            /app #

        Yeah! there is my goapp running, but **its network is exposed only inside Kubernetes cluster**. Now it's time to expose it outside cluster.

    - expose as Service (`NodePort`) - use mapping of some port on cluster node to external cluster ip.

        Now create our service definition file (`kubernetes/servicey.yml`)

        ```yml
            apiVersion: v1
            kind: Service
            metadata:
            name: goapp
            labels:
                app: goapp
            spec:
            type: NodePort
            ports:
                - port: 8080
                nodePort: 30080
            selector:
                app: goapp
        ```

        And synchronise our cluster:

            ❯ kubectl create -f kubernetes/service.yml

        After that we can check if our service is created correctly:

            ❯ kubectl get service goapp
            NAME      TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
            goapp     NodePort   10.106.164.215   <none>        8080:30080/TCP   26m

        `30080` is our port which will be visible outside cluster


        Next we need to get somehow ip address of kubernetes cluster. I'm using minikube so it's quite simple

            ❯ minikube ip
            # we assign to env variable
            ❯ IP=$(minikube ip)


        When we have external cluster IP now we can access our service on given port.

            ❯ IP=$(minikube ip)
            ❯ curl $IP:30080
            Hello World! 2018-03-19 19:15:47.543450202 +0000 UTC from goapp-deployment-684d96ff7-ltl7h

        Yeah it's working.
