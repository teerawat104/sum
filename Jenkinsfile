pipeline {
   agent any

environment {
    GOPATH="${env.WORKSPACE}
}
  stages {
    stage("pullcode"){
           steps{
             git(
                  url: 'ttps://github.com/teerawat104/sum.git',
             )
             }
           }
    stage("unit test"){
        steps{
          bat "cd src && go test -cover ./..."
        }
     }
   }
}