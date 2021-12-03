pipeline {
    agent any
    stages {
        stage('docker-build') {
            steps {
                sh 'ls -la'
                sh 'docker build .'
            }
        }
    }
}