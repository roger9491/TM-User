pipeline {
    agent any
    tools {
        go 'Go'
    }
    
    environment{
        PROJECT = 'tm-user'
    }

    stages {
        stage('Build Test image'){
            steps{
                sh "docker build --target test -t ${env.PROJECT} ."
            }
        }
        stage('Test'){
            steps{
                sh "docker run ${env.PROJECT}"
                sh 'docker system prune -a'
            }
        }
        stage('Build Deployment image'){
            steps{
                echo 'build image'
                sh "docker build . -t ${env.PROJECT}"
            }
        }
        stage('Push image'){
            steps{
                echo 'docker login'
                withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
                    sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                    echo 'docker login successful'
                    sh "docker tag ${env.PROJECT} ${env.dockerHubUser}/${env.PROJECT}"
                    sh "docker push ${env.dockerHubUser}/${env.PROJECT}"
                }
            }
        }
        
    }
}