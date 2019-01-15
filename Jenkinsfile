// workspaces
def goWs = '/go/src/github.com/logrhythm/salt-auto-update'

pipeline {
    agent { node label: "golang-1.10", customWorkspace: goWs }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'go get gopkg.in/yaml.v2'
                sh 'go build'
                chmod +x salt-auto-update
                stash includes: 'salt-auto-update', name: 'app'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                unstash 'app'
                archiveArtifacts artifacts: 'salt-auto-update', fingerprint: true
                sh "curl --insecure --user ${USERNAME}:${SSH_PASSWORD} -T salt-auto-update sftp://${REMOTE_SERVER}/~/"
            }
        }
    }
    post {
        always {
            echo 'Post'
        }
    }
}