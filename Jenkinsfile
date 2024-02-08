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
                    def version = "v0.${env.BUILD_NUMBER}-${commitId}"
                    env.IMAGE_TAG = "${DOCKER_REGISTRY}/${IMAGE_NAME}:${version}"
                    sh "docker build -t ${env.IMAGE_TAG} ."
                }
            }
        }
        stage('Push') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'registry-cred', passwordVariable: 'REGISTRY_PASS', usernameVariable: 'REGISTRY_USER')]) {
                        sh "echo '$REGISTRY_PASS' | docker login ${DOCKER_REGISTRY} -u '$REGISTRY_USER' --password-stdin"
                        sh "docker push ${env.IMAGE_TAG}"
                    }   
                }
            }
        }
        stage('Update K8s Deployment') {
            steps {
                script {
                    // Kubernetes manifestindeki imaj versiyonunu güncelleme
                    sh "sed -i 's|image: ${DOCKER_REGISTRY}/${IMAGE_NAME}:.*|image: ${env.IMAGE_TAG}|' k8s/deployment.yaml"
                    sh "git add k8s/deployment.yaml"
                    sh "git commit -m 'Update image version to ${version}'"
                    sh "git push origin main"
                }
            }
        }
    }
    post {
        always {
            // Temizlik adımları, örneğin Docker image'ları silme
        }
    }
}
