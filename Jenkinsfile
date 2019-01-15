// workspaces
def goWs = '/go/src/github.com/logrhythm/salt-auto-update'

pipeline {
    agent { node label: "golang-1.10", customWorkspace: goWs }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'pwd'
                sh 'ls'
                sh 'go get gopkg.in/yaml.v2'
                sh 'go build'
                sh 'ls'
                stash includes: 'salt-auto-update', name: 'app'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                unstash 'app'
                archiveArtifacts artifacts: 'salt-auto-update', fingerprint: true
                script {
                  def remote = [:]
                  remote.name = 'saltmaster'
                  remote.host = "${REMOTE_SERVER}"
                  remote.user = "${USERNAME}"
                  remote.password = "${SSH_PASSWORD}"
                  remote.allowAnyHosts = true
                  stage('Remote SSH') {
                    sshPut remote: remote, from: 'salt-auto-update', into: '~'
                  }
                }
            }
        }
    }
    post {
        always {
            echo 'Post'
        }
    }
}