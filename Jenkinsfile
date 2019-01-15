pipeline {
    agent any
    stages {
        stage('Build') {
            agent { label "nodejs" }
            steps {
                echo 'Building..'
                sh 'ls'
                sh (
                        script: './echo.sh',
                        returnStatus: true
                )
                sh 'ls'
                stash includes: '**/*.jar', name: 'jars'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
    post {
        always {
            unstash 'jars'
            archiveArtifacts artifacts: '**/*.jar', fingerprint: true
        }
    }
}