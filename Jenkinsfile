pipeline {
  agent {
    label 'linux-default'
  }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        GOBIN = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/bin"
        PATH = "${PATH}:${GOBIN}"
    }
    tools {
        go 'go-1.17.7'
    }

  stages {
    stage('Cleaning workspace'){
        steps {
            sh 'rm -rf $JENKINS_HOME/workspace/$JOB_NAME/content'
            sh 'rm -rf $JENKINS_HOME/workspace/$JOB_NAME/templates'
            sh 'rm -rf $JENKINS_HOME/workspace/$JOB_NAME/*.nupkg'
            sh 'rm -rf $JENKINS_HOME/workspace/$JOB_NAME/*.zip'
            sh 'rm -rf $JENKINS_HOME/workspace/$JOB_NAME/*.exe'
        }
    }

    stage('Build') {
      steps {
        sh 'GOOS=linux go build'
        sh 'GOOS=windows go build'
        sh 'tar cf $JENKINS_HOME/workspace/$JOB_NAME/tools/bolt_exec_puppet.zip *.exe'
      }
    }

    stage('Goquette - NuGet Packaging') {
        steps {
            sh 'go install github.com/PatrickLaabs/goquette@latest'
            sh 'cd $JENKINS_HOME/workspace/$JOB_NAME && $JENKINS_HOME/jobs/$JOB_NAME/builds/$BUILD_ID/bin/goquette'
            }
        }

    stage('nFPM - rpm Packaging') {
        steps {
            sh 'go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest'
            sh '$JENKINS_HOME/jobs/$JOB_NAME/builds/$BUILD_ID/bin/nfpm pkg --packager rpm --target $JENKINS_HOME/workspace/$JOB_NAME/'
        }
    }

    stage('Deploy .nupkg to Nexus') {
        steps {
            echo 'deploying to nexus..'
            nexusArtifactUploader(
                nexusVersion: 'nexus2',
                protocol: 'http',
                nexusUrl: '192.168.86.222:8081/nexus',
                groupId: 'com.example',
                version: '1.0.2',
                repository: 'nuget',
                credentialsId: 'nexus-user-credentials',
                artifacts: [
                    [artifactId: 'bolt_exec_puppet',
                    classifier: 'release',
                     file: '$JENKINS_HOME/workspace/$JOB_NAME/bolt_exec_puppet.nupkg',
                     type: 'nuget']
                ]
             )
        }
    }

    stage('Deploy .rpm to Nexus') {
        steps {
            nexusArtifactUploader(
                nexusVersion: 'nexus2',
                protocol: 'http',
                nexusUrl: '192.168.86.222:8081/nexus',
                groupId: 'com.example',
                version: '1.0.2',
                repository: 'rpm',
                credentialsId: 'nexus-user-credentials',
                artifacts: [
                    [artifactId: 'bolt_exec_puppet',
                     classifier: 'release',
                     file: '$JENKINS_HOME/workspace/$JOB_NAME/bolt_exec_puppet-1.0.2.x86_64.rpm',
                     type: 'rpm']
                ]
             )
        }
    }
  }
}