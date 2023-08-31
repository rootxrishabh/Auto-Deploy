pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Starting build...'
                git 'https://github.com/rootxrishabh/Auto-Deploy'
                sh "docker build -t autodeployapp/autodeploy:${env.BUILD_NUMBER}.0 ."
                echo 'Build completed'
            }
        }
    stage('Push to HUB') {
        steps {
                echo 'Pushing to docker hub ...'
                sh "docker login -u autodeployapp -p password123"
                sh "docker push autodeployapp/autodeploy:${env.BUILD_NUMBER}.0"
                echo 'Push completed' 
            }
        }
    stage('Version') {
        steps {
                echo 'Starting update...'
                git 'https://github.com/rootxrishabh/autodeploy-yaml.git'
                sh "sed -i 's|image: autodeployapp/autodeploy:.*|image: autodeployapp/autodeploy:${env.BUILD_NUMBER}.0|' deployment.yaml"
                sh "cat deployment.yaml"
                sh "git config user.name 'rootxrishabh'"
                sh "git config user.email 'risrock02@gmail.com'"
                sh "git add ."
                sh "git commit -m 'Jenkins build version ${env.BUILD_NUMBER}.0'"
                sh "git config --local credential.helper '!f() { echo username=rootxrishabh; echo password=PAT; }; f'"
                sh "git push origin master"
                echo 'Update completed'   
            }
        }
    }
}
