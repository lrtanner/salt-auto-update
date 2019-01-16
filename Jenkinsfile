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
                sh 'chmod +x salt-auto-update'
                sh 'tar -cvf salt-auto-update.tar config.yaml salt-script.sh salt-auto-update'
                stash includes: 'salt-auto-update.tar', name: 'app'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                set +x
                unstash 'app'
                archiveArtifacts artifacts: 'salt-auto-update.tar', fingerprint: true
                sh "curl --insecure --user ${USERNAME}:${SSH_PASSWORD} -T salt-auto-update.tar sftp://${REMOTE_SERVER}/~/"
                sh "sshpass -p ${SSH_PASSWORD} ssh -t ${USERNAME}@${REMOTE_SERVER} 'sh /home/logrhythm/temp/salt-script.sh'"
            }
        }
    }
    post {
        always {
            echo 'Post'
        }
    }
}