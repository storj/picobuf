pipeline {
    agent {
        docker {
            label 'main'
            image 'storjlabs/ci:latest'
            alwaysPull true
            args '-u root:root --cap-add SYS_PTRACE -v "/tmp/gomod":/go/pkg/mod'
        }
    }

    options {
          timeout(time: 26, unit: 'MINUTES')
    }

    environment {
        GOTRACEBACK = 'all'
    }

    stages {
        stage('Checkout') {
            steps {
                // Delete any content left over from a previous run.
                sh "chmod -R 777 ."
                // Bash requires extglob option to support !(.git) syntax,
                // and we don't want to delete .git to have faster clones.
                sh 'bash -O extglob -c "rm -rf !(.git)"'

                checkout scm

                sh 'mkdir -p .build'

                // make a backup of the mod file, because sometimes they get modified by tools
                // this allows to lint the unmodified files
                sh 'cp go.mod .build/go.mod.orig'

                // download dependencies
                sh 'go mod download'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -v ./...'
                sh 'go test -v -p 16 -bench XYZXYZXYZXYZ -run XYZXYZXYZXYZ ./...'
            }
        }

        stage('Verification') {
            parallel {
                stage('Lint') {
                    steps {
                        sh 'check-copyright'
                        sh 'check-imports -race ./...'
                        sh 'check-peer-constraints -race'
                        sh 'check-atomic-align ./...'
                        sh 'check-monkit ./...'
                        sh 'check-errs ./...'
                        sh 'staticcheck ./...'
                        sh 'golangci-lint --config /go/ci/.golangci.yml -j=2 run'
                        sh 'go-licenses check ./...'
                        sh 'check-mod-tidy -mod .build/go.mod.orig'
                     }
                }

                stage('Tests') {
                    environment {
                        COVERFLAGS = "${ env.BRANCH_NAME == 'main' ? '-coverprofile=.build/coverprofile -coverpkg=./...' : ''}"
                    }
                    steps {
                        sh 'go vet ./...'
                        sh 'go test -parallel 4 -p 6 -vet=off $COVERFLAGS -timeout 20m -json -race ./... | tee .build/tests.json | xunit -out .build/tests.xml'
                    }
                    post {
                        always {
                            sh script: 'cat .build/tests.json | tparse -all -top -slow 100', returnStatus: true
                            archiveArtifacts artifacts: '.build/tests.json'
                            junit '.build/tests.xml'

                            script {
                                if(fileExists(".build/coverprofile")){
                                    sh script: 'filter-cover-profile < .build/coverprofile > .build/clean.coverprofile', returnStatus: true
                                    sh script: 'gocov convert .build/clean.coverprofile > .build/cover.json', returnStatus: true
                                    sh script: 'gocov-xml  < .build/cover.json > .build/cobertura.xml', returnStatus: true
                                    cobertura coberturaReportFile: '.build/cobertura.xml'
                                }
                            }
                        }
                    }
                }

                stage('Go Compatibility') {
                    steps {
                        sh 'GOOS=linux   GOARCH=amd64 go build ./...'
                        sh 'GOOS=linux   GOARCH=386   go build ./...'
                        sh 'GOOS=linux   GOARCH=arm64 go build ./...'
                        sh 'GOOS=linux   GOARCH=arm   go build ./...'
                        sh 'GOOS=windows GOARCH=amd64 go build ./...'
                        sh 'GOOS=windows GOARCH=386   go build ./...'
                        sh 'GOOS=windows GOARCH=arm64 go build ./...'
                        sh 'GOOS=darwin  GOARCH=amd64 go build ./...'
                        sh 'GOOS=darwin  GOARCH=arm64 go build ./...'

                        sh 'GOOS=linux   GOARCH=amd64 go1.14 build ./...'
                        sh 'GOOS=linux   GOARCH=386   go1.14 build ./...'
                        sh 'GOOS=linux   GOARCH=arm64 go1.14 build ./...'
                        sh 'GOOS=linux   GOARCH=arm   go1.14 build ./...'
                        sh 'GOOS=windows GOARCH=amd64 go1.14 build ./...'
                        sh 'GOOS=windows GOARCH=386   go1.14 build ./...'
                        sh 'GOOS=darwin  GOARCH=amd64 go1.14 build ./...'
                    }
                }
            }
        }
    }
    post {
        always {
            sh "chmod -R 777 ." // ensure Jenkins agent can delete the working directory
            deleteDir()
        }
    }
}