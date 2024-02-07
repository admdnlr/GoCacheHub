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
                    // Assuming the Dockerfile is in the root of your project
                    sh """
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v \$(pwd):/workspace gcr.io/kaniko-project/executor:latest --context /workspace --dockerfile /workspace/Dockerfile --destination ${REGISTRY_URL}/${DOCKER_IMAGE}:${IMAGE_TAG}
"""
                }
            }
        }
    }
}
