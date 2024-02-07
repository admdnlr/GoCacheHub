pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'gocachehub'
        IMAGE_TAG = "${env.BUILD_NUMBER}"
        REGISTRY_URL = 'registry.digitalocean.com/admdnlr'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/admdnlr/GoCacheHub.git'
            }
        }
        stage('Build and Push Docker Image') {
            steps {
                script {
                    // Assuming the Dockerfile is in the root of your project
                    sh """
                    docker run --rm \
                        -v /var/run/docker.sock:/var/run/docker.sock \
                        -v $(pwd):/workspace \
                        gcr.io/kaniko-project/executor:latest \
                        --context /workspace \
                        --dockerfile /workspace/Dockerfile \
                        --destination ${REGISTRY_URL}/${DOCKER_IMAGE}:${IMAGE_TAG}
                    """
                }
            }
        }
    }
}
