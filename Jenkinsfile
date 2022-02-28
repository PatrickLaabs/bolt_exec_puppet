pipeline {
  agent any
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    tools {
        go 'go-1.17.7'
    }

  stages {
    stage('Build') {
      steps {
        echo 'Building..'
        sh 'go build'
      }
    }

    stage ('Release') {
        environment {
            GITHUB_TOKEN = credentials('GH_PAT')
        }
        steps {
            sh 'curl -sL https://git.io/goreleaser | bash'
        }
    }
  }
}
