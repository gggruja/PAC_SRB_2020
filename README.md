# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

How to run it:

0. If you have minikube you can delete it first:

minikube delete

rm -rf .minikube 
   
1. install minikube 

brew install minikube

2. Set up your minikube

minikube config view              
- vm-driver: hyperkit
- cpus: 4
- disk-size: 100g
- memory: 8000

3. Start your minikube

minikube start 

4. Add on ingress

minikube addons enable ingres
        
5. add in hosts file: 

minikube_ip	conference.frontend

minikube_ip	conference.backend

minikube_ip	conference.keycloak

minikube_ip	conference.grafana

minikube_ip	conference.prometheus

6. Go to ../PAC_SRB_2020/infrastructure/terraform running `install.sh` script

./install

7. Go to http://conference.frontend/ 
