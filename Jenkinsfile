pipeline {
    agent any
    stages {
        stage('docker-build') {
            steps {
                sh 'ls -la'
                withCredentials([string(credentialsId: '	SOME_TEXT', variable: '	SOME_TEXT')]) {
                    sh 'echo $SOME_TEXT'
                }
                sh 'docker run hello-world'
            }
        }
    }
}