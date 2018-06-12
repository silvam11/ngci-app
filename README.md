# ngci-app
Below are the steps to deploy this application.

## Prerequisistes
The steps described here assume that:

1. istioctl version 0.8 is installed and available on the PATH environment variable.
2. Kubernetes and docker are installed.
3. Git is installed.
4. All the scripts mentioned on these instructions have to the executables.

## Deploy code form git repo
To get access to this code, please run the following command

```
git clone https://github.com/silvam11/ngci-app
```

The previously command creates a folder named ngci-app under the current directory. Then, run the command:

```
cd ngci-app
```

## Create docker images

This project provides to docker images which are created running the script:

```
./create-docker-images.sh
```

## Deploy the demo application

To deploy the demo application, run the command below.

```
./setup.sh
```
## Clean the environment

After using this application, the kubernetes environment can be cleaned using the following command:

```
./cleanup.sh
```