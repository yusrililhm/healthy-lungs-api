pipeline {
    agent any

    environtment {
        AUTHOR = "Yusril Ilham"
    }

    stages {
        stage("Test") {
            sh "go test -v ./test/..."
        }

        stage("Build") {
            sh "go build -o main"
        }

        stage("Deploy") {
            echo "deploy"
        }
    }

    post {
        success {
            mail bcc: '', body: "See <${env.JOB_DISPLAY_URL}>", cc: '', from: '', replyTo: '', subject: "Jenkins Build: ${env.JOB_NAME} #${env.BUILD_ID}", to: 'yusrililham62@gmail.com'
        }

        failure {
            mail bcc: '', body: "See <${env.JOB_DISPLAY_URL}>", cc: '', from: '', replyTo: '', subject: "Jenkins Build Failed: ${env.JOB_NAME} #${env.BUILD_ID}", to: 'yusrililham62@gmail.com'
        }
    }
}
