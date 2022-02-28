pipeline {
  agent any
  tools {
    go 'go-1.17'
  }
  environment {
    GO111MODULE = 'on'
    CGO_ENABLED = 0
    GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }

  stages {
    stage('Build') {
      steps {
        echo 'Building..'
        sh 'go build'
      }
    }

    stage ('Release') {
      when {
        buildingTag()
      }

      environment {
        GITHUB_TOKEN = credentials('github-token')
      }

      steps {
        sh 'curl -sL https://git.io/goreleaser | bash'
      }
    }
  }
}
