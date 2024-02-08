pipeline {
    agent any
    environment {
        DOCKER_REGISTRY = 'registry.digitalocean.com/admdnlr'
        IMAGE_NAME = 'gocachehub'
    }
    stages {
        stage('Build') {
            steps {
                script {
                    def commitId = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
                    env.VERSION = "v0.${env.BUILD_NUMBER}-${commitId}"
                    env.IMAGE_TAG = "${DOCKER_REGISTRY}/${IMAGE_NAME}:${env.VERSION}"
                    sh "/usr/local/bin/docker build -t ${env.IMAGE_TAG} ."
                }
            }
        }
        stage('Push') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'registry-cred', passwordVariable: 'REGISTRY_PASS', usernameVariable: 'REGISTRY_USER')]) {
                        sh "echo '$REGISTRY_PASS' | /usr/local/bin/docker login ${DOCKER_REGISTRY} -u '$REGISTRY_USER' --password-stdin"
                        sh "/usr/local/bin/docker push ${env.IMAGE_TAG}"
                    }
                }
            }
        }
        stage('Update K8s Deployment') {
            steps {
                script {
                    git branch: 'main', credentialsId: 'github-cred', url: 'https://github.com/admdnlr/GoCacheHub.git'
                    
                    sh "sed -i '' 's|image: ${DOCKER_REGISTRY}/${IMAGE_NAME}:.*|image: ${env.IMAGE_TAG}|' Deployment-Manifests/gocachehub-deployment.yaml"
                    sh "git add Deployment-Manifests/gocachehub-deployment.yaml"
                    sh "git commit -m 'Update image version to ${env.VERSION}'"
                    sh "git push origin main"
                }
            }
        }
    }
    post {
        always {
            cleanWs()
            echo 'Geçici dosyalar temizlendi.'
        }
    }
}
