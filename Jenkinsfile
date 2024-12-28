pipeline {
    agent any

     environment {
        DOCKER_IMAGE = 'smemory/cdci'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = '7811591595:AAEAhsr5jqVTEPDHUGBbwh6bJE5pxfgJFbc'
        TELEGRAM_CHAT_ID = '-4662212571'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/PhanNgocNguyen/CICD_Finals.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    echo 'Clearing final_cicd-related images and containers'
                    sh '''
                        docker container stop final_cicd || echo "No container named final_cicd to stop"
                        docker container rm final_cicd || echo "No container named final_cicd to remove"
                        docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image named ${DOCKER_IMAGE}:${DOCKER_TAG} to remove"
                    '''
                }
                echo 'Deploying to DEV environment...'
                sh 'docker image pull smemory/cdci:latest'
                sh 'docker container stop final_cicd || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network still exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name final_cicd -p 4000:4000 --network dev smemory/cdci:latest'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}
