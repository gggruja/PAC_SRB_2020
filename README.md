# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

How to run it:

* If you have minikube you can delete it first:

        minikube delete

        rm -rf .minikube 
   
* Install minikube: 

        brew install minikube

* Set up your minikube as below:

        command: minikube config set XXXXX
        
        minikube config view              
        - vm-driver: hyperkit
        - cpus: 4
        - disk-size: 100g
        - memory: 8000

* Start your minikube

        minikube start 

* Add on ingress:

        minikube addons enable ingress
        
* Add in hosts file: 

        <minikube ip>	conference conference.backend conference.keycloak conference.grafana conference.prometheus

* Add eval for images

         eval $(minikube docker-env) 

* Go to PATH ../frontend and ../backend run image creation:

        docker build -f Dockerfile -t backend .
        docker build -f Dockerfile -t frontend .

* Go to PATH ../PAC_SRB_2020/infrastructure/terraform run `install.sh` script:

        ./install

* Open URL in browser: 

        http://conference/ 
 
* User/pass from user in keycloak: 

        ggrujic/ggrujic     
