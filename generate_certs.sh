openssl genrsa -out ~/domain.com.ssl/domain.com.key 2048 -days 365

openssl req -new -sha256 -key cert/tls.key -out cert/tls.csr

openssl x509 -req -days 365 -in cert/tls.csr -signkey cert/tls.key -sha256 -out cert/tls.crt

#Copying certificates
cp cert/tls.key /etc/istio/ingressgateway-certs/tls.key
cp cert/tls.crt /etc/istio/ingressgateway-certs/tls.crt