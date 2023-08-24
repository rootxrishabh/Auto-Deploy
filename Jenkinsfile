pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building the go application\'s docker image ...'
                git 'https://github.com/rootxrishabh/AutoDeploy.git'
                sh "docker build -t rootxrishabh/Library-API-autodeploy:${env.BUILD_NUMBER}.0 ."
                echo 'Your image has been built'
            }
        }
        stage('Push to HUB') {
            steps {
                echo 'Pushing to docker hub ...'
                sh "docker login -u rootxrishabh -p password"
                sh "docker push rootxrishabh/AutoDeploy:${env.BUILD_NUMBER}.0"
                echo 'Push completed' 
            }
        }
        stage('Update Deployment') {
            steps {
                echo 'Updating the deploying script ...'
                git 'https://github.com/rootxrishabh/AutoDeploy-yaml.git'
                sh "sed -i 's|image: rootxrishabh/AutoDeploy:.*|image: rootxrishabh/AutoDeploy:${env.BUILD_NUMBER}.0|' deployment.yaml"
                sh "cat deployment.yaml"
                sh "git config user.name 'rootxrishabh'"
                sh "git config user.email 'risrock02@gmail.com'"
                sh "git add ."
                sh "git commit -m 'Jenkins build version ${env.BUILD_NUMBER}.0'"
                sh "git config --local credential.helper '!f() { echo username=rootxrishabh; echo password=PAT; }; f'"
                sh "git push origin master"
                echo 'Script Updated'   
            }
        }
    }
}