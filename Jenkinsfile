pipeline {
        environment {
    GOPATH="${env.WORKSPACE}/.gopath"
  }
  stages {
         stage("pullcode"){
           steps{
             git(
                  url: 'https://github.com/teerawat104/sum.git',
             )
           }
         }
         stage("Setup path") {
            steps {
            sh "mkdir -p .gopath/app/pipeline_example"
            sh "ln -sf ${env.WORKSPACE} .gopath/app/pipeline_example"
      }
    }
         stage("Run test") {
  steps {
        sh "go test -v"
      }
    }
            
          stage("Test build cmd") {
  steps {
        sh "go install github.com/go-debos/fakemachine/cmd/fakemachine"
      }
    }
}

