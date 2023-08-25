# Auto-Deploy

This is an API which helps manage a library system. Deploy this application using a pipeline with the following commands below

# Step 1
Build the docker image locally with the provided file.

# Step 2
The below will create a jenkins image which will run on docker, it is also enabled to talk to host docker daemon.
```
FROM jenkins/jenkins:lts
USER root
RUN apt-get update
RUN curl -sSL https://get.docker.com/ | sh
```
And run this image using the command
```
sudo docker run -p 8080:8080 -p 50000:50000 -v /var/run/docker.sock:/var/run/docker.sock go/jenkins:latest
```
go/jenkins:latest is the name of the built image in our case.

# Step 3
Setup Jenkins dashboard

# Step 4
Create a new Pipeline project in Jenkins and enable GitSCM polling to integrate with GitHub webhooks.
Add the Jenkinsfile to the script to create the pipeline steps and initialize the pipeline.

# Step 5
Expose your localhost port 8080 which is running Jenkins to the internet using ngrok.
You will have to create a ngrok account which will provide you with a key to use ngrok
'''
ngrok http 8080
'''
Verify the link provided if it is now accessible.

# Step 6
Go to your GitHub project of your application and in settings add a new webhook with the ngrok url and add '/github-webhook/' at the end of the url, enable webhook to run on pushes, PRs and commits. Jenkins recieves webhook request at this particular URL only

# Step 7
Now startup minikube and install ArgoCD using the official documentation and create a new app that links to your other repo which will include all k8s config yaml files for GitOps.

# Step 8
Simply expose your service using the minikube official documentation.
'''
minikube service pipeline-service --url
'''

## You now have a fully functional CI/CD pipeline which uses docker, jenkins, argo CD, kubernetes, github/git.
