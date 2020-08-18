# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

INFRASTRUCTURE

There are terraform script which set up whole environment!

If you want to set up environment ona a push a button you can execute `install.sh` script

Note: Before you need to build two images for frontend and backend service by entering:

         docker build -f Dockerfile -t backend .
         docker build -f Dockerfile -t frontend .

From directory of backend and frontend where the Dockerfile is.
