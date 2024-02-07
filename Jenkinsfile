pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'gocachehub'
        IMAGE_TAG = "${env.BUILD_NUMBER}" // Veya dinamik bir tag, örneğin: "${env.BUILD_NUMBER}"
        REGISTRY_URL = 'registry.digitalocean.com/admdnlr'
        REGISTRY_CREDENTIALS_ID = '34253c30-a2d1-4368-b280-ec92c1509fb7'
        GIT_CREDENTIALS_ID = 'aa140a10-b6d6-4e82-ad1a-8c5ad3657cea' // Git credentials ID'nizi buraya ekleyin
    }

    stages {
        stage('Clone Repository') {
            steps {
                git credentialsId: env.GIT_CREDENTIALS_ID, url: 'https://github.com/admdnlr/GoCacheHub.git', branch: 'main'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("$REGISTRY_URL/$DOCKER_IMAGE:$IMAGE_TAG")
                }
            }
        }
        stage('Push Image to Registry') {
            steps {
                script {
                    docker.withRegistry("https://$REGISTRY_URL", env.REGISTRY_CREDENTIALS_ID) {
                        docker.image("$REGISTRY_URL/$DOCKER_IMAGE:$IMAGE_TAG").push()
                    }
                }
            }
        }
        stage('Update Git Repository') {
            steps {
                script {
                    // Tag'i güncelleyen ve commit eden script.
                    // Bu örnekte, deployment.yml dosyanızın içindeki image kısmını güncellemeniz gerekecek.
                    // Ayrıca, git kullanıcı adı ve e-posta adresinizi de konfigüre etmeniz gerekmektedir.
                    sh """
                      sed -i 's|$REGISTRY_URL/$DOCKER_IMAGE:.*|$REGISTRY_URL/$DOCKER_IMAGE:$IMAGE_TAG|g' ./Deployment-Manifests/gocachehub-deployment.yaml
                      git config user.name 'admdnlr'
                      git config user.email 'ademdnlr60@gmail.com'
                      git commit -am "Update image tag to $IMAGE_TAG"
                      git push
                    """
                }
            }
        }
    }
}
