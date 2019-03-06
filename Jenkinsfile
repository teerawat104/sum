pipeline {
   agent any

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
    stage("unit test"){
        steps{
          bat "cd src && go test -cover ./..."
        }
     }
     stage("build docker"){
       steps{
          bat "docker build -t baac/sum -f Dockerfile ."
       }
     }
   }
}