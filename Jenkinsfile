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
                  remote.host = '10.1.23.103'
                  remote.user = 'logrhythm'
                  remote.password = 'logrhythm!1'
                  remote.allowAnyHosts = true
                  stage('Remote SSH') {
                    writeFile file: 'abc.sh', text: 'ls -lrt'
                    sshPut remote: remote, from: 'abc.sh', into: '.'
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