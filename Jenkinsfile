pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh "echo 'Build number ${BUILD_ID}'"
                sh "docker build -t my-app:${BUILD_ID} ."
            }
        }
    }
}
