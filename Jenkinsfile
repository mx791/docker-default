pipeline {
    agent any
    stages {
        stage('docker-build') {
            steps {
                sh 'ls -la'
                withCredentials([string(credentialsId: 'SECRET_WORD', variable: 'SECRET_WORD')]) {
                    sh 'echo $SECRET_WORD'
                }
                withCredentials([string(credentialsId: 'TEST', variable: 'TEST')]) {
                    sh 'echo $TEST'
                }
                sh 'docker run hello-world'
            }
        }
    }
}