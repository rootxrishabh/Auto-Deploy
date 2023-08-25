pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building the go application\'s docker image ...'
                git 'https://github.com/rootxrishabh/Auto-Deploy'
                sh "docker build -t autodeploy/autodeploy:1.0"
                echo 'Your image has been built'
            }
        }
        stage('Push to HUB') {
            steps {
                echo 'Pushing to docker hub ...'
                sh "docker login -u autodeploy -p password123"
                sh "docker push autodeploy/autodeploy:1.0"
                echo 'Push completed' 
            }
        }
        stage('Update Deployment') {
            steps {
                echo 'Updating the deploying script ...'
                git 'https://github.com/rootxrishabh/AutoDeploy-yaml.git'
                sh "sed -i 's|image: autodeploy/autodeploy:.*|image: autodeploy/autodeploy:${env.BUILD_NUMBER}.0|' deployment.yaml"
                sh "cat deployment.yaml"
                sh "git config user.name 'rootxrishabh'"
                sh "git config user.email 'risrock02@gmail.com'"
                sh "git add ."
                sh "git commit -m 'Jenkins build version ${env.BUILD_NUMBER}.0'"
                sh "git config --local credential.helper '!f() { echo username=rootxrishabh; echo password=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; }; f'"
                sh "git push origin master"
                echo 'Script Updated'   
            }
        }
    }
}
