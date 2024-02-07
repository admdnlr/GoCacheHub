pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'gocachehub'
        IMAGE_TAG = "${env.BUILD_NUMBER}"
        REGISTRY_URL = 'registry.digitalocean.com/admdnlr'
        REGISTRY_CREDENTIALS_ID = '34253c30-a2d1-4368-b280-ec92c1509fb7'
        GIT_CREDENTIALS_ID = 'aa140a10-b6d6-4e82-ad1a-8c5ad3657cea'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git url: 'https://github.com/admdnlr/GoCacheHub.git', branch: 'main', credentialsId: env.GIT_CREDENTIALS_ID
            }
        }
         stage('Build and Push Docker Image') {
            steps {
                script {
                    // Kaniko executor konteynerını kullanarak Docker imajını build edin ve push edin
                    sh """
                    /kaniko/executor --context ${WORKSPACE} \
                                     --dockerfile ${WORKSPACE}/Dockerfile \
                                     --destination ${REGISTRY_URL}/${DOCKER_IMAGE}:${IMAGE_TAG}
                    """
                }
            }
        }
    }
}
