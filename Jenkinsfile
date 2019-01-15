// workspaces
def goWs = '/go/src/github.com/logrhythm/case-api'

pipeline {
    agent any
    stages {
        stage('Build') {
            agent { label "golang-1.10", customWorkspace: goWs }
            steps {
                echo 'Building..'
                sh 'ls'
                sh go build
                sh 'ls'
                stash includes: 'salt-auto-update', name: 'app'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                unstash 'app'
                archiveArtifacts artifacts: 'salt-auto-update', fingerprint: true
            }
        }
    }
    post {
        always {
            echo 'Post'
        }
    }
}