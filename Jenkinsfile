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
           }
}