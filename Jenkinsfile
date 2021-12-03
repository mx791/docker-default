pipeline {
    agent any
    stages {
        stage('docker-build') {
            steps {
                sh 'ls -la'
                withCredentials([string(credentialsId: 'SECRET_WORD', variable: 'SECRET_WORD')]) {
                    sh 'echo $SECRET_WORD'
                }
                sh 'docker run hello-world'
            }
        }
    }
}