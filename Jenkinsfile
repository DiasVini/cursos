String determineRepoName() {
    return scm.getUserRemoteConfigs()[0].getUrl().tokenize('/')[3].split("\\.")[0]
}
pipeline {
  agent {
    node {
      label 'jenkinsAgentBuild-Basic'
    }

  }
    stage('Create Sonar Properties File') {
      steps {
          def repoName = determineRepoName()
        sh '''
sonar.projectKey=atris:${repoName}
sonar.coverage.exclusions=**/bandit/**, **/flake8/**, **/pylint/**, **/govet/**, **/golangci/**
sonar.python.bandit.reportPaths="./bandit"
sonar.python.flake8.reportPaths="./flake8"
sonar.python.pylint.reportPaths="./pylint"
sonar.go.govet.reportPaths="./govet"
sonar.go.golangci-lint.reportPaths="./golangci"
        '''
      }
    }

  }

    stage('External Analyzers') {
      steps {
        sh '''mkdir govet
        mkdir golangci
go tool vet ./ > govet/govet.txt
golangci-lint run ./ > golangci/golanci.txt
'''
      }
    }

  }
}
