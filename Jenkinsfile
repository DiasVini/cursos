String determineRepoName() {
    return scm.getUserRemoteConfigs()[0].getUrl().tokenize('/')[3].split("\\.")[0]
}

pipeline {

  agent {
    node {
      label 'jenkinsAgentBuild-Basic'
    }
  }
  parameters{
      string(name:'REPONAME', defaultValue: determineRepoName(), description: 'repository name')
  }
  stages{
    stage('Create Sonar Properties File') {
      steps {
        repoName = determineRepoName()
        sh '''
            sonar.projectKey=atris:'''+${params.REPONAME}+'''
            sonar.coverage.exclusions=**/bandit/**, **/flake8/**, **/pylint/**, **/govet/**, **/golangci/**
            sonar.python.bandit.reportPaths="./bandit"
            sonar.python.flake8.reportPaths="./flake8"
            sonar.python.pylint.reportPaths="./pylint"
            sonar.go.govet.reportPaths="./govet"
            sonar.go.golangci-lint.reportPaths="./golangci"
            '''
      }
    }
    stage('External Analyzers Python') {
      steps {
        sh '''
            mkdir govet
            mkdir golangci
            go tool vet ./ > govet/govet.txt
            golangci-lint run ./ > golangci/golanci.txt
            '''
      }
    }
    stage('Prepare and run SonarQube') {
      steps {
        withSonarQubeEnv(installationName: 'SonarQube', credentialsId: 'SonarQubeAuthentication') {
          sh '/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/SonarQube/bin/sonar-scanner'
        }

      }
    }  
  }
  }
