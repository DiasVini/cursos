pipeline {
  agent any
  stages {
    stage('Create Sonar Properties File') {
      steps {
          script {
              writeFile file: 'sonar-project.properties', text: 'sonar.projectKey=atris:${scm.getUserRemoteConfigs()[0].getUrl().tokenize(\'/\')[3].split("\\\\.")[0]}\n sonar.coverage.exclusions=**/bandit/**, **/flake8/**, **/pylint/**, **/govet/**, **/golangci/**\n sonar.python.bandit.reportPaths="./bandit"\n sonar.python.flake8.reportPaths="./flake8"\n sonar.python.pylint.reportPaths="./pylint"\n sonar.go.govet.reportPaths="./govet"\n sonar.go.golangci-lint.reportPaths="./golangci '
          }
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
