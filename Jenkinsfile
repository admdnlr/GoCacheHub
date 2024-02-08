pipeline {
    agent any
    environment {
        DOCKER_REGISTRY = 'registry.digitalocean.com/admdnlr'
        IMAGE_NAME = 'gocachehub'
        version = '1.0.0'
    }
    stages {
        stage('Build') {
            steps {
                script {
                    def commitId = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
                    version = "v0.${env.BUILD_NUMBER}-${commitId}"
                    env.IMAGE_TAG = "${DOCKER_REGISTRY}/${IMAGE_NAME}:${version}"
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
                    // Repository'yi checkout etme ve main branch'e geçiş
                    checkout([
                        $class: 'GitSCM',
                        branches: [[name: '*/main']],
                        doGenerateSubmoduleConfigurations: false,
                        extensions: [],
                        submoduleCfg: [],
                        userRemoteConfigs: [[
                            credentialsId: 'github-cred',
                            url: 'https://github.com/admdnlr/GoCacheHub.git'
                        ]]
                    ])

                    // Kubernetes manifestindeki imaj versiyonunu güncelleme
                    sh "sed -i '' 's|image: ${DOCKER_REGISTRY}/${IMAGE_NAME}:.*|image: ${env.IMAGE_TAG}|' Deployment-Manifests/gocachehub-deployment.yaml"
                    sh "git add Deployment-Manifests/gocachehub-deployment.yaml"
                    // 'version' değişkenini 'env.VERSION' olarak güncelledim. 'version' değişkeninizin tanımını ve değer atamasını doğrulayın.
                    sh "git commit -m 'Update image version to ${version}'"
                    // '-u' seçeneği ile birlikte 'git push' kullanırken, branch'in uzaktaki karşılığıyla ilişkilendirilmesi sağlanır.
                    sh "git push -u origin main"
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
